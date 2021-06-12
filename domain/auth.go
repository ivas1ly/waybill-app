package domain

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"

	"github.com/ivas1ly/waybill-app/internal"

	"go.uber.org/zap"

	"github.com/pquerna/otp/totp"

	"github.com/sethvargo/go-password/password"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/ivas1ly/waybill-app/models"
)

func (d *Domain) LoginUser(ctx context.Context, input models.Login) (*models.AuthResponse, error) {
	user, err := d.UsersRepository.GetUserByEmail(input.Email)
	if err != nil {
		d.Logger.Error("Bad credentials. Incorrect login.")
		return nil, gqlerror.Errorf("Incorrect login or password.")
	}

	err = user.CompareUserPassword(input.Password)
	if err != nil {
		d.Logger.Error("Bad credentials. Incorrect password.")
		return nil, gqlerror.Errorf("Incorrect login or password.")
	}
	token, err := user.GenerateTokenPair()
	if err != nil {
		d.Logger.Error("Can't generate token pair.")
		return nil, gqlerror.Errorf("Internal server error.")
	}
	d.Logger.Info(token["accessToken"])
	d.Logger.Info(token["refreshToken"])

	_, err = d.UsersRepository.UpdateRefresh(user, token["refreshToken"])
	if err != nil {
		d.Logger.Error("Can't update user refresh token.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	c, err := d.GetCurrentServerCTX(ctx)
	if err != nil {
		return nil, err
	}
	refreshTime, _ := time.Parse(time.RFC3339, token["refreshExpiresAt"])
	cookie := new(fiber.Cookie)
	cookie.Name = "waybill-app-refresh-token"
	cookie.Value = token["refreshToken"]
	cookie.Expires = refreshTime
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return &models.AuthResponse{
		AccessToken: &models.Token{
			AccessToken:      token["accessToken"],
			AccessExpiredAt:  token["accessExpiresAt"],
			RefreshToken:     token["refreshToken"],
			RefreshExpiredAt: token["refreshExpiresAt"],
		},
		User: user,
	}, nil
}

func (d *Domain) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	_, err := d.UsersRepository.GetUserByEmail(input.Email)
	if err == nil {
		d.Logger.Error("Email already in use.")
		return nil, gqlerror.Errorf("This email already in use.")
	}

	genPassword, err := password.Generate(15, 6, 0, false, false)
	if err != nil {
		d.Logger.Error("An error occurred while creating a user.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	var role models.Role
	if len(input.Role.String()) != 0 && input.Role.IsValid() {
		role = *input.Role
	} else {
		d.Logger.Error("Invalid role.")
		return nil, gqlerror.Errorf("Invalid role.")
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Waybill",
		AccountName: input.Email,
	})
	if err != nil {
		d.Logger.Error("Error while creating TOTP.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, gqlerror.Errorf("Internal server error.")
	}
	png.Encode(&buf, img)

	user := &models.User{
		Email:  input.Email,
		Role:   role,
		Secret: key.Secret(),
	}

	d.Logger.Info(key.Secret())

	err = user.HashPassword(genPassword)
	if err != nil {
		d.Logger.Error("Password hashing error", zap.Error(err))
		return nil, gqlerror.Errorf("Internal server error.")
	}

	if _, err := d.UsersRepository.CreateUser(user); err != nil {
		d.Logger.Error("Error while creating a new user.", zap.Error(err))
		return nil, gqlerror.Errorf("Internal server error.")
	}

	err = internal.SendEmail(input.Email, genPassword, key.Secret(), buf.Bytes())
	if err != nil {
		d.Logger.Error("Error while sending email.", zap.Error(err))
		return nil, gqlerror.Errorf("Internal server error.")
	}

	return user, nil
}

func (d *Domain) GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {

	if ctx.Value("CurrentUser") == nil {
		d.Logger.Error("There is no user in the context. Access is denied.")
		return nil, gqlerror.Errorf("Unauthenticated request. Access Denied.")
	}

	user, ok := ctx.Value("CurrentUser").(*models.User)
	if !ok || user.ID == "" {
		d.Logger.Error("There is no user in the context. Access is denied.")
		return nil, gqlerror.Errorf("Unauthenticated request. Access Denied.")
	}

	return user, nil
}

func (d *Domain) GetCurrentServerCTX(ctx context.Context) (*fiber.Ctx, error) {

	if ctx.Value("serverContext") == nil {
		d.Logger.Error("There is no server ctx in the context.")
		return nil, gqlerror.Errorf("No server ctx.")
	}

	c, ok := ctx.Value("serverContext").(*fiber.Ctx)
	if !ok || c.String() == "" {
		d.Logger.Error("There is no server ctx in the context.")
		return nil, gqlerror.Errorf("There is no server ctx in the context.")
	}

	return c, nil
}

func (d *Domain) Refresh(ctx context.Context) (*models.AuthResponse, error) {

	c, err := d.GetCurrentServerCTX(ctx)
	if err != nil {
		return nil, err
	}

	noCookie := "No cookie."
	tokenFromCookie := c.Cookies("waybill-app-refresh-token", noCookie)
	if tokenFromCookie == noCookie {
		d.Logger.Error("There is no cookie. Can't refresh tokens.")
		return nil, gqlerror.Errorf("Internal server error.")
	}
	fmt.Println(tokenFromCookie)
	claims := &jwt.StandardClaims{}
	tokenFromJWT, err := jwt.ParseWithClaims(tokenFromCookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("auth.signing_key")), nil
	})
	fmt.Printf("%s ", viper.GetString("auth.signing_key"))
	fmt.Println(err)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			d.Logger.Error("JWT Signature Invalid.")
			return nil, gqlerror.Errorf("Unauthorized.")
		}
		d.Logger.Error("Can't parse token from JWT.")
		return nil, gqlerror.Errorf("Bad request.")
	}
	if !tokenFromJWT.Valid {
		d.Logger.Error("Invalid refresh token.")
		return nil, gqlerror.Errorf("Unauthorized.")
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 5*time.Minute {
		d.Logger.Error("Expired token.")
		return nil, gqlerror.Errorf("Bad request.")
	}

	user, err := d.UsersRepository.GetUserByID(claims.Subject)
	if user.RefreshToken != tokenFromJWT.Raw {
		d.Logger.Error("There is a different token in the database.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	token, err := user.GenerateTokenPair()
	if err != nil {
		d.Logger.Error("Can't generate token pair.")
		return nil, gqlerror.Errorf("Internal server error.")
	}
	d.Logger.Info(token["accessToken"])
	d.Logger.Info(token["refreshToken"])

	_, err = d.UsersRepository.UpdateRefresh(user, token["refreshToken"])
	if err != nil {
		d.Logger.Error("Can't update user refresh token.")
		return nil, gqlerror.Errorf("Internal server error.")
	}

	refreshTime, _ := time.Parse(time.RFC3339, token["refreshExpiresAt"])
	cookie := new(fiber.Cookie)
	cookie.Name = "waybill-app-refresh-token"
	cookie.Value = token["refreshToken"]
	cookie.Expires = refreshTime
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return &models.AuthResponse{
		AccessToken: &models.Token{
			AccessToken:      token["accessToken"],
			AccessExpiredAt:  token["accessExpiresAt"],
			RefreshToken:     token["refreshToken"],
			RefreshExpiredAt: token["refreshExpiresAt"],
		},
		User: user,
	}, nil
}

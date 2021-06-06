package domain

import (
	"bytes"
	"context"
	"image/png"

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

	err = user.ComparePassword(input.Password)
	if err != nil {
		d.Logger.Error("Bad credentials. Incorrect password.")
		return nil, gqlerror.Errorf("Incorrect login or password.")
	}
	token, err := user.GenerateTokenPair()
	d.Logger.Info(token["accessToken"])
	d.Logger.Info(token["refreshToken"])

	return &models.AuthResponse{
		AccessToken: &models.Token{
			AccessToken:  token["accessToken"],
			ExpiredAt:    token["expiresAt"],
			RefreshToken: token["refreshToken"],
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

	err = internal.SendEmail(input.Email, genPassword, key.Secret(), buf.Bytes())
	if err != nil {
		d.Logger.Error("Error while sending email.", zap.Error(err))
		return nil, gqlerror.Errorf("Internal server error.")
	}

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

	if _, err := user.GenerateTokenPair(); err != nil {
		d.Logger.Error("Error while creating a pair of tokens", zap.Error(err))
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

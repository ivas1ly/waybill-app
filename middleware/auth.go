package middleware

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/ivas1ly/waybill-app/domain"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(d *domain.Domain) fiber.Handler {

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(viper.GetString("auth.signing_key")),
		SuccessHandler: func(c *fiber.Ctx) error {
			d.Logger.Info("Valid JWT")
			u := c.Locals("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)

			user, err := d.UsersRepository.GetUserByID(claims["sub"].(string))
			if err != nil {
				d.Logger.Info("Can't get data from database.")
				return gqlerror.Errorf("Internal server error.")
			}

			c.Locals("CurrentUser", user)
			return c.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			d.Logger.Info("Empty or invalid JWT.")
			return ctx.Next()
		},
		ContextKey: "user",
	})
}

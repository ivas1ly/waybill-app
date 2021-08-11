module github.com/ivas1ly/waybill-app

go 1.16

// +heroku goVersion go1.16

replace github.com/99designs/gqlgen v0.13.0 => github.com/arsmn/gqlgen v0.13.2

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/gofiber/fiber/v2 v2.12.0
	github.com/gofiber/jwt/v2 v2.2.2
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/klauspost/compress v1.13.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pquerna/otp v1.3.0
	github.com/sethvargo/go-password v0.2.0
	github.com/spf13/viper v1.7.1
	github.com/vektah/gqlparser/v2 v2.1.0
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)

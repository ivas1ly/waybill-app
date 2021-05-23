module github.com/ivas1ly/waybill-app

go 1.16

// +heroku goVersion go1.16

replace github.com/99designs/gqlgen v0.13.0 => github.com/arsmn/gqlgen v0.13.2

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/gofiber/fiber/v2 v2.7.1
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/klauspost/compress v1.12.1 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/savsgio/gotils v0.0.0-20210316171653-c54912823645 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/valyala/fasthttp v1.23.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20210415045647-66c3f260301c // indirect
	golang.org/x/tools v0.1.0 // indirect
	gorm.io/gorm v1.21.8
)

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/ivas1ly/waybill-app) [![CodeQL](https://github.com/ivas1ly/waybill-app/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/ivas1ly/waybill-app/actions/workflows/codeql-analysis.yml) [![Go](https://github.com/ivas1ly/waybill-app/actions/workflows/go.yml/badge.svg)](https://github.com/ivas1ly/waybill-app/actions/workflows/go.yml) [![gographs](https://gographs.io/badge.svg)](https://gographs.io/repo/github.com/ivas1ly/waybill-app?cluster=true)

# Waybill App
Waybill processing service (module) written with Golang.

## Requirements
- Go 1.16
- PostgreSQL 13

## Tech Stack
* Golang
  * [Fiber](https://github.com/gofiber/fiber)
  * [GORM](https://github.com/go-gorm/gorm)
  * [Gqlgen](https://github.com/99designs/gqlgen)
  * [Zap](https://github.com/uber-go/zap)
  * [OTP](https://github.com/pquerna/otp/)
* PostgreSQL

## Features

* 2fa authentication with OTP and JWT tokens.
* Auto migrations.
* Clean Architecture.
* Creation of reports on completed waybills.
* GraphQL API.

## License

This repository is available under the [MIT License](https://github.com/ivas1ly/waybill-app/blob/main/LICENSE).

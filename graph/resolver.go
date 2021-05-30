package graph

import "github.com/ivas1ly/waybill-app/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	Domain *domain.Domain
}

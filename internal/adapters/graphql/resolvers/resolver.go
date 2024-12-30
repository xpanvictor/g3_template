package graph

import "go_graphql_gin/internal/core/ports"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService ports.UserService
}

type Config struct {
	UserService ports.UserService
	// Add other services as needed
}

func NewResolver(cfg Config) *Resolver {
	return &Resolver{
		UserService: cfg.UserService,
	}
}

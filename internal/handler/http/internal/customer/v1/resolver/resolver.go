package resolver

import "github.com/Pochirify/pochirify-backend/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App usecase.App
}

type Config struct {
	App usecase.App
}

func NewResolver(c *Config) *Resolver {
	return &Resolver{
		App: c.App,
	}
}

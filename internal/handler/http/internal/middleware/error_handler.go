package middleware

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// hide internal error message inside pochirify server
func NewErrorHandler(lf logger.Factory) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		lf(ctx).Error(err, "internal error")
		return graphql.DefaultErrorPresenter(ctx, errors.New("internal error"))
	}
}

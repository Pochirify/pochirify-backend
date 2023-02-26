package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/spanner"
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	appspanner "github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/Pochirify/pochirify-backend/internal/domain/ec/shopify"
	"github.com/Pochirify/pochirify-backend/internal/domain/payment"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/resolver"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/customer/v1/schema"
	"github.com/Pochirify/pochirify-backend/internal/handler/http/internal/middleware"
	webhookv1 "github.com/Pochirify/pochirify-backend/internal/handler/http/internal/webhook/v1"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"github.com/Pochirify/pochirify-backend/internal/usecase"
)

type Server struct {
	port   uint
	router *chi.Mux
}

type Config struct {
	Port    uint
	Logger  logger.Logger
	Spanner *spanner.Client

	PayPayClient     payment.PaypayClient
	CreditCardClient payment.CreditCardClient
	ShopifyClient    shopify.ShopifyClient
}

func NewServer(ctx context.Context, c *Config) *Server {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		// AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))
	r.Use(middleware.NewRequestLogger(newLoggerFactory(c.Logger.WithName("request_logger"))))

	spanner := appspanner.NewSpanner(c.Spanner, newLoggerFactory(c.Logger.WithName("spanner")))
	repositories := appspanner.InitRepositories(spanner)
	app := usecase.NewApp(&usecase.Config{
		PaypayClient:     c.PayPayClient,
		CreditCardClient: c.CreditCardClient,
		ShopifyClient:    c.ShopifyClient,
		Repositories:     repositories,
	})

	config := &resolver.Config{App: app}
	srv := gqlhandler.NewDefaultServer(
		schema.NewExecutableSchema(
			schema.Config{Resolvers: resolver.NewResolver(config)},
		),
	)
	srv.SetErrorPresenter(middleware.NewErrorHandler(newLoggerFactory(c.Logger.WithName("error_handler"))))

	webhookHandler := webhookv1.NewWebhookHandler(app)
	r.Handle("/api/webhook", webhookHandler.PayPayTransactionEventHandler())

	r.Handle("/api/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", c.Port)

	return &Server{
		port:   c.Port,
		router: r,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}

// TODO: add configure traceID for logger
func newLoggerFactory(l logger.Logger) logger.Factory {
	return func(ctx context.Context) logger.Logger {
		// md, ok := metadata.FromIncomingContext(ctx)
		// if !ok {
		// 	return l
		// }

		// rid := md.Get("X-ASIA-REQUEST-ID")
		// if len(rid) == 0 {
		// 	return l
		// }

		// traceContextID := tracerctx.GetTraceContextID(ctx)

		// return l.WithValues("request_id", rid).WithValues("logging.googleapis.com/trace", traceContextID)
		return l
	}
}

// func handler() http.Handler {}

package spanner

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner/yo"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const databaseNamePattern = "projects/%s/instances/%s/databases/%s"

type ctxTnxKey struct{}

type Spanner struct {
	client *spanner.Client
	logger logger.Factory
}

type ClientConfig struct {
	ProjectID  string
	InstanceID string
	DatabaseID string
	// SessionPoolMaxIdle       uint64
	// SessionPoolWriteSessions float64
}

func NewClient(ctx context.Context, conf *ClientConfig) (*spanner.Client, error) {
	client, err := spanner.NewClient(
		ctx,
		getDatabaseName(conf.ProjectID, conf.InstanceID, conf.DatabaseID),
		// NOTE: https://github.com/x-asia/kauche-app/pull/721/files#r747595937
		// spanner.ClientConfig{
		// 	SessionPoolConfig: spanner.SessionPoolConfig{
		// 		MaxIdle:       conf.SessionPoolMaxIdle,
		// 		WriteSessions: conf.SessionPoolWriteSessions,
		// 	},
		// },
		option.WithGRPCDialOption(grpc.WithBlock()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.WaitForReady(true))),
		// option.WithGRPCDialOption(grpc.WithUnaryInterceptor(tracer.NewGRPCClientInterceptor())),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to initialize spanner client", err)
	}

	return client, nil
}

func getDatabaseName(projectID, instanceID, databaseID string) string {
	databaseName := fmt.Sprintf(databaseNamePattern, projectID, instanceID, databaseID)

	return databaseName
}

func NewSpanner(client *spanner.Client, logger logger.Factory) *Spanner {
	s := &Spanner{
		client: client,
		logger: logger,
	}

	return s
}

func (s Spanner) Ctx(ctx context.Context) yo.YORODB {
	if tnx, ok := ctx.Value(ctxTnxKey{}).(*spanner.ReadWriteTransaction); ok {
		return tnx
	}
	return s.client.Single()
}

func (s Spanner) ApplyMutations(ctx context.Context, ms []*spanner.Mutation) (*time.Time, error) {
	if tnx, ok := ctx.Value(ctxTnxKey{}).(*spanner.ReadWriteTransaction); ok {
		if err := tnx.BufferWrite(ms); err != nil {
			return nil, err
		}
		t := time.Now()
		return &t, nil
	}

	t, err := s.client.Apply(ctx, ms)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// func (s Spanner) Transaction(ctx context.Context, fc func(context.Context) error) error {
// 	if tnx := ctx.Value(ctxTnxKey{}); tnx != nil {
// 		return errors.New("already in transaction")
// 	}

// 	if _, err := s.client.ReadWriteTransaction(ctx, func(ctx context.Context, tnx *spanner.ReadWriteTransaction) error {
// 		return fc(context.WithValue(ctx, ctxTnxKey{}, tnx))
// 	}); err != nil {
// 		return fmt.Errorf("in transaction: %w", err)
// 	}

// 	return nil
// }

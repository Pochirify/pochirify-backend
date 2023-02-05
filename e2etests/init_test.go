// go:build e2e
package e2etests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	appspanner "github.com/Pochirify/pochirify-backend/internal/handler/db/spanner"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
	"github.com/Pochirify/pochirify-backend/internal/handler/logger/json"

	"github.com/Pochirify/pochirify-backend/e2etests/gqlgenc"
	"github.com/Pochirify/pochirify-backend/internal/domain/repository"
)

func newClient(_ *testing.T) gqlgenc.GraphQLClient {
	return gqlgenc.NewClient(
		http.DefaultClient,
		"http://localhost:"+port+"/api/query",
	)
}

func initRepositories() repository.Repositories {
	client, err := spanner.NewClient(
		context.Background(),
		&appspanner.ClientConfig{
			ProjectID:  projectID,
			InstanceID: instanceID,
			DatabaseID: databaseID,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create spanner client: %s", err.Error()))
	}

	return appspanner.InitRepositories(appspanner.NewSpanner(client, newLoggerFactory()))
}

func newLoggerFactory() logger.Factory {
	l, err := json.NewLogger()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize new logger: %s", err.Error()))
	}
	return func(ctx context.Context) logger.Logger {
		return l
	}
}

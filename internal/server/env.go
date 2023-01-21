package server

import (
	"github.com/kelseyhightower/envconfig"
)

type environments struct {
	Port              uint   `envconfig:"PORT" required:"true"`
	GCPProjectID      string `envconfig:"GCP_PROJECT_ID" required:"true"`
	SpannerInstanceID string `envconfig:"SPANNER_INSTANCE_ID" required:"true"`
	SpannerDatabaseID string `envconfig:"SPANNER_DATABASE_ID" required:"true"`
}

func getEnvironments() (*environments, error) {
	var e environments
	if err := envconfig.Process("", &e); err != nil {
		return nil, err
	}

	return &e, nil
}

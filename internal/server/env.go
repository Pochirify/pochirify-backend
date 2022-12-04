package server

import (
	"github.com/kelseyhightower/envconfig"
)

type environments struct {
	Port uint `envconfig:"PORT" required:"true"`
}

func getEnvironments() (*environments, error) {
	var e environments
	if err := envconfig.Process("", &e); err != nil {
		return nil, err
	}

	return &e, nil
}

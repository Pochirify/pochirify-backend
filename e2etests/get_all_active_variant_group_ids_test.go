// go:build e2e
package e2etests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllActiveVariantGroupIDs_Normal(t *testing.T) {
	client := newClient(t)
	_, err := client.GetAllActiveVariantGroupIDs(context.Background())
	assert.NoError(t, err)
}

//go:build tools
// +build tools

// TODO: dev配下へ移動
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

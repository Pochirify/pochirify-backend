package spanner

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isNotFoundErr(err error) bool {
	return status.Code(err) == codes.NotFound
}

type field struct {
	name  string
	value string
}

func findError(fields []field, rawError error, wrapperError ...error) error {
	format := "failed to find by"
	for i, f := range fields {
		if i != 0 {
			format += " and"
		}
		format += " " + f.name + "=" + f.value
	}

	if len(wrapperError) == 0 {
		format += ", err=%w"
		return fmt.Errorf(format, rawError)
	}
	format += ", err=%s: %w"
	return fmt.Errorf(format, rawError, wrapperError[0])
}

package spanner

import "cloud.google.com/go/spanner"

func toSpannerNullString(val *string) spanner.NullString {
	if val == nil {
		return spanner.NullString{}
	}

	return spanner.NullString{StringVal: *val, Valid: true}
}

func fromSpannerNullString(val spanner.NullString) *string {
	if val.IsNull() {
		return nil
	}
	return &val.StringVal
}

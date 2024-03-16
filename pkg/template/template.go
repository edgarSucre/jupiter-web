package template

import "context"

type hxRequestKey string

var IsHxRequestKey hxRequestKey = "isHxRequestKey"

func IsHxRequest(ctx context.Context) bool {
	if val, ok := ctx.Value(IsHxRequestKey).(bool); ok {
		return val
	}

	return false
}

package ctxdata

import (
	"context"
	"encoding/json"
)

func GetUId(ctx context.Context) int64 {
	if u, ok := ctx.Value(Identify).(json.Number); ok {
		if uid, err := u.Int64(); err == nil {
			return uid
		}
	}
	return 0
}

package ctxdata

import (
	"context"
	"encoding/json"
)

func GetUId(ctx context.Context) uint64 {
	if u, ok := ctx.Value(Identify).(json.Number); ok {
		if uid, err := u.Int64(); err == nil {
			return uint64(uid)
		}
	}
	return 0
}

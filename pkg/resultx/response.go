/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package resultx

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zrpcErr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"

	"xlife/pkg/xerr"
)

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Success: true,
		Msg:     "",
		Data:    data,
	}
}

func Fail(code int, err string) *Response {
	return &Response{
		Success: false,
		Msg:     err,
		Data:    nil,
	}
}

func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := xerr.ErrMsg(errCode)

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*zrpcErr.CodeMsg); ok {
			errCode = e.Code
			errMsg = e.Msg
		} else {
			if gStatus, ok := status.FromError(causeErr); ok {
				errCode = int(gStatus.Code())
				errMsg = gStatus.Message()
			}
		}

		// 日志记录
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)

		return http.StatusBadRequest, Fail(errCode, errMsg)
	}
}

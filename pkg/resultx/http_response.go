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

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpSuccess(data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: 200,
		Msg:  "",
		Data: data,
	}
}

func HttpFail(code int, err string) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg:  err,
		Data: nil,
	}
}

func HttpOkHandler(_ context.Context, v interface{}) any {
	return HttpSuccess(v)
}

func HttpErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := xerr.ErrMsg(errCode)

		causeErr := errors.Cause(err)
		var zErr *zrpcErr.CodeMsg
		if errors.As(causeErr, &zErr) {
			errCode = zErr.Code
			errMsg = zErr.Msg
		} else {
			if gStatus, ok := status.FromError(causeErr); ok {
				errCode = int(gStatus.Code())
				errMsg = gStatus.Message()
			}
		}

		// 日志记录
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)

		return http.StatusBadRequest, HttpFail(errCode, errMsg)
	}
}

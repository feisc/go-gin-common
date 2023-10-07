package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab/go-gin/common/common/types"
	"gitlab/go-gin/common/util"
	"go.uber.org/zap"
	"reflect"
)

// wrapper the gin.HandlerFunc with error
func ErrorWrapper(handler types.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// request id
		rid := util.GetRequestIdFromCtx(c)
		// response code
		code := CodeSuccess
		// output error message
		var msg string
		err := handler(c)
		if err != nil {
			var apiErr types.APIError

			// as types.APIError
			if errors.As(err, &apiErr) {
				e := apiErr.Err
				switch e := e.(type) {
				case *json.UnmarshalTypeError:
					code = CodeInvalidParam
					msg = e.Error()
					zap.L().Error(apiErr.From,
						zap.Any("request_id", rid),
						zap.Any("result", types.CtrlActFail),
						zap.Error(e),
						zap.Any("error_type", reflect.TypeOf(e)),
					)
				case *json.SyntaxError:
					code = CodeInvalidParam
					msg = e.Error()
					zap.L().Error(apiErr.From,
						zap.Any("request_id", rid),
						zap.Any("result", types.CtrlActFail),
						zap.Error(e),
						zap.Any("error_type", reflect.TypeOf(e)),
					)
				default:
					// generic
					code = CodeGenericFailure
					msg = apiErr.Message()
					zap.L().Error(apiErr.From,
						zap.Any("request_id", rid),
						zap.Any("result", types.CtrlActFail),
						zap.Error(apiErr),
						zap.Any("error_type", reflect.TypeOf(apiErr)),
					)
				}
			} else {
				code = CodeGenericFailure
				msg = code.Msg()
				zap.L().Error("error occured in handler",
					zap.Any("request_id", rid),
					zap.Any("result", types.CtrlActFail),
					zap.Error(err),
					zap.Any("error_type", reflect.TypeOf(err)))
			}

			// response
			ResponseErrorWithMsg(c, code, msg)
			return
		}
	}
}

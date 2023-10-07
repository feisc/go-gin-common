package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab/go-gin/common/common/types"
	"gitlab/go-gin/common/util"
	"go.uber.org/zap"
	"net/url"
)

func BatchMysqlHandlerTest(c *gin.Context) error {
	// 1、接收文件的url
	handler := "BatchMysqlHandlerTest"
	rid := util.GetRequestIdFromCtx(c)

	var req types.BatchMysqlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return types.APIError{
			From: handler,
			Err:  err,
		}
	}

	// 2、判断url是否正确
	if _, err := url.ParseRequestURI(req.SqlFileUrl); err != nil {
		return types.APIError{
			From: handler,
			Err:  err,
		}
	}

	// 3、返回消息
	resp := &types.BatchMysqlResponse{
		Msg: "receive success",
	}
	// success
	zap.L().Info(handler,
		zap.Any("request_id", rid),
		zap.String("url", req.SqlFileUrl),
		zap.Any("result", types.CtrlActSuccess),
		zap.Any("resource", types.ResouceMysql),
	)
	ResponseSuccess(c, resp)
	return nil
}

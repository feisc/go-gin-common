package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab/go-gin/common/util"
	"net/http"
)

/*
{
	"code": 10000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}

*/

type ResponseData struct {
	RequestId string      `json:"requestId"`
	Code      ResCode     `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	requestId := c.GetString(util.RidFlag)
	c.JSON(http.StatusOK, &ResponseData{
		RequestId: requestId,
		Code:      code,
		Msg:       code.Msg(),
		Data:      nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	requestId := c.GetString(util.RidFlag)
	c.JSON(http.StatusOK, &ResponseData{
		RequestId: requestId,
		Code:      code,
		Msg:       msg,
		Data:      nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	requestId := c.GetString(util.RidFlag)
	c.JSON(http.StatusOK, &ResponseData{
		RequestId: requestId,
		Code:      CodeSuccess,
		Msg:       CodeSuccess.Msg(),
		Data:      data,
	})
}

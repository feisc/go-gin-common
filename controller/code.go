package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeServerBusy
	CodeInvalidManifest
	Code404
	CodeGenericFailure
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidManifest: "无效的资源清单",
	Code404:             "api not found",
	CodeGenericFailure:  "通用错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

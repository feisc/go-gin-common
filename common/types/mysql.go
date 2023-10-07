package types

type BatchMysqlRequest struct {
	SqlFileUrl string `json:"sqlFileUrl" binding:"required"`
	File       string `json:"file" binding:"required" `
	NameSpace  string `json:"namespace"`
	NanoId     string `json:"nanoId"`
}

type BatchMysqlResponse struct {
	Msg string `json:"msg"`
}

type ReqCode int64

const (
	CodeBatchExec ReqCode = 1000 + iota
	CodeBatchFailed
	CodeBatchSuccess
	CodeDownloadFailed
	CodeFileErr
	CodeServerErr
	CodeServerUnkonw
)

var codeMsgMap = map[ReqCode]string{
	CodeBatchExec:      "执行中",
	CodeBatchFailed:    "执行失败",
	CodeBatchSuccess:   "执行完成",
	CodeDownloadFailed: "下载失败",
	CodeFileErr:        "文件错误",
	CodeServerErr:      "服务出错",
	CodeServerUnkonw:   "unknow",
}

func (c ReqCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerUnkonw]
	}
	return msg
}

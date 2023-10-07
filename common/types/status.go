package types

type MysqlStatus string

const (
	MysqlDownloadFailed MysqlStatus = "download filed"
	MysqlExecFailed     MysqlStatus = "batch mysql failed"
	MysqlExecSuccess    MysqlStatus = "batch mysql success"
	mysqlImplement      MysqlStatus = "batch mysql implement"
)

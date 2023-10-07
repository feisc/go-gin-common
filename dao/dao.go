package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gitlab/go-gin/common/setting"
)

var DB *sql.DB

func Init(cfg *setting.MysqlConfig) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s&multiStatements=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Charset)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 设置参数
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	//DB.SetConnMaxLifetime(30 * time.Second)

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"gitlab/go-gin/common/dao"
	"gitlab/go-gin/common/logger"
	"gitlab/go-gin/common/routers"
	"gitlab/go-gin/common/setting"
	"go.uber.org/zap"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("config file is required, e.g.: ./zvos-edge-command-control command-control.yaml")
		return
	}

	// init setting
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("failed to load config file: %v\n", err)
		return
	}

	// init logger
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("failed to init logger: %v\n", err)
		return
	}

	// init mysql
	if err := dao.Init(setting.Conf.MysqlConfig); err != nil {
		fmt.Printf("failed to init mysql: %v\n", err)
		return
	}

	// init routers
	r := routers.SetupRouter(setting.Conf.Mode, setting.Conf.HTTPS, setting.Conf.HTTPSPort)

	// run server https
	var err error
	if setting.Conf.HTTPS {
		err = r.RunTLS(fmt.Sprintf(":%d", setting.Conf.HTTPSPort), "/etc/zvos-edge/tls/server.crt", "/etc/zvos-edge/tls/server.key")
	} else {
		// http
		err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	}

	// fatal and log error
	if err != nil {
		zap.L().Fatal("failed to run server", zap.Error(err))
		return
	}
}

package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name            string `mapstructure:"name"`
	Mode            string `mapstructure:"mode"`
	Version         string `mapstructure:"version"`
	UploadStatusUrl string `mapstructure:"upload_status_url"`
	Port            int    `mapstructure:"port"`
	HTTPSPort       int    `mapstructure:"https_port"`
	HTTPS           bool   `mapstructure:"https"`

	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// TODO
type MysqlConfig struct {
	UserName     string `mapstructure:"user_name"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Charset      string `mapstructure:"charset"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)

	// read config
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// unmarshal configurations
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config was changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}

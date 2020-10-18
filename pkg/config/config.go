package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

const defaultConfigFile = "./configs/config.yaml"

// init 初始化配置文件
func init() {
	pflag.StringP("configFile", "c", "", "choose config file.")
	pflag.Parse()

	// 优先级: 命令行 > 环境变量 > 默认值
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)

	configFile := v.GetString("configFile")
	if configFile == "" {
		configFile = defaultConfigFile
	}

	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(Conf); err != nil {
			fmt.Println("v.OnConfigChange -> v.Unmarshal(Conf) failed,err:", err)
		}
	})
	if err := v.Unmarshal(Conf); err != nil {
		fmt.Println("v.Unmarshal(Conf) failed,err:", err)
	}

	fmt.Println("Config sets success，Conf:", Conf)
}

// AppConfig 应用配置
type AppConfig struct {
	AppName   string `mapstructure:"app_name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

// MySQLConfig mysql配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

// LogConfig zap log配置
type LogConfig struct {
	Prefix     string `mapstructure:"prefix"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

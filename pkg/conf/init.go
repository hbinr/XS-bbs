package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "../config.yaml"

// Init init config
func Init() (*Config, error) {
	var (
		conf Config
		err  error
	)
	pflag.StringP("conf", "c", "", "choose conf file.")
	pflag.Parse()

	// 优先级: 命令行 > 环境变量 > 默认值
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)

	configFile := v.GetString("conf")
	if configFile == "" {
		configFile = defaultConfigFile
	}

	v.SetConfigFile(configFile)
	if err = v.ReadInConfig(); err != nil {
		fmt.Println("ReadInConfig:", err)
		panic(fmt.Errorf("Fatal error conf file: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println("v.OnConfigChange -> v.Unmarshal(&conf) failed,err:", err)
			return
		}
	})
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println("v.Unmarshal failed,err:", err)
		return &conf, nil
	}
	fmt.Println("Config sets success，Conf:", &conf)
	return &conf, nil
}

package conf

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "../config.yaml"

// Init init config
func Init() (conf *Config, err error) {

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
		panic(fmt.Sprintf("Fatal error conf file: %s", err))
	}
	// 监控config改变
	watchConfig(conf, v)

	if err = v.Unmarshal(&conf); err != nil {
		log.Fatalf("config unmarshal, err: %+v", err)
	}

	return
}

func watchConfig(conf *Config, v *viper.Viper) {
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err := v.Unmarshal(conf); err != nil {
			fmt.Println("v.OnConfigChange -> v.Unmarshal(&conf) failed,err:", err)
			return
		}
	})
}

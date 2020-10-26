package conf

// Config 应用配置
type Config struct {
	System      `mapstructure:"system"`
	LogConfig   `mapstructure:"log"`
	MySQLConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
}

type System struct {
	AppName   string `mapstructure:"app_name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`
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

// MySQLConfig mysql配置
type MySQLConfig struct {
	Addr         string `mapstructure:"addr"`           // host:port,for trace
	DSN          string `mapstructure:"dsn"`            // write data source name.
	LogMode      bool   `mapstructure:"log_mode"`       // whether to open the log
	MaxOpenConns int    `mapstructure:"max_open_conns"` // max open conns
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // max idle conns
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

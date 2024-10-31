package config

import (
	"github.com/spf13/viper"
	"time"
)

var Appconfig *Config

type Config struct {
	DBDriver      string     `mapstructure:"DB_DRIVER"`
	DBSource      string     `mapstructure:"DB_SOURCE"`
	ServerAddress string     `mapstructure:"SERVER_ADDRESS"`
	PostConfig    PostConfig `mapstructure:"POST_CONFIG"`
	LogLevel      int8       `mapstructure:"LOG_LEVEL"`
	Env           string     `mapstructure:"ENV"`
}

type PostConfig struct {
	DBHost       string        `mapstructure:"DB_HOST"`
	DBPort       string        `mapstructure:"DB_PORT"`
	DBUser       string        `mapstructure:"POSTGRES_USER"`
	DBPassword   string        `mapstructure:"POSTGRES_PASSWORD"`
	DBName       string        `mapstructure:"POSTGRES_DB"`
	DBSSLMode    string        `mapstructure:"DB_SSLMODE"`
	MaxOpenConns int           `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns int           `mapstructure:"MAX_IDLE_CONNS"`
	MaxIdleTime  time.Duration `mapstructure:"MAX_IDLE_TIME"`
	Timeout      time.Duration `mapstructure:"TIMEOUT"`
}

func LoadingConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	Appconfig = &config
	return nil
}

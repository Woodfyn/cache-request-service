package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	v = viper.New()
)

type Config struct {
	RedisConfig RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int `mapstructure:"db"`
}

func NewConfig(folder string, filename string) *Config {
	var cfg *Config

	v.AddConfigPath(folder)
	v.SetConfigName(filename)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	if err := setEnvConfig(cfg); err != nil {
		panic(err)
	}

	return cfg
}

func setEnvConfig(cfg *Config) error {
	if err := godotenv.Load("main.env"); err != nil {
		return err
	}

	cfg.RedisConfig.Addr = os.Getenv("REDIS_ADDR")
	cfg.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")

	return nil
}

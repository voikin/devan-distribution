package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Pg   PGConfig
	HTTP HTTPConfig
	JWT  JWTConfig
}

type PGConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"ssl_mode"`
}

type HTTPConfig struct {
	Host                  string        `json:"host"`
	Port                  string        `json:"port"`
	ShutdownServerTimeout time.Duration `josn:"shutdownServerTimeout"`
}

type JWTConfig struct {
	Salt       string `json:"salt"`
	SigningKey string `json:"signingKey"`
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func (c PGConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s auth=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSLMode)
}

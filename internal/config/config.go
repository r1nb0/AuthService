package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Password PasswordConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type PasswordConfig struct {
	Salt             string
	MinLength        uint
	MaxLength        uint
	IncludeLowercase bool
	IncludeUppercase bool
	IncludeChars     bool
	IncludeDigits    bool
}

type JWTConfig struct {
	Secret string
}

func GetConfig(configPath, configName, configType string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg, err := parseConfig(v)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	cfg.Postgres.Password = os.Getenv("DATABASE_PASS")
	cfg.Password.Salt = os.Getenv("SALT")
	cfg.JWT.Secret = os.Getenv("JWT_SECRET")
	return &cfg, nil
}

package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	cfgPath string = "./configs"
	cfgName string = "config-prod"
	cfgType string = "yaml"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Password PasswordConfig
	JWT      JWTConfig
	Logger   LoggerConfig
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
	MinLength        int
	MaxLength        int
	IncludeLowercase bool
	IncludeUppercase bool
	IncludeChars     bool
	IncludeDigits    bool
	IncludeSpecial   bool
}

type LoggerConfig struct {
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

type JWTConfig struct {
	Secret string
}

func GetConfig() *Config {
	v := viper.New()
	v.AddConfigPath(cfgPath)
	v.SetConfigName(cfgName)
	v.SetConfigType(cfgType)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error of initializing config: %s", err.Error())
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatalf("error of parsing config: %s", err.Error())
	}
	return cfg
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

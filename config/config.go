package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	ListenPort  string
	TokenSecret string
	Postgres    *PostgresConfig
	Redis       *RedisConfig
}

type PostgresConfig struct {
	Host          string
	Database      string
	User          string
	Password      string
	Port          string
	RunMigrations string
}

type RedisConfig struct {
	Host            string
	SessionDuration string
	Port            string
}

func Init() error {
	if !EnvLoaded() {
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf("failed to load .env file: %v", err)
		}
	}
	return nil
}

func EnvLoaded() bool {
	requiredVars := []string{
		"POSTGRES_HOST",
		"POSTGRES_DB",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_PORT",
		"REDIS_HOST",
		"SESSION_DURATION",
		"REDIS_PORT",
		"LISTEN_PORT",
		"RUN_MIGRATIONS",
		"TOKEN_SECRET",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return false
		}
	}
	return true
}

func Get() *Config {
	return &Config{
		ListenPort:  os.Getenv("LISTEN_PORT"),
		TokenSecret: os.Getenv("TOKEN_SECRET"),
		Postgres: &PostgresConfig{
			Host:          os.Getenv("POSTGRES_HOST"),
			Database:      os.Getenv("POSTGRES_DB"),
			User:          os.Getenv("POSTGRES_USER"),
			Password:      os.Getenv("POSTGRES_PASSWORD"),
			Port:          os.Getenv("POSTGRES_PORT"),
			RunMigrations: os.Getenv("RUN_MIGRATIONS"),
		},
		Redis: &RedisConfig{
			Host:            os.Getenv("REDIS_HOST"),
			SessionDuration: os.Getenv("SESSION_DURATION"),
			Port:            os.Getenv("REDIS_PORT"),
		},
	}
}

func (c *Config) SessionDuration() (time.Duration, error) {
	d, err := strconv.Atoi(c.Redis.SessionDuration)
	if err != nil {
		return 0, err
	}
	return time.Duration(d), nil
}

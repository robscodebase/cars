package tests

import (
	"cars/config"
	"log"
	"os"
)

const (
	postgresHost     = "localhost"
	postgresPort     = "5432"
	postgresDB       = "cars"
	postgresUser     = "postgres"
	postgresPassword = "postgres"
	runMigrations    = "false"
	redisHost        = "localhost"
	redisPort        = "6379"
	sessionDuration  = "3600"
	listenPort       = "8080"
	tokenSecret      = "7cP0iH2EfUv9y$B^5p4s%#jL0kZvV@r&"
)

func SetTestEnvVars() {
	configLoaded := config.EnvLoaded()
	if configLoaded {
		log.Println("config already loaded")
	}
	if !configLoaded {
		os.Setenv("POSTGRES_HOST", postgresHost)
		os.Setenv("POSTGRES_PORT", postgresPort)
		os.Setenv("POSTGRES_DB", postgresDB)
		os.Setenv("POSTGRES_USER", postgresUser)
		os.Setenv("POSTGRES_PASSWORD", postgresPassword)
		os.Setenv("RUN_MIGRATIONS", runMigrations)
		os.Setenv("REDIS_HOST", redisHost)
		os.Setenv("REDIS_PORT", redisPort)
		os.Setenv("SESSION_DURATION", sessionDuration)
		os.Setenv("LISTEN_PORT", listenPort)
		os.Setenv("TOKEN_SECRET", tokenSecret)
	}
	if !config.EnvLoaded() {
		panic("failed to set test environment variables")
	}
}

package config

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var env sync.Map

type archetypeConfiguration struct {
	// HTTP client is enabled by default
	EnableSecretManager bool
	EnvironmentPath     string
	EnablePostgreSQLDB  bool
	EnablePubSub        bool
	EnableFirestore     bool
	EnableCobraCli      bool
	EnableHTTPServer    bool
	EnableRedis         bool
	EnableRestyClient   bool
}

func (e *archetypeConfiguration) SetPubsub(enable bool) {
	e.EnablePubSub = enable
}

type Config string

// Archetype default configuration
const (
	PORT         Config = "PORT"
	COUNTRY      Config = "COUNTRY"
	PROJECT_NAME Config = "PROJECT_NAME"
	ENVIRONMENT  Config = "ENVIRONMENT"
)

// Google Cloud Platform configuration
const GOOGLE_PROJECT_ID Config = "GOOGLE_PROJECT_ID"

// PostgreSQL configuration
const (
	DATABASE_POSTGRES_HOSTNAME Config = "DATABASE_POSTGRES_HOSTNAME"
	DATABASE_POSTGRES_PORT     Config = "DATABASE_POSTGRES_PORT"
	DATABASE_POSTGRES_NAME     Config = "DATABASE_POSTGRES_NAME"
	DATABASE_POSTGRES_USERNAME Config = "DATABASE_POSTGRES_USERNAME"
	DATABASE_POSTGRES_PASSWORD Config = "DATABASE_POSTGRES_PASSWORD"
	DATABASE_POSTGRES_SSL_MODE Config = "DATABASE_POSTGRES_SSL_MODE"
)

// Datadog configuration
const (
	DD_SERVICE    Config = "DD_SERVICE"
	DD_ENV        Config = "DD_ENV"
	DD_VERSION    Config = "DD_VERSION"
	DD_AGENT_HOST Config = "DD_AGENT_HOST"
)

// OpenTelemetry Configuration
const (
	OTEL_EXPORTER_OTLP_ENDPOINT Config = "OTEL_EXPORTER_OTLP_ENDPOINT"
)

// Redis configuration
const (
	REDIS_ADDRESS  Config = "REDIS_ADDRESS"
	REDIS_PASSWORD Config = "REDIS_PASSWORD"
	REDIS_DB       Config = "REDIS_DB"
)

func (e Config) Get() string {
	if val, ok := env.Load(string(e)); ok {
		return val.(string)
	}
	val := os.Getenv(string(e))
	env.Store(string(e), val)
	return val
}

func Set(key, val string) {
	env.Store(string(key), val)
}

var Installations = archetypeConfiguration{
	EnableSecretManager: false,
	EnableHTTPServer:    false,
	EnableFirestore:     false,
	EnablePubSub:        false,
	EnableRedis:         false,
	EnableRestyClient:   false,
	EnablePostgreSQLDB:  false,
}

func Setup() error {
	errs := []string{}

	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found getting environments from system")
		Installations.EnableSecretManager = true
	}

	ddService := DD_SERVICE.Get()
	ddEnv := DD_ENV.Get()
	ddVersion := DD_VERSION.Get()

	if ddService != "" && ddEnv != "" && ddVersion != "" && DD_AGENT_HOST.Get() != "" && OTEL_EXPORTER_OTLP_ENDPOINT.Get() == "" {
		os.Setenv(string(OTEL_EXPORTER_OTLP_ENDPOINT), DD_AGENT_HOST.Get()+":4317")
	}

	// Check that all required environment variables are set
	requiredEnvVars := []Config{
		//PUT YOUR CUSTOM REQUIRED ENVIRONMENTS
	}

	if Installations.EnablePubSub || Installations.EnableFirestore {
		requiredEnvVars = append(requiredEnvVars, GOOGLE_PROJECT_ID)
	}

	if Installations.EnablePostgreSQLDB {
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_HOSTNAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_PORT)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_NAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_USERNAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_PASSWORD)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_SSL_MODE)
	}

	if Installations.EnableRedis {
		requiredEnvVars = append(requiredEnvVars, REDIS_ADDRESS)
		requiredEnvVars = append(requiredEnvVars, REDIS_PASSWORD)
	}

	if PORT.Get() == "" {
		Set("PORT","8080")
	}

	for _, envVar := range requiredEnvVars {
		value := envVar.Get()
		if value == "" {
			errs = append(errs, string(envVar))
		}
	}

	if len(errs) > 0 {
		slog.Error("error loading environment variables", "notFoundEnvironments", errs)
		//log.Error().Strs("notFoundEnvironments", errs).Msg("error loading environment variables")
		return fmt.Errorf("error loading environment variables: %v", errs)
	}

	return nil
}

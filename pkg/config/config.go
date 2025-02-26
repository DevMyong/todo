package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	KeyEnvironment = "ENVIRONMENT"
	KeyVersion     = "VERSION"
	KeyPort        = "PORT"

	KeyMongoURI       = "MONGO_URI"
	KeyPostgresURI    = "POSTGRES_URI"
	KeyRedisURI       = "REDIS_URI"
	KeyNatsClusterURI = "NATS_CLUSTER_URI"
	KeyAwsS3BucketURI = "AWS_S3_BUCKET_URI"

	KeyLogLevel    = "LOG_LEVEL"
	KeyLogFormat   = "LOG_FORMAT"
	KeyServiceName = "SERVICE_NAME"
)

const (
	EnvLocal      = "local"
	EnvDev        = "dev"
	EnvStage      = "stage"
	EnvProduction = "prod"

	LogFormatText = "text"
	LogFormatJSON = "json"

	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"

	ServiceNameMain = "main"
	ServiceNameUser = "user"
)

type Config struct {
	Environment    string // "prod", "stage", "dev", "local"
	MongoURI       string
	PostgresURI    string
	RedisURI       string
	NatsClusterURI string
	AwsS3BucketURI string
	Port           string

	LogLevel    string // "debug", "info", "warn", "error"
	LogFormat   string // "text", "json"
	ServiceName string // "user", "order", "payment"
	Version     string // "1.0.0"
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	environment := os.Getenv(KeyEnvironment)
	if environment == "" {
		environment = "local"
	}

	version := os.Getenv(KeyVersion)
	if version == "" {
		version = "1.0.0"
	}

	logLevel := os.Getenv(KeyLogLevel)
	if logLevel == "" {
		logLevel = "info"
	}

	logFormat := os.Getenv(KeyLogFormat)
	if logFormat == "" {
		logFormat = "json"
	}

	serviceName := os.Getenv(KeyServiceName)
	if serviceName == "" {
		serviceName = "main"
	}

	port := os.Getenv(KeyPort)
	if port == "" {
		port = "8080"
	}

	cfg := &Config{
		Environment:    environment,
		MongoURI:       os.Getenv(KeyMongoURI),
		PostgresURI:    os.Getenv(KeyPostgresURI),
		RedisURI:       os.Getenv(KeyRedisURI),
		NatsClusterURI: os.Getenv(KeyNatsClusterURI),
		AwsS3BucketURI: os.Getenv(KeyAwsS3BucketURI),
		Port:           port,
		LogLevel:       logLevel,
		LogFormat:      logFormat,
		ServiceName:    serviceName,
		Version:        version,
	}
	return cfg
}

type UserConfig struct {
	MongoURI       string
	PostgresURI    string
	RedisURI       string
	NatsClusterURI string
	AwsS3BucketURI string
}

func (c *Config) BuildUserConfig() *UserConfig {
	return &UserConfig{
		MongoURI:       c.MongoURI,
		PostgresURI:    c.PostgresURI,
		RedisURI:       c.RedisURI,
		NatsClusterURI: c.NatsClusterURI,
		AwsS3BucketURI: c.AwsS3BucketURI,
	}
}

type LoggerConfig struct {
	ServiceName string // "user", "order", "payment"...
	Level       string // "debug", "info", "warn", "error"...
	Environment string // "prod", "stage", "dev", "local"...
	Format      string // "json", "console", "text"...
	Version     string // "1.0.0"...
}

func (c *Config) BuildLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:       c.LogLevel,
		Format:      c.LogFormat,
		Version:     c.Version,
		Environment: c.Environment,
		ServiceName: c.ServiceName,
	}
}

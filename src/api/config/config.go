package config

import (
	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

// App Config
type Config struct {
	Server ServerConfig
	PgSQL  PostgreSQLConfig
	JWT    JWTConfig
	Redis  RedisConfig
}

// Fiber Config
type ServerConfig struct {
	Host string `env:"FIBER_HOST,required"`
	Port string `env:"FIBER_PORT,required"`
}

// PostgteSQL Config
type PostgreSQLConfig struct {
	Host     string `env:"PG_HOST,required"`
	Port     int    `env:"PG_PORT,required"`
	Username string `env:"PG_USER,required"`
	Password string `env:"PG_PASSWORD,required"`
	DBName   string `env:"PG_DBNAME,required"`
	SSLMode  string `env:"PG_SSLMODE,required"`
	Timezone string `env:"PG_TIMEZONE,required"`
}

// JWT Config
type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

// Redis Config
type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT,required"`
	Username string `env:"REDIS_USERNAME,required"`
	Password string `env:"REDIS_PASSWORD,required"`
	Protocol string `env:"REDIS_PROTOCOL,required"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Environments has been loaded.")

	return cfg
}

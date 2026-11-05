package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

// Config — все настройки приложения в одном месте.
// main вызывает New() один раз и передаёт cfg дальше (db, server).
type Config struct {
	HTTP     HTTP     `validate:"required"`
	Postgres Postgres `validate:"required"`
}

// HTTP — на чём слушает API.
type HTTP struct {
	Port string `env:"HTTP_PORT" env-default:"8080"`
	Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
}

// Postgres — подключение к БД.
// Для старта достаточно ConnStr; отдельные Host/User пригодятся позже.
type Postgres struct {
	ConnStr string `env:"DATABASE_URL" env-required:"true"`
}

// New читает .env (через autoload) и переменные окружения в Config.
func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

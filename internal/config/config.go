package config

import (
	"sync"
	"time"

	"github.com/aintsashqa/go-project-template/pkg/postgres"
	slogext "github.com/aintsashqa/go-project-template/pkg/slog-ext"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Logger   slogext.Config
	HTTP     HTTPConfig
	Postgres postgres.Config
}

type HTTPConfig struct {
	Port         int           `envconfig:"HTTP_PORT" default:"50051"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"1s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"1s"`
	IdleTimeout  time.Duration `envconfig:"HTTP_IDLE_TIMEOUT" default:"30s"`
}

func Get() Config {
	once.Do(func() {
		_ = godotenv.Overload()
		envconfig.MustProcess("", &conf)
	})

	return conf
}

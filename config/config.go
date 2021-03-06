package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	API struct {
		Port         string        `default:":8080" envconfig:"API_PORT"`
		ReadTimeout  time.Duration `default:"5s" envconfig:"API_READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"5s" envconfig:"API_WRITE_TIMEOUT"`
	}
}

// ParseConfig func
func ParseConfig(app string) (cfg Config, err error) {
	if err := envconfig.Process(app, &cfg); err != nil {
		if err := envconfig.Usage(app, &cfg); err != nil {
			return cfg, err
		}
		return cfg, err
	}
	return cfg, nil
}

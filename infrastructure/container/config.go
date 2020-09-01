package container

import (
	"github.com/mihnealun/prox/infrastructure/rconfig"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config contains all the variables from the .route file
type Config struct {
	Env   string `required:"true"`
	Debug bool   `required:"true"`

	LogLevel string `required:"true" envconfig:"log_level"`
	LogFile  string `required:"true" envconfig:"log_file"`

	Interface string `required:"true"`
	Port      int    `required:"true"`

	Routes rconfig.Config
}

var instanceConfig *Config
var onceConfig sync.Once

// GetConfigInstance method reads the .env file by default, validate the fields
func getConfigInstance() (*Config, error) {
	var err error
	onceConfig.Do(func() {
		err = godotenv.Load()
		if err != nil {
			return
		}

		instanceConfig = &Config{}

		err = envconfig.Process("", instanceConfig)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	instanceConfig.Routes, err = getRouteConfig()
	if err != nil {
		return nil, err
	}

	return instanceConfig, nil
}

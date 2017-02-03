package environment

import (
	"bitbucket.org/dkfbasel/scratch/src/backend/repository"
	"github.com/jinzhu/configor"
	"github.com/uber-go/zap"
)

// --- Globally available items ---

// Items is used for globally necessary items
type Items struct {
	// define a global logger
	Logger zap.Logger

	// define a global database
	SampleDB repository.SampleDBInterface

	// make the configuration available globally
	Config Configuration
}

// func withEnv(env Environment, echo.HandlerFunc)

// --- Basic application configuration ---

// Configuration is used to hold basic application configuration
type Configuration struct {
	Host       string `default:"0.0.0.0:80" env:"HOST"`
	RequestLog bool   `default:"false" env:"REQUEST_LOG"`
}

// LoadConfiguration will load the basic application configuration from the
// specified config file
func LoadConfiguration(filepath string) (Configuration, error) {
	config := Configuration{}
	err := configor.Load(&config, filepath)
	return config, err
}

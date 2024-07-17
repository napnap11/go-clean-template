package config

import (
	"fmt"
	"time"
	"strings"

	"github.com/spf13/viper"
)

type (
	// Config is the structure that holds all configuration for our application
	Config struct {
		HTTP   HTTPConfig   `mapstructure:"http"`
		Logger LoggerConfig `mapstructure:"logger"`
		// Add other configuration structs here
	}

	// HTTPConfig holds all HTTP related configuration
	HTTPConfig struct {
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
		IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	}

	// LoggerConfig holds all Logger related configuration
	LoggerConfig struct {
		Level string `mapstructure:"level"`
	}
)

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AutomaticEnv()          // read in environment variables that match

	// Replace dots with underscores in environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("No config file found. Using defaults and environment variables.")
		} else {
			// Config file was found but another error was produced
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config

	// Set default values
	setDefaults()

	// Unmarshal the configuration
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}

// setDefaults sets default values for configuration
func setDefaults() {
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.read_timeout", 10*time.Second)
	viper.SetDefault("http.write_timeout", 10*time.Second)
	viper.SetDefault("http.idle_timeout", 60*time.Second)
	viper.SetDefault("logger.level", "info")
}
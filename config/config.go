package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App                    `yaml:"app"`
		HTTP                   `yaml:"http"`
		Log                    `yaml:"logger"`
		Admin                  `yaml:"admin"`
		Schedule               `yaml:"schedule"`
		TheHillAdminProject    `yaml:"the_hill_admin"`
		TheHillCustomerProject `yaml:"the_hill_customer"`
		TheHillStoreProject    `yaml:"the_hill_store"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		ApiKey  string `env-required:"true" yaml:"api_key" env:"APP_API_KEY"`
	}

	Admin struct {
		Username string `env-required:"true" yaml:"username" env:"USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"PASSWORD"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Schedule struct {
		Period int `env-required:"true" env:"SCHEDULE_PERIOD"`
	}

	TheHillAdminProject struct {
		Name           string `env-required:"true" yaml:"name" env:"THE_HILL_ADMIN_NAME"`
		AmountToDelete int    `env-required:"true" yaml:"amount" env:"THE_HILL_ADMIN_AMOUNT_TO_DELETE"`
	}

	TheHillCustomerProject struct {
		Name           string `env-required:"true" yaml:"name" env:"THE_HILL_CUSTOMER_NAME"`
		AmountToDelete int    `env-required:"true" yaml:"amount" env:"THE_HILL_CUSTOMER_AMOUNT_TO_DELETE"`
	}

	TheHillStoreProject struct {
		Name           string `env-required:"true" yaml:"name" env:"THE_HILL_STORE_NAME"`
		AmountToDelete int    `env-required:"true" yaml:"amount" env:"THE_HILL_STORE_AMOUNT_TO_DELETE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

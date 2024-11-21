package config

import (
	"github.com/duongbui2002/core-package/config"
	"github.com/duongbui2002/core-package/config/environment"
	"strings"
)

type Config struct {
	AppOptions AppOptions `mapstructure:"appOptions" env:"AppOptions"`
}

type AppOptions struct {
	DeliveryType string `mapstructure:"deliveryType" env:"DeliveryType"`
	ServiceName  string `mapstructure:"serviceName"  env:"serviceName"`
}

func NewConfig(env environment.Environment) (*Config, error) {
	cfg, err := config.BindConfig[*Config](env)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}

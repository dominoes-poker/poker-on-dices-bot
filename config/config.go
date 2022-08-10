package config

import "os"

type Config interface {
	GetDataProviderURL() string
	GetToken() string
}

type EnvConfig struct {
	dataProviderURL string
	token           string
}

func (config EnvConfig) GetDataProviderURL() string {
	return config.dataProviderURL
}

func (config EnvConfig) GetToken() string {
	return config.token
}

func NewConfig() *EnvConfig {
	return &EnvConfig{
		dataProviderURL: os.Getenv("DATA_PROVIDER_URL"),
		token:           os.Getenv("TOKEN"),
	}
}

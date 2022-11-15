package config

type Config struct {
	Name        string `env:"APPLICATION_NAME"`
	Environment string `env:"APPLICATION_ENV"`
	ProjectID   string `env:"APPLICATION_PROJECT_ID"`
	Region      string `env:"APPLICATION_REGION"`
	Version     string `json:"APPLICATION_VERSION"`
}

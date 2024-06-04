// YOU CAN EDIT YOUR CUSTOM CONFIG HERE

package config

import (
	"go.tekoapis.com/tekone/library/conf"
	"go.tekoapis.com/tekone/library/database"
)

// Config application
type Config struct {
	conf.Base `mapstructure:",squash"`
	// Custom here
	// PostgreSQL hold the config of PostgreSQL.
	PostgreSQL database.DBConfig `json:"postgresql" mapstructure:"postgresql"`
	TMHost     string            `json:"tm_host" mapstructure:"tm_host"`
	IAMHost    string            `json:"iam_host" mapstructure:"iam_host"`
}

func loadDefaultConfig() *Config {
	c := &Config{
		Base:       *conf.DefaultBaseConfig(),
		PostgreSQL: database.PostgresSQLDefaultConfig(),
	}
	return c
}

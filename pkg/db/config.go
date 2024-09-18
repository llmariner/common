package db

import (
	"fmt"
	"os"
)

// Config specifies the configurations to connect to the database.
type Config struct {
	Host            string    `yaml:"host"`
	Port            int       `yaml:"port"`
	Username        string    `yaml:"username"`
	Database        string    `yaml:"database"`
	PasswordEnvName string    `yaml:"passwordEnvName"`
	SSL             SSLConfig `yaml:"ssl"`

	// CreateDatabase specifies whether to create the database if it does not exist.
	CreateDatabase bool `yaml:"createDatabase"`
	// OriginalDatabase specifies the original database to connect to before creating the database.
	// If empty, use "template1".
	OriginalDatabase string `yaml:"originalDatabase"`
}

// SSLConfig specifies the configurations for SSL.
type SSLConfig struct {
	Mode     string `yaml:"mode"`
	RootCert string `yaml:"rootCert"`
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port <= 0 {
		return fmt.Errorf("port must be greater than 0")
	}
	if c.Username == "" {
		return fmt.Errorf("username is required")
	}
	if c.Database == "" {
		return fmt.Errorf("database is required")
	}
	if c.PasswordEnvName == "" {
		return fmt.Errorf("passwordEnvName is required")
	}
	if c.SSL.Mode == "" {
		return fmt.Errorf("ssl.mode is required")
	}
	return nil
}

// password returns the password for the connection to the database.
func (c Config) password() string {
	return os.Getenv(c.PasswordEnvName)
}

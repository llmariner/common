package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenDB initializes GORM using the configuration parameters.
func OpenDB(c Config) (*gorm.DB, error) {
	conf := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Database, c.password(), c.SSL.Mode)
	if cert := c.SSL.RootCert; cert != "" {
		conf += fmt.Sprintf(" sslrootcert=%s", cert)
	}

	db, err := gorm.Open(postgres.Open(conf), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect to database: %s", err)
	}

	db.Logger = logger.Default

	return db, nil
}

package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenDB initializes GORM using the configuration parameters.
func OpenDB(c Config) (*gorm.DB, error) {
	if c.CreateDatabase {
		if err := createDatabaseIfNotExists(c); err != nil {
			return nil, fmt.Errorf("create database: %s", err)
		}
	}
	return open(c)
}

func createDatabaseIfNotExists(c Config) error {
	// Connect to the original database.
	copied := c
	copied.Database = c.OriginalDatabase
	if d := copied.OriginalDatabase; d != "" {
		copied.Database = d
	} else {
		copied.Database = "template1"
	}

	db, err := open(copied)
	if err != nil {
		return fmt.Errorf("open database %q: %s", c.OriginalDatabase, err)
	}
	defer func() {
		db, err := db.DB()
		if err != nil {
			log.Printf("Failed to access generic database: %s\n", err)
		}
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database: %s\n", err)
		}
	}()

	// Check if the database already exists.
	var count int
	if err := db.Raw("SELECT count(datname) FROM pg_database where datname = ?", c.Database).Scan(&count).Error; err != nil {
		return fmt.Errorf("select database: %s", err)
	}
	if count > 0 {
		log.Printf("Database %q already exists. Skipping creation.\n", c.Database)
		// Database already exists. Do nothing.
		return nil
	}

	log.Printf("Database %q does not exist. Creating.\n", c.Database)
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", c.Database)).Error; err != nil {
		return fmt.Errorf("create database: %s", err)
	}
	return nil
}

func open(c Config) (*gorm.DB, error) {
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

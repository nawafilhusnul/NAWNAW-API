package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Database        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	Driver          string
}

// getDatabaseConfig reads the database configuration from the config file using viper
// and returns a Database struct with the configuration values.
func getDatabaseConfig() *Database {
	return &Database{
		Host:            viper.GetString("database.host"),
		Port:            viper.GetString("database.port"),
		Username:        viper.GetString("database.username"),
		Password:        viper.GetString("database.password"),
		Database:        viper.GetString("database.database"),
		MaxIdleConns:    viper.GetInt("database.max_idle_conns"),
		MaxOpenConns:    viper.GetInt("database.max_open_conns"),
		ConnMaxLifetime: viper.GetDuration("database.conn_max_lifetime"),
		Driver:          viper.GetString("database.driver"),
	}
}

// NewDatabase initializes a new Database struct by reading the configuration
// from the config file and returns the struct.
func NewDatabase() *Database {
	return getDatabaseConfig()
}

// getDSN constructs the Data Source Name (DSN) string for connecting to the MySQL database
// using the configuration values stored in the Database struct.
func (d *Database) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Database)
}

// GetDB establishes a connection to the MySQL database using GORM, configures the connection pool
// settings, and returns the GORM DB object. If there is an error during the connection process,
// the function logs the error and terminates the application.
func (d *Database) GetDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(d.getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(d.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(d.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(d.ConnMaxLifetime * time.Second)

	return db
}

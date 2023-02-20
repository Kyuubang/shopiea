package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// Config is a struct to store database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	TZ       string
}

// DB is a global database connection pool
var DB *gorm.DB

func (config Config) InitDB(migrate bool) error {
	var err error

	var dsn = "host=" + config.Host +
		" user=" + config.User +
		" password=" + config.Password +
		" dbname=" + config.DBName +
		" port=" + config.Port

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if migrate {
		// Auto-migrate the database schema
		err = DB.AutoMigrate(&User{}, &Course{}, &Lab{}, &Score{})
		if err != nil {
			return err
		}
		os.Exit(0)
	}

	return nil
}

package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/aldhipradana/apidemo/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Configuration *Config
var DB *gorm.DB

type Config struct {
	// Database connection string
	Dsn      string
	DbDriver string
}

func GetConfig() *Config {
	// if not set yet, load from toml.
	if Configuration == nil {
		var config Config
		if _, err := toml.DecodeFile("config.toml", &config); err != nil {
			log.Fatal("Cannot log config.toml", err)
		}
		Configuration = &config
	}
	return Configuration
}

func GetDb() *gorm.DB {
	if DB == nil {
		config := GetConfig()

		var dialector gorm.Dialector
		switch config.DbDriver {
		case "sqlite":
			dialector = sqlite.Open(config.Dsn)
		default:
			log.Fatal("Unsupported database driver: ", config.DbDriver)
		}

		db, err := gorm.Open(dialector, &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		err = db.AutoMigrate(
			schemas.User{},
			schemas.Notes{},
		)
		if err != nil {
			log.Fatal("Cannot migrate database", err)
			panic("failed to migrate database")
		}

		DB = db
	}

	return DB
}

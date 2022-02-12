package db

import (
	"fmt"
	"log"

	"github.com/den19980107/go-fiber-gorm-starter/config"
	"github.com/den19980107/go-fiber-gorm-starter/db/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var ORM *gorm.DB

func Connect() {
	var (
		err error
		dsn string
	)

	switch config.App.Database.Driver {
	case "sqlserver":
		dsn = fmt.Sprintf(
			"%s://%s:%s@%s:%d?database=%s",
			config.App.Database.Driver,
			config.App.Database.Username,
			config.App.Database.Password,
			config.App.Database.Host,
			config.App.Database.Port,
			config.App.Database.DBName,
		)
		ORM, err = gorm.Open(sqlserver.Open(dsn))
	default:
		ORM, err = gorm.Open(sqlite.Open(config.App.Database.SQLiteFile))
	}

	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}
}

func Migrate() {
	log.Println("Initiating migration...")

	err := ORM.Migrator().AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	log.Println("Migration Completed.")
}

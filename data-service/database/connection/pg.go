package connection

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)

type PgDataBase struct {
	DB *gorm.DB
}

func (pgDb *PgDataBase) Connect(dbUrl string) error {
	var err error

	dialector := postgres.Open(dbUrl)

	if pgDb.DB, err = gorm.Open(dialector, &gorm.Config{}); err != nil {
		log.Println("Failed to connect database.")
	}

	//if err := pgDb.DB.Exec("PRAGMA foreign_keys = ON", nil).Error; err != nil {
	//	log.Println("Failed to enable foreign keys.")
	//}

	if err == nil {
		log.Println("Connection to the Database is opened.")
	}

	return err
}

func (pgDb *PgDataBase) InitializeTables(models ...interface{}) error {
	err := pgDb.DB.AutoMigrate(models...)
	if err != nil {
		log.Printf("Creating tables %s has been failed", models)
	}
	return err
}

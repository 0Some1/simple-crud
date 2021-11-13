package utilties

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simpleCrud/models"
)

var DB *gorm.DB

func Connect() {

	dnsPostgres := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT)

	database, err := gorm.Open(postgres.Open(dnsPostgres), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}
	DB = database
	migration(DB)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Comment{},
	)

}

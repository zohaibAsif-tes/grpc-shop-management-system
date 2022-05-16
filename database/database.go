package database

import (
	"fmt"

	"github.com/zohaibAsif-tes/grpc-shop-management-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const (
	// connString = "postgres://postgres:postgres@localhost:5432/shop-management-system?sslmode=disable"
	dbSource = "host=172.17.0.2 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
)

func EstablishConnection() {
	DB, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection to database failed.")
	}
	//enabling auto migration.
	DB.AutoMigrate(&models.Bill{})
}

package database

import (
	"fabiloco/hotel-trivoli-api/config"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"strconv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
  var err error

  p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

  if err != nil {
    println("Error parsing DB_PORT variable.")
  }

  dns := fmt.Sprintf(
    "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
    config.Config("DB_USER"), 
    config.Config("DB_PASSWORD"), 
    config.Config("DB_HOST"), 
    port,
    config.Config("DB_NAME"),
  )
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

  // err = DB.Migrator().DropColumn(&model.Product{}, "type")
  // if err != nil {
  //     // Do whatever you want to do!
  //     fmt.Print("ERROR: We expect the description column to be drop-able")
  // }

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.ProductType{})
	fmt.Println("Database Migrated")
}

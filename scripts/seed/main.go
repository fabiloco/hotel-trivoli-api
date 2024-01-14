package main

import (
	"fabiloco/hotel-trivoli-api/api/config"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
  db.AutoMigrate(
		&entities.Product{},
		&entities.ProductType{},
		&entities.Service{},
		&entities.Room{},
		&entities.RoomHistory{},
		&entities.Receipt{},
	)

	// Create ProductTypes
	productTypes := []entities.ProductType{
		{Name: "Electronics"},
		{Name: "Clothing"},
		{Name: "Books"},
	}

	for _, pt := range productTypes {
		db.Create(&pt)
	}

	// Create Products
	products := []entities.Product{
		{Name: "Laptop", Stock: 10, Price: 999.99, Type: []entities.ProductType{productTypes[0]}},
		{Name: "T-Shirt", Stock: 50, Price: 19.99, Type: []entities.ProductType{productTypes[1]}},
		{Name: "Book1", Stock: 30, Price: 29.99, Type: []entities.ProductType{productTypes[2]}},
		{Name: "Smartphone", Stock: 20, Price: 799.99, Type: []entities.ProductType{productTypes[0]}},
		{Name: "Jeans", Stock: 40, Price: 49.99, Type: []entities.ProductType{productTypes[1]}},
		{Name: "Book2", Stock: 25, Price: 39.99, Type: []entities.ProductType{productTypes[2]}},
		{Name: "Headphones", Stock: 15, Price: 129.99, Type: []entities.ProductType{productTypes[0]}},
		// Add more mock data as needed
	}

	for _, p := range products {
		db.Create(&p)
	}

	// Create Services
	services := []entities.Service{
		{Name: "Room Cleaning", Price: 50.0},
		{Name: "Laundry Service", Price: 30.0},
		// Add more mock data as needed
	}

	for _, s := range services {
		db.Create(&s)
	}

	// Create Rooms
	rooms := []entities.Room{
		{Number: 101},
		{Number: 102},
		{Number: 103},
		// Add more mock data as needed
	}

	for _, r := range rooms {
		db.Create(&r)
	}

	// Create RoomHistories
	roomHistories := []entities.RoomHistory{
		{
			StartDate: time.Now().Add(-24 * time.Hour),
			EndDate:   time.Now(),
			RoomID:    rooms[0].ID,
			ServiceID: services[0].ID,
		},
		{
			StartDate: time.Now().Add(-48 * time.Hour),
			EndDate:   time.Now().Add(-24 * time.Hour),
			RoomID:    rooms[1].ID,
			ServiceID: services[1].ID,
		},
		// Add more mock data as needed
	}

	for _, rh := range roomHistories {
		db.Create(&rh)
	}

	// Create Receipts
	receipts := []entities.Receipt{
		{
			TotalPrice:  150.0,
			TotalTime:   24 * time.Hour,
			Products:    []entities.Product{products[0], products[1]},
			ServiceID:   services[0].ID,
			RoomID:      rooms[0].ID,
		},
		{
			TotalPrice:  80.0,
			TotalTime:   24 * time.Hour,
			Products:    []entities.Product{products[2], products[3]},
			ServiceID:   services[1].ID,
			RoomID:      rooms[1].ID,
		},
		// Add more mock data as needed
	}

	for _, rec := range receipts {
		db.Create(&rec)
	}
}

func main() {
  var DB *gorm.DB
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

	seed(DB)
}

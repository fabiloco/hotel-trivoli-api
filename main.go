package main

import (
	"fabiloco/hotel-trivoli-api/database"
	_ "fabiloco/hotel-trivoli-api/docs"
	"fabiloco/hotel-trivoli-api/handler"
	"fabiloco/hotel-trivoli-api/middleware"
	"fabiloco/hotel-trivoli-api/model"
	"fabiloco/hotel-trivoli-api/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// Serve static assets
	app.Static("/public", "./public")

	database.ConnectDB()

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "http://example.com/swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:3001/swagger/oauth2-redirect.html",
	}))

	app.Use(middleware.FormatResponse())

	// Configurar la base de datos
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	// Migrar modelos
	db.AutoMigrate(&model.User{}) // Asegúrate de tener otros modelos aquí si es necesario

	// Crear instancias de los stores
	productStore := store.NewProductStore(database.DB)
	productTypeStore := store.NewProductTypeStore(database.DB)
	userStore := store.NewUserStore(database.DB)

	// Crear instancias de los handlers
	StoreHandler := handler.NewHandler(productStore, userStore, productTypeStore)

	// Registrar las rutas
	StoreHandler.Register(app)
	app.Listen(":3001")
}

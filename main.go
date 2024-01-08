package main

import (
	"fabiloco/hotel-trivoli-api/api/routes"
	"fabiloco/hotel-trivoli-api/database"
	_ "fabiloco/hotel-trivoli-api/docs"
	"fabiloco/hotel-trivoli-api/middleware"
	"fabiloco/hotel-trivoli-api/pkg/product"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
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

	// Crear instancias de los stores
	// productStore := store.NewProductStore(database.DB)

  productRepo := product.NewRepository(database.DB)
  productService := product.NewService(productRepo)

	api := app.Group("/api/v1", logger.New())

  routes.ProductRouter(api, productService)

	// StoreHandler.Register(app)
	app.Listen(":3001")
}

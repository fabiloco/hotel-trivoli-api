package main

import (
	"fabiloco/hotel-trivoli-api/api/routes"
	"fabiloco/hotel-trivoli-api/api/database"
	_ "fabiloco/hotel-trivoli-api/docs"
	"fabiloco/hotel-trivoli-api/api/middleware"
	"fabiloco/hotel-trivoli-api/pkg/product"
	producttype "fabiloco/hotel-trivoli-api/pkg/product_type"
	"fabiloco/hotel-trivoli-api/pkg/user"

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

  productRepo := product.NewRepository(database.DB)
  productTypeRepo := producttype.NewRepository(database.DB)
  userRepo := user.NewRepository(database.DB)

  productService := product.NewService(productRepo)
  productTypeService := producttype.NewService(productTypeRepo)
  userService := user.NewService(userRepo)

	api := app.Group("/api/v1", logger.New())

  routes.ProductRouter(api, productService)
  routes.ProductTypeRouter(api, productTypeService)
  routes.UserRouter(api, userService)

	// StoreHandler.Register(app)
	app.Listen(":3001")
}

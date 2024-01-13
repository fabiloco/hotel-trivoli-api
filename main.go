package main

import (
	"fabiloco/hotel-trivoli-api/api/config"
	"fabiloco/hotel-trivoli-api/api/database"
	"fabiloco/hotel-trivoli-api/api/routes"
	_ "fabiloco/hotel-trivoli-api/docs"
	"fabiloco/hotel-trivoli-api/pkg/product"
	productType "fabiloco/hotel-trivoli-api/pkg/product_type"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"fabiloco/hotel-trivoli-api/pkg/reports"
	"fabiloco/hotel-trivoli-api/pkg/room"
	roomHistory "fabiloco/hotel-trivoli-api/pkg/room_history"
	"fabiloco/hotel-trivoli-api/pkg/service"
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

	// app.Use(middleware.FormatResponse())

  productRepo := product.NewRepository(database.DB)
  productTypeRepo := productType.NewRepository(database.DB)
  userRepo := user.NewRepository(database.DB)
  serviceRepo := service.NewRepository(database.DB)
  roomRepo := room.NewRepository(database.DB)
  roomHistoryRepo := roomHistory.NewRepository(database.DB)
  receiptRepo := receipt.NewRepository(database.DB)

  productService := product.NewService(productRepo, productTypeRepo)
  productTypeService := productType.NewService(productTypeRepo)
  userService := user.NewService(userRepo)
  serviceService := service.NewService(serviceRepo)
  roomService := room.NewService(roomRepo)
  roomHistoryService := roomHistory.NewService(roomHistoryRepo, roomRepo, serviceRepo)
  receptService := receipt.NewService(receiptRepo, serviceRepo, productRepo, roomRepo)

  reportService := reports.NewService(productRepo, receiptRepo)

	api := app.Group("/api/v1", logger.New())

  routes.ProductRouter(api, productService)
  routes.ProductTypeRouter(api, productTypeService)
  routes.UserRouter(api, userService)
  routes.ServiceRouter(api, serviceService)
  routes.RoomRouter(api, roomService)
  routes.RoomHistoryRouter(api, roomHistoryService)
  routes.ReceiptRouter(api, receptService)

  routes.ReportsRouter(api, reportService)

	// StoreHandler.Register(app)
	app.Listen(getPort())
}

func getPort() string {
	port := config.Config("PORT")
	if port == "" {
		port = ":3001"
	} else {
		port = ":" + port
	}

	return port
}

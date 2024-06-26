package main

import (
	"fabiloco/hotel-trivoli-api/api/config"
	"fabiloco/hotel-trivoli-api/api/database"
	"fabiloco/hotel-trivoli-api/api/routes"
	_ "fabiloco/hotel-trivoli-api/docs"
	"fabiloco/hotel-trivoli-api/pkg/auth"
	individualreceipt "fabiloco/hotel-trivoli-api/pkg/individual_receipt"
	"fabiloco/hotel-trivoli-api/pkg/product"
	productType "fabiloco/hotel-trivoli-api/pkg/product_type"
	"fabiloco/hotel-trivoli-api/pkg/receipt"
	"fabiloco/hotel-trivoli-api/pkg/reports"
	"fabiloco/hotel-trivoli-api/pkg/role"
	"fabiloco/hotel-trivoli-api/pkg/room"
	roomHistory "fabiloco/hotel-trivoli-api/pkg/room_history"
	"fabiloco/hotel-trivoli-api/pkg/service"
	"fabiloco/hotel-trivoli-api/pkg/shift"
	"fabiloco/hotel-trivoli-api/pkg/user"
	"fabiloco/hotel-trivoli-api/printer"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func getPort() string {
	port := config.Config("PORT")
	if port == "" {
		port = ":3001"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		// AllowOrigins: "http://localhost:5173, http://localhost:5174, http://localhost:5175, http://localhost:4173",
		AllowOriginsFunc: func(origin string) bool { return true }, //--> this is dangerous
		AllowHeaders:     "Authorization, Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	// app.Use(cors.New())

	// Serve static assets
	app.Static("/public", "./public")

	database.ConnectDB()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"healthcheck": true,
		})
	})

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

	printer := printer.GetESCPOSPrinter()
	err := printer.InitPrinter()
	if err != nil {
		fmt.Println("Error initializing printer:", err)
	}

	productRepo := product.NewRepository(database.DB)
	productTypeRepo := productType.NewRepository(database.DB)
	userRepo := user.NewRepository(database.DB)
	serviceRepo := service.NewRepository(database.DB)
	roomRepo := room.NewRepository(database.DB)
	roomHistoryRepo := roomHistory.NewRepository(database.DB)
	receiptRepo := receipt.NewRepository(database.DB)
	individualReceiptRepo := individualreceipt.NewRepository(database.DB)
	shiftRepo := shift.NewRepository(database.DB)

	repositoryRepo := role.NewRepository(database.DB)

	productService := product.NewService(productRepo, productTypeRepo)
	productTypeService := productType.NewService(productTypeRepo)
	userService := user.NewService(userRepo)
	serviceService := service.NewService(serviceRepo)
	roomService := room.NewService(roomRepo)
	roomHistoryService := roomHistory.NewService(roomHistoryRepo, roomRepo, serviceRepo)
	receptService := receipt.NewService(receiptRepo, serviceRepo, productRepo, roomRepo, userRepo, individualReceiptRepo)

	reportService := reports.NewService(productRepo, receiptRepo, individualReceiptRepo)
	shiftService := shift.NewService(shiftRepo, receiptRepo, individualReceiptRepo)

	authService := auth.NewService(userRepo, repositoryRepo)

	api := app.Group("/api/v1", logger.New())

	routes.ProductRouter(api, productService)
	routes.ProductTypeRouter(api, productTypeService)
	routes.UserRouter(api, userService)
	routes.ServiceRouter(api, serviceService)
	routes.RoomRouter(api, roomService)
	routes.RoomHistoryRouter(api, roomHistoryService)
	routes.ReceiptRouter(api, receptService, shiftService)

	routes.ReportsRouter(api, reportService)
	routes.ShiftRouter(api, shiftService)
	routes.AuthRouter(api, authService)

	// StoreHandler.Register(app)
	app.Listen(fmt.Sprint("0.0.0.0", getPort()))
}

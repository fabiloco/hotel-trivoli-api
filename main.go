package main

import (
	"fabiloco/hotel-trivoli-api/database"
	"fabiloco/hotel-trivoli-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/swagger" // swagger handler
  _ "fabiloco/hotel-trivoli-api/docs"
)

// @title           Hotel Trivoli API
// @version         1.0
// @description     This is the awesome API for the Hotel Trivoli project. 
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  faalsaru@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
  app := fiber.New()

  app.Use(cors.New())

  // Serve static assets
  app.Static("/public", "./public")

  database.ConnectDB()

  app.Get("/docs/*", swagger.HandlerDefault) // default

	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/swagger.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3001/swagger/oauth2-redirect.html",
	}))

  router.SetupRoutes(app)

  app.Listen(":3001")
}

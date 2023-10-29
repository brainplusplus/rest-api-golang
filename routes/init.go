package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/configs"
)

var (
	log       = configs.GetLogger()
	app       *fiber.App
	apiAuth   fiber.Router
	apiAdmin  fiber.Router
	webAuth   fiber.Router
	webRouter fiber.Router
)

func SetupRoutes() {
	app = configs.GetWebApp()
	SetupApiRoutes()
	SetupWebRoutes()
}

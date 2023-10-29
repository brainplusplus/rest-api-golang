package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html/v2"
	"simple-ecommerce/responses"
	"simple-ecommerce/utils"
	"time"
)

var app *fiber.App

func InitWebApp() {
	// Create a new engine by passing the template folder
	// and template extension using <engine>.New(dir, ext string)
	engine := html.New("./views", ".html")
	for key, function := range utils.GetAllFuncMap() {
		engine.AddFunc(key, function)
	}
	app = fiber.New(fiber.Config{
		Views: engine,
	})

	//Task no 7
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			resp := new(responses.Response)
			resp.Message = "Rate Limit : Unavailable"
			return c.Status(fiber.StatusUnavailableForLegalReasons).JSON(resp)
			//return c.Status(fiber.StatusUnavailableForLegalReasons).SendFile("./views/toofast.html")
		},
	}))
	//End task no 7
}

func GetWebApp() *fiber.App {
	if app == nil {
		InitWebApp()
	}
	return app
}

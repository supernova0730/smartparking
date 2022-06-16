package transport

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"smartparking/config"
	"smartparking/internal/transport/middlewares"
	"smartparking/internal/transport/routes"
)

func ServerStart() error {
	app := initRest()
	address := config.GlobalConfig.Web.String()

	if config.GlobalConfig.Web.TLSEnable {
		return app.ListenTLS(address, config.GlobalConfig.Web.CertFile, config.GlobalConfig.Web.KeyFile)
	}
	return app.Listen(address)
}

func initRest() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	middlewares.RegisterMiddlewares(app)
	routes.RegisterRoutes(app)
	return app
}

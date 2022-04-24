package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"smartparking/config"
	_ "smartparking/docs"
	"smartparking/internal/manager"
	"smartparking/internal/transport/middlewares"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/swagger/*", swagger.HandlerDefault)

	r.Static("/images", config.GlobalConfig.Web.FileStorage)

	v1 := r.Group("/api/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	ctl := manager.MManager.Controller()

	v1.Post("/upload-image", ctl.Image().Upload)

	auth := v1.Group("/auth")
	{
		auth.Post("/register", ctl.Auth().Register)
		auth.Post("/login", ctl.Auth().Login)
		auth.Post("/refresh", ctl.Auth().RefreshTokens)
		auth.Post("/forget-password", ctl.Auth().ForgetPassword)
		auth.Post("/verify", ctl.Auth().CheckEmailVerification)
	}

	personal := v1.Group("/personal")
	personal.Use(middlewares.AuthByToken(manager.MManager))
	{
		personal.Get("/my-cars", ctl.Personal().GetAllMyCars)
		personal.Post("/add-car", ctl.Personal().AddCar)
		personal.Put("/car-activate/:id", ctl.Personal().ActivateCarByID)
		personal.Put("/car-deactivate/:id", ctl.Personal().DeactivateCarByID)
		personal.Post("/buy-ticket", ctl.Personal().BuyTicket)
		personal.Get("/my-tickets", ctl.Personal().GetAllMyTickets)
		personal.Get("/my-entry-history", ctl.Personal().GetAllMyEntryHistories)
	}

	tax := v1.Group("/tax")
	{
		tax.Get("/", ctl.Tax().GetAll)
		tax.Get("/:id", ctl.Tax().GetByID)
		tax.Post("/", ctl.Tax().Create)
		tax.Put("/:id", ctl.Tax().UpdateByID)
		tax.Delete("/:id", ctl.Tax().DeleteByID)
	}

	parkingZone := v1.Group("/parking-zone")
	{
		parkingZone.Get("/", ctl.ParkingZone().GetAll)
		parkingZone.Get("/:id", ctl.ParkingZone().GetByID)
		parkingZone.Post("/", ctl.ParkingZone().Create)
		parkingZone.Put("/:id", ctl.ParkingZone().UpdateByID)
		parkingZone.Delete("/:id", ctl.ParkingZone().DeleteByID)
	}

	parkingPlace := v1.Group("/parking-place")
	{
		parkingPlace.Get("/", ctl.ParkingPlace().GetAll)
		parkingPlace.Get("/by-parking-zone-id/:parking_zone_id", ctl.ParkingPlace().GetAllByParkingZoneID)
		parkingPlace.Get("/:id", ctl.ParkingPlace().GetByID)
		parkingPlace.Post("/", ctl.ParkingPlace().Create)
		parkingPlace.Put("/:id", ctl.ParkingPlace().UpdateByID)
		parkingPlace.Delete("/:id", ctl.ParkingPlace().DeleteByID)
	}

	client := v1.Group("/client")
	{
		client.Get("/", ctl.Client().GetAll)
		client.Get("/:id", ctl.Client().GetByID)
		client.Put("/:id", ctl.Client().UpdateByID)
		client.Delete("/:id", ctl.Client().DeleteByID)
	}

	car := v1.Group("/car")
	{
		car.Get("/", ctl.Car().GetAll)
		car.Get("/:id", ctl.Car().GetByID)
		car.Post("/", ctl.Car().Create)
		car.Put("/:id", ctl.Car().UpdateByID)
		car.Delete("/:id", ctl.Car().DeleteByID)
	}

	ticket := v1.Group("/ticket")
	{
		ticket.Get("/", ctl.Ticket().GetAll)
		ticket.Get("/:id", ctl.Ticket().GetByID)
		ticket.Put("/:id", ctl.Ticket().UpdateByID)
		ticket.Delete("/:id", ctl.Ticket().DeleteByID)
	}

	entryHistory := v1.Group("/entry-history")
	{
		entryHistory.Get("/", ctl.EntryHistory().GetAll)
		entryHistory.Get("/:id", ctl.EntryHistory().GetByID)
	}

	core := v1.Group("/core")
	{
		core.Post("/check", ctl.Core().Check)
	}
}

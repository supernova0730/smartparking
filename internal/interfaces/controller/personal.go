package controller

import "github.com/gofiber/fiber/v2"

type Personal interface {
	GetAllMyCars(c *fiber.Ctx) error
	AddCar(c *fiber.Ctx) error
	ActivateCarByID(c *fiber.Ctx) error
	DeactivateCarByID(c *fiber.Ctx) error
	GetAllMyTickets(c *fiber.Ctx) error
	BuyTicket(c *fiber.Ctx) error
	GetAllMyEntryHistories(c *fiber.Ctx) error
}

package controller

import "github.com/gofiber/fiber/v2"

type Ticket interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	UpdateByID(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
}

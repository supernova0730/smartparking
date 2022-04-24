package controller

import "github.com/gofiber/fiber/v2"

type Core interface {
	Check(c *fiber.Ctx) error
}

package controller

import "github.com/gofiber/fiber/v2"

type Image interface {
	Upload(c *fiber.Ctx) error
}

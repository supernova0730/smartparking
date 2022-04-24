package controller

import "github.com/gofiber/fiber/v2"

type EntryHistory interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
}

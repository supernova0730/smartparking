package controller

import "github.com/gofiber/fiber/v2"

type Auth interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ForgetPassword(c *fiber.Ctx) error
	RefreshTokens(c *fiber.Ctx) error
	CheckEmailVerification(c *fiber.Ctx) error
}

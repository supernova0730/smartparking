package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"smartparking/internal/apiError"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/utils"
	"smartparking/pkg/validate"
	"sync"
)

func RegisterMiddlewares(r fiber.Router) {
	r.Use(recover.New())
	r.Use(logger.New())
	r.Use(SetupContextHolder())
	r.Use(Errors())
}

func SetupContextHolder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(utils.ContextHolderKey, &sync.Map{})
		return c.Next()
	}
}

func Errors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			if apiErr := apiError.Parse(err); apiErr != nil {
				return c.Status(apiErr.Status).JSON(apiErr)
			}

			if validationErr := validate.Parse(err); validationErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(validationErr)
			}

			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return nil
	}
}

func AuthByToken(m manager.Manager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := utils.GetBearerToken(c)
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		clientID, err := m.Service().Auth().ValidateToken(token)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		utils.SetAttributeInt64(c.Context(), utils.AttributeClientID, clientID)
		return c.Next()
	}
}

func AuthByBasic() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin",
		},
	})
}

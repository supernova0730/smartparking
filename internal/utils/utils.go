package utils

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"smartparking/pkg/tools"
	"strings"
)

const (
	authPrefix = "Bearer "
)

func GetBearerToken(c *fiber.Ctx) string {
	authorization := c.Get(fiber.HeaderAuthorization)
	if !strings.HasPrefix(authorization, authPrefix) {
		return ""
	}

	return strings.TrimPrefix(authorization, authPrefix)
}

func FormFileToBuff(c *fiber.Ctx, field string) (buff *bytes.Buffer, filename string, err error) {
	formFile, err := c.FormFile(field)
	if err != nil {
		return
	}

	file, err := formFile.Open()
	if err != nil {
		return
	}

	buff, err = tools.WriteToBuff(file)
	if err != nil {
		return
	}

	filename = formFile.Filename
	return
}

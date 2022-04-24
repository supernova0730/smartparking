package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/transport/response"
	"smartparking/internal/utils"
	"smartparking/pkg/tools"
)

type CoreController struct {
	m manager.Manager
}

func NewCoreController(m manager.Manager) *CoreController {
	return &CoreController{m: m}
}

func (ctl *CoreController) Check(c *fiber.Ctx) error {
	parkingZoneID := tools.StringToInt64(c.FormValue("parking_zone_id", ""))

	buff, filename, err := utils.FormFileToBuff(c, "upload")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Core().Check(filename, buff, parkingZoneID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

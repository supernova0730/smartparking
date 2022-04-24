package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
)

type ParkingZoneController struct {
	m manager.Manager
}

func NewParkingZoneController(m manager.Manager) *ParkingZoneController {
	return &ParkingZoneController{m: m}
}

// GetAll godoc
// @Description get all parking zone
// @Tags parking zone
// @Produce json
// @Success 200 {object} []views.ParkingZoneView
// @Failure 500 {object} response.ErrorResponse
// @Router /parking-zone/ [get]
func (ctl *ParkingZoneController) GetAll(c *fiber.Ctx) error {
	parkingZones, err := ctl.m.Service().ParkingZone().GetAll()
	if err != nil {
		return err
	}

	result := models.ListParkingZone(parkingZones).ToView()
	return response.Success(c, result)
}

// GetByID godoc
// @Description get parking zone detail
// @Tags parking zone
// @Produce json
// @Param id path int true "parking zone id"
// @Success 200 {object} views.ParkingZoneView
// @Failure 400 {object} response.ErrorResponse
// @Router /parking-zone/{id} [get]
func (ctl *ParkingZoneController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	parkingZone, err := ctl.m.Service().ParkingZone().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := parkingZone.ToView()
	return response.Success(c, result)
}

// Create - ...
func (ctl *ParkingZoneController) Create(c *fiber.Ctx) error {
	var (
		createDTO dtos.ParkingZoneCreateUpdateDTO
		model     models.ParkingZone
		result    views.ParkingZoneView
	)

	err := c.BodyParser(&createDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	model.SetFromCreateUpdateDTO(createDTO)

	parkingZone, err := ctl.m.Service().ParkingZone().Create(model)
	if err != nil {
		return err
	}

	result = parkingZone.ToView()
	return response.Success(c, result)
}

// UpdateByID - ...
func (ctl *ParkingZoneController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.ParkingZoneCreateUpdateDTO
		model     models.ParkingZone
		result    views.ParkingZoneView
	)

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	model.SetFromCreateUpdateDTO(updateDTO)

	parkingZone, err := ctl.m.Service().ParkingZone().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = parkingZone.ToView()
	return response.Success(c, result)
}

// DeleteByID - ...
func (ctl *ParkingZoneController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().ParkingZone().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

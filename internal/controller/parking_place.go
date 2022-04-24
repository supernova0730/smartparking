package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
)

type ParkingPlaceController struct {
	m manager.Manager
}

func NewParkingPlaceController(m manager.Manager) *ParkingPlaceController {
	return &ParkingPlaceController{m: m}
}

// GetAll godoc
// @Description get all parking places
// @Tags parking place
// @Produce json
// @Success 200 {object} []views.ParkingPlaceListView
// @Failure 500 {object} response.ErrorResponse
// @Router /parking-place/ [get]
func (ctl *ParkingPlaceController) GetAll(c *fiber.Ctx) error {
	parkingPlaces, err := ctl.m.Service().ParkingPlace().GetAll()
	if err != nil {
		return err
	}

	result := models.ListParkingPlace(parkingPlaces).ToView()
	return response.Success(c, result)
}

// GetAllByParkingZoneID godoc
// @Description get all parking places by parking zone id
// @Tags parking place
// @Produce json
// @Param id path int true "parking zone id"
// @Success 200 {object} []views.ParkingPlaceListView
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /parking-place/by-parking-zone-id/{id} [get]
func (ctl *ParkingPlaceController) GetAllByParkingZoneID(c *fiber.Ctx) error {
	parkingZoneID, err := c.ParamsInt("parking_zone_id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	parkingPlaces, err := ctl.m.Service().ParkingPlace().GetAllBy(models.ParkingPlace{ParkingZoneID: int64(parkingZoneID)})
	if err != nil {
		return err
	}

	result := models.ListParkingPlace(parkingPlaces).ToView()
	return response.Success(c, result)
}

// GetByID godoc
// @Description get parking place detail
// @Tags parking place
// @Produce json
// @Param id path int true "parking place id"
// @Success 200 {object} views.ParkingPlaceDetailView
// @Failure 400 {object} response.ErrorResponse
// @Router /parking-place/{id} [get]
func (ctl *ParkingPlaceController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	parkingPlace, err := ctl.m.Service().ParkingPlace().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := parkingPlace.ToView()
	return response.Success(c, result)
}

// Create - ...
func (ctl *ParkingPlaceController) Create(c *fiber.Ctx) error {
	var (
		createDTO dtos.ParkingPlaceCreateDTO
		model     models.ParkingPlace
		result    views.ParkingPlaceDetailView
	)

	err := c.BodyParser(&createDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	model.SetFromCreateDTO(createDTO)
	parkingPlace, err := ctl.m.Service().ParkingPlace().Create(model)
	if err != nil {
		return err
	}

	result = parkingPlace.ToView()
	return response.Success(c, result)
}

// UpdateByID - ...
func (ctl *ParkingPlaceController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.ParkingPlaceUpdateDTO
		model     models.ParkingPlace
		result    views.ParkingPlaceDetailView
	)

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	model.SetFromUpdateDTO(updateDTO)
	parkingPlace, err := ctl.m.Service().ParkingPlace().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = parkingPlace.ToView()
	return response.Success(c, result)
}

// DeleteByID - ...
func (ctl *ParkingPlaceController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().ParkingPlace().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
)

type CarController struct {
	m manager.Manager
}

func NewCarController(m manager.Manager) *CarController {
	return &CarController{m: m}
}

func (ctl *CarController) GetAll(c *fiber.Ctx) error {
	cars, err := ctl.m.Service().Car().GetAll()
	if err != nil {
		return err
	}

	result := models.ListCar(cars).ToView()
	return response.Success(c, result)
}

func (ctl *CarController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	car, err := ctl.m.Service().Car().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := car.ToView()
	return response.Success(c, result)
}

func (ctl *CarController) Create(c *fiber.Ctx) error {
	var (
		createDTO dtos.CarCreateUpdateDTO
		model     models.Car
		result    views.CarDetailView
	)

	err := c.BodyParser(&createDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	model.SetFromCreateUpdateDTO(createDTO)
	car, err := ctl.m.Service().Car().Create(model)
	if err != nil {
		return err
	}

	result = car.ToView()
	return response.Success(c, result)
}

func (ctl *CarController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.CarCreateUpdateDTO
		model     models.Car
		result    views.CarDetailView
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
	car, err := ctl.m.Service().Car().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = car.ToView()
	return response.Success(c, result)
}

func (ctl *CarController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Car().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

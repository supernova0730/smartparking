package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
	"smartparking/pkg/validate"
)

type ClientController struct {
	m manager.Manager
}

func NewClientController(m manager.Manager) *ClientController {
	return &ClientController{m: m}
}

func (ctl *ClientController) GetAll(c *fiber.Ctx) error {
	clients, err := ctl.m.Service().Client().GetAll()
	if err != nil {
		return err
	}

	result := clients.ToView()
	return response.Success(c, result)
}

func (ctl *ClientController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	client, err := ctl.m.Service().Client().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := client.ToView()
	return response.Success(c, result)
}

func (ctl *ClientController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.ClientUpdateDTO
		model     models.Client
		result    views.ClientView
	)

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = validate.Struct(updateDTO)
	if err != nil {
		return err
	}

	model.SetFromUpdateDTO(updateDTO)
	client, err := ctl.m.Service().Client().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = client.ToView()
	return response.Success(c, result)
}

func (ctl *ClientController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Client().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

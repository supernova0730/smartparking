package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
)

type TicketController struct {
	m manager.Manager
}

func NewTicketController(m manager.Manager) *TicketController {
	return &TicketController{m: m}
}

func (ctl *TicketController) GetAll(c *fiber.Ctx) error {
	tickets, err := ctl.m.Service().Ticket().GetAll()
	if err != nil {
		return err
	}

	result := models.ListTicket(tickets).ToView()
	return response.Success(c, result)
}

func (ctl *TicketController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	ticket, err := ctl.m.Service().Ticket().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := ticket.ToView()
	return response.Success(c, result)
}

func (ctl *TicketController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.TicketUpdateDTO
		model     models.Ticket
		result    views.TicketDetailView
	)

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = model.SetFromUpdateDTO(updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	ticket, err := ctl.m.Service().Ticket().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = ticket.ToView()
	return response.Success(c, result)
}

func (ctl *TicketController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Ticket().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

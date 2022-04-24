package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
)

type EntryHistoryController struct {
	m manager.Manager
}

func NewEntryHistoryController(m manager.Manager) *EntryHistoryController {
	return &EntryHistoryController{m: m}
}

func (ctl *EntryHistoryController) GetAll(c *fiber.Ctx) error {
	entryHistories, err := ctl.m.Service().EntryHistory().GetAll()
	if err != nil {
		return err
	}

	result := models.ListEntryHistory(entryHistories).ToView()
	return response.Success(c, result)
}

func (ctl *EntryHistoryController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	entryHistory, err := ctl.m.Service().EntryHistory().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := entryHistory.ToView()
	return response.Success(c, result)
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/views"
)

type TaxController struct {
	m manager.Manager
}

func NewTaxController(m manager.Manager) *TaxController {
	return &TaxController{m: m}
}

// GetByID godoc
// @Description get tax detail
// @Tags tax
// @Produce json
// @Param id path int true "tax id"
// @Success 200 {object} views.TaxView
// @Failure 400 {object} response.ErrorResponse
// @Router /tax/{id} [get]
func (ctl *TaxController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	tax, err := ctl.m.Service().Tax().GetByID(int64(id))
	if err != nil {
		return err
	}

	result := tax.ToView()
	return response.Success(c, result)
}

// GetAll godoc
// @Description get all taxes
// @Tags tax
// @Produce json
// @Success 200 {object} []views.TaxView
// @Failure 500 {object} response.ErrorResponse
// @Router /tax/ [get]
func (ctl *TaxController) GetAll(c *fiber.Ctx) error {
	taxes, err := ctl.m.Service().Tax().GetAllOrderByPrice()
	if err != nil {
		return err
	}

	result := taxes.ToView()
	return response.Success(c, result)
}

// Create - ...
func (ctl *TaxController) Create(c *fiber.Ctx) error {
	var (
		createDTO dtos.TaxCreateUpdateDTO
		model     models.Tax
		result    views.TaxView
	)

	err := c.BodyParser(&createDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = model.SetFromCreateUpdateDTO(createDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	tax, err := ctl.m.Service().Tax().Create(model)
	if err != nil {
		return err
	}

	result = tax.ToView()
	return response.Success(c, result)
}

// UpdateByID - ...
func (ctl *TaxController) UpdateByID(c *fiber.Ctx) error {
	var (
		updateDTO dtos.TaxCreateUpdateDTO
		model     models.Tax
		result    views.TaxView
	)

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = c.BodyParser(&updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = model.SetFromCreateUpdateDTO(updateDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	tax, err := ctl.m.Service().Tax().Update(int64(id), model)
	if err != nil {
		return err
	}

	result = tax.ToView()
	return response.Success(c, result)
}

// DeleteByID - ...
func (ctl *TaxController) DeleteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Tax().Delete(int64(id))
	if err != nil {
		return err
	}

	return response.Success(c, nil)
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/internal/transport/response"
	"smartparking/internal/utils"
	"smartparking/internal/views"
	"smartparking/pkg/validate"
)

type PersonalController struct {
	m manager.Manager
}

func NewPersonalController(m manager.Manager) *PersonalController {
	return &PersonalController{m: m}
}

// GetAllMyCars godoc
// @Description get client cars
// @Tags personal
// @Produce json
// @Success 200 {object} []views.CarListView
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/my-cars [get]
// @Security ApiKeyAuth
func (ctl *PersonalController) GetAllMyCars(c *fiber.Ctx) error {
	clientId := utils.GetClientID(c.Context())

	cars, err := ctl.m.Service().Car().GetAllByClientID(clientId)
	if err != nil {
		return err
	}

	result := cars.ToView()
	return response.Success(c, result)
}

// AddCar godoc
// @Description add car by client
// @Tags personal
// @Accept json
// @Produce json
// @Param createDTO body dtos.PersonalCarCreateDTO true "PersonalCarCreateDTO"
// @Success 200 {object} views.CarDetailView
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/add-car [post]
// @Security ApiKeyAuth
func (ctl *PersonalController) AddCar(c *fiber.Ctx) error {
	var (
		createDTO dtos.PersonalCarCreateDTO
		model     models.Car
		result    views.CarDetailView
	)

	clientID := utils.GetClientID(c.Context())

	if err := c.BodyParser(&createDTO); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	if err := validate.Struct(createDTO); err != nil {
		return err
	}

	model.SetFromPersonalCarCreateDTO(createDTO)
	model.ClientID = clientID
	car, err := ctl.m.Service().Car().Create(model)
	if err != nil {
		return err
	}

	result = car.ToView()
	return response.Success(c, result)
}

// ActivateCarByID godoc
// @Description activate car by client
// @Tags personal
// @Param id path int true "car id"
// @Success 200
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/car-activate/{id} [put]
// @Security ApiKeyAuth
func (ctl *PersonalController) ActivateCarByID(c *fiber.Ctx) error {
	clientID := utils.GetClientID(c.Context())

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Car().SetStatusByClientAndCarID(clientID, int64(id), true)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

// DeactivateCarByID godoc
// @Description deactivate car by client
// @Tags personal
// @Param id path int true "car id"
// @Success 200
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/car-deactivate/{id} [put]
// @Security ApiKeyAuth
func (ctl *PersonalController) DeactivateCarByID(c *fiber.Ctx) error {
	clientID := utils.GetClientID(c.Context())

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = ctl.m.Service().Car().SetStatusByClientAndCarID(clientID, int64(id), false)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

// GetAllMyTickets godoc
// @Description get client tickets
// @Tags personal
// @Produce json
// @Success 200 {object} []views.PersonalTicketListView
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/my-tickets [get]
// @Security ApiKeyAuth
func (ctl *PersonalController) GetAllMyTickets(c *fiber.Ctx) error {
	clientID := utils.GetClientID(c.Context())
	tickets, err := ctl.m.Service().Ticket().GetAllBy(models.Ticket{ClientID: clientID})
	if err != nil {
		return err
	}

	result := tickets.ToPersonalTicketListView()
	return response.Success(c, result)
}

// BuyTicket godoc
// @Description buy ticket by client
// @Tags personal
// @Accept json
// @Produce json
// @Param buyTicketDTO body dtos.BuyTicketDTO true "BuyTicketDTO"
// @Success 200 {object} views.TicketDetailView
// @Failure 400 {object} validate.ValidationError
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/buy-ticket [post]
// @Security ApiKeyAuth
func (ctl *PersonalController) BuyTicket(c *fiber.Ctx) error {
	var (
		buyTicketDTO dtos.BuyTicketDTO
		result       views.TicketDetailView
	)

	clientID := utils.GetClientID(c.Context())

	err := c.BodyParser(&buyTicketDTO)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	err = validate.Struct(buyTicketDTO)
	if err != nil {
		return err
	}

	ticket, err := ctl.m.Service().Ticket().BuyTicket(clientID, buyTicketDTO)
	if err != nil {
		return err
	}

	result = ticket.ToView()
	return response.Success(c, result)
}

// GetAllMyEntryHistories godoc
// @Description get client entry histories
// @Tags personal
// @Produce json
// @Param car_id query int false "car id"
// @Param selected_period query string false "selected period (ex: day, week, month, all)"
// @Param date_from query string false "date from (ex: 2022-04-02T00:00:00.000Z)"
// @Param date_to query string false "date to (ex: 2022-04-02T00:00:00.000Z)"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} []views.PersonalEntryHistoryListView
// @Failure 500 {object} response.ErrorResponse
// @Router /personal/my-entry-history [get]
// @Security ApiKeyAuth
func (ctl *PersonalController) GetAllMyEntryHistories(c *fiber.Ctx) error {
	clientID := utils.GetClientID(c.Context())

	filter := dtos.EntryHistoryFilter{}
	err := filter.Fill(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	entryHistories, total, err := ctl.m.Service().EntryHistory().GetAllByClientIDAndFilter(clientID, filter)
	if err != nil {
		return err
	}

	result := entryHistories.ToPersonalEntryHistoryView()
	return response.Success(c, fiber.Map{
		"rows":  result,
		"total": total,
	})
}

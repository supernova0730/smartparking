package apiError

import "net/http"

const (
	TicketExpired  = "TICKET_EXPIRED"
	TicketNotFound = "TICKET_NOT_FOUND"
)

var ticketErrors = []*apiError{
	{
		ID:      TicketExpired,
		Message: "ticket was expired",
		Status:  http.StatusBadRequest,
	},
	{
		ID:      TicketNotFound,
		Message: "ticket not found",
		Status:  http.StatusNotFound,
	},
}

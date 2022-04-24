package apiError

import "net/http"

const (
	EntryHistoryNotFound = "ENTRY_HISTORY_NOT_FOUND"
)

var entryHistoryErrors = []*apiError{
	{
		ID:      EntryHistoryNotFound,
		Message: "entry history not found",
		Status:  http.StatusNotFound,
	},
}

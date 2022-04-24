package apiError

type apiError struct {
	ID      string `json:"-"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"-"`
}

func (e *apiError) Error() string {
	return e.Message
}

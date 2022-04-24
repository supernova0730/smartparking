package apiError

var apiErrors []*apiError

func init() {
	apiErrors = append(apiErrors, authErrors...)
	apiErrors = append(apiErrors, carErrors...)
	apiErrors = append(apiErrors, clientErrors...)
	apiErrors = append(apiErrors, emailVerificationErrors...)
	apiErrors = append(apiErrors, entryHistoryErrors...)
	apiErrors = append(apiErrors, parkingPlaceErrors...)
	apiErrors = append(apiErrors, parkingZoneErrors...)
	apiErrors = append(apiErrors, taxErrors...)
	apiErrors = append(apiErrors, ticketErrors...)
}

func findByID(id string) *apiError {
	for _, err := range apiErrors {
		if err.ID == id {
			return err
		}
	}
	panic("apiError doesn't exist: " + id)
}

func Throw(id string) *apiError {
	return findByID(id)
}

func Parse(err error) *apiError {
	apiErr, ok := err.(*apiError)
	if !ok {
		return nil
	}
	return apiErr
}

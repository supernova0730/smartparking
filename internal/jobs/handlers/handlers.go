package handlers

import (
	"smartparking/internal/interfaces/jobs"
	"smartparking/internal/interfaces/manager"
)

const (
	ECHO                   = "ECHO"
	EXPIRED_PARKING_PLACES = "EXPIRED_PARKING_PLACES"
)

func RegisterHandlers(m manager.Manager) (handlers []jobs.Job, err error) {
	allJobs, err := m.Repository().Job().GetAll()
	if err != nil {
		return
	}

	for _, j := range allJobs {
		switch j.Code {
		case ECHO:
			handlers = append(handlers, NewEcho(m, j))
		case EXPIRED_PARKING_PLACES:
			handlers = append(handlers, NewExpiredParkingPlaces(m, j))
		}
	}

	return
}

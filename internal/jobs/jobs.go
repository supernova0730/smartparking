package jobs

import (
	"github.com/go-co-op/gocron"
	"smartparking/internal/jobs/handlers"
	"smartparking/internal/manager"
	"time"
)

func Start() error {
	s := gocron.NewScheduler(time.UTC)

	jobHandlers, err := handlers.RegisterHandlers(manager.MManager)
	if err != nil {
		return err
	}

	for _, handler := range jobHandlers {
		_, err = s.Cron(handler.GetSchedule()).Do(handler.Do)
		if err != nil {
			return err
		}
	}

	s.StartAsync()
	return nil
}

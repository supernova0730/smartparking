package dtos

import (
	"github.com/gofiber/fiber/v2"
	"smartparking/internal/consts"
	"smartparking/pkg/tools"
	"time"
)

type EntryHistoryFilter struct {
	CarID          int64     `query:"car_id"`
	ParkingZoneID  int64     `query:"parking_zone_id"`
	SelectedPeriod string    `query:"selected_period"`
	DateFrom       time.Time `query:"date_from"`
	DateTo         time.Time `query:"date_to"`
	Page           int       `query:"page"`
	Size           int       `query:"size"`
}

func (f *EntryHistoryFilter) Fill(c *fiber.Ctx) error {
	if err := c.QueryParser(f); err != nil {
		return err
	}

	if f.Size == 0 {
		f.Size = consts.DefaultSize
	}

	f.DateFrom, f.DateTo = tools.GetDatePeriod(f.SelectedPeriod, f.DateFrom, f.DateTo)

	return nil
}

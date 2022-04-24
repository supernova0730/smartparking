package models

import (
	"smartparking/internal/dtos"
	"smartparking/internal/views"
	"smartparking/pkg/tools"
	"time"
)

type ListTax []Tax
type Tax struct {
	ID       int64         `gorm:"column:id"`
	Title    string        `gorm:"column:title"`
	Duration time.Duration `gorm:"column:duration"`
	Price    int           `gorm:"column:price"`
}

func (Tax) TableName() string {
	return "tax"
}

func (t Tax) ToView() views.TaxView {
	return views.TaxView{
		ID:       t.ID,
		Title:    t.Title,
		Duration: tools.DurationToString(t.Duration),
		Price:    t.Price,
	}
}

func (list ListTax) ToView() (result []views.TaxView) {
	for _, tax := range list {
		result = append(result, tax.ToView())
	}
	return
}

func (t *Tax) SetFromCreateUpdateDTO(dto dtos.TaxCreateUpdateDTO) error {
	duration, err := tools.ParseDuration(dto.Duration)
	if err != nil {
		return err
	}

	t.Title = dto.Title
	t.Duration = duration
	t.Price = dto.Price
	return nil
}

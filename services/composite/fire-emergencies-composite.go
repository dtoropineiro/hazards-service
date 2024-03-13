package composite

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
	"hazards-emergencies-go/dto"
	"time"
)

type FireEmergencyClientService interface {
	GetAllFires() ([]dto.FireDto, error)
}

type FireEmergencyServiceComposite struct {
	fireEmergencyClientServices []FireEmergencyClientService
}

func NewCompositeFireEmergencyService(central123, cvb FireEmergencyClientService) *FireEmergencyServiceComposite {
	return &FireEmergencyServiceComposite{
		fireEmergencyClientServices: []FireEmergencyClientService{central123, cvb},
	}
}

func (c *FireEmergencyServiceComposite) GetAllFires() ([]dto.FireDto, error) {
	var allFires []dto.FireDto
	lastFires := make([]dto.FireDto, 0)
	for _, service := range c.fireEmergencyClientServices {
		fires, err := service.GetAllFires()
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		allFires = append(allFires, fires...)
		for _, fire := range allFires {
			isDateFromLastDay, err := isWithinLast24Hours(fire.DateTime)
			if err != nil {
				logrus.Error(err)
				return nil, err
			}
			if isDateFromLastDay {
				lastFires = append(lastFires, fire)
			}
		}
	}
	return lastFires, nil
}

func isWithinLast24Hours(dateStr string) (bool, error) {
	inputDateTime, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return false, fmt.Errorf("failed to parse date: %w", err)
	}
	now := time.Now()
	duration := now.Sub(inputDateTime)
	return duration <= 24*time.Hour, nil
}

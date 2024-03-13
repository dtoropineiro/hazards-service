package composite

import (
	"github.com/sirupsen/logrus"
	"hazards-emergencies-go/clients"
	"hazards-emergencies-go/dto"
	"hazards-emergencies-go/utils"
	"hazards-emergencies-go/utils/datesutil"
)

type Central123FireEmergenciesCompositeService struct{}

func (s *Central123FireEmergenciesCompositeService) GetAllFires() ([]dto.FireDto, error) {
	logrus.Info("In Central123FireEmergenciesCompositeService GetAllFires")

	fires, err := clients.GetLastFires()
	if err != nil {
		logrus.Info("", err)
	}
	features := fires.Features
	if features == nil {
		return []dto.FireDto{}, nil
	}
	return mapFireResponseDto(features)

}

func mapFireResponseDto(features []clients.Feature) ([]dto.FireDto, error) {
	var fireDtos []dto.FireDto
	for _, feature := range features {
		longitude := feature.Geometry.Coordinates[0]
		latitude := feature.Geometry.Coordinates[1]
		id := feature.Properties.Id

		district := feature.Properties.Comuna
		key := feature.Properties.Clave
		place := feature.Properties.Ubicacion
		dateTime := feature.Properties.Fecha
		date, time := datesutil.GetDateAndTimeSplit(dateTime)
		keyDescription := utils.GetKeyDescriptionByKey(key)
		fireId := utils.GenerateFireId(place, latitude, longitude, key, dateTime)

		fireDto := dto.FireDto{
			Id:             id,
			FireId:         fireId,
			Key:            key,
			KeyDescription: keyDescription,
			Longitude:      longitude,
			Latitude:       latitude,
			PlaceName:      place,
			DateTime:       dateTime,
			Date:           date,
			Time:           time,
			District:       district,
		}

		fireDtos = append(fireDtos, fireDto)
	}
	return fireDtos, nil
}

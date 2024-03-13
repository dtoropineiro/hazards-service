package composite

import (
	"github.com/sirupsen/logrus"
	"hazards-emergencies-go/clients"
	"hazards-emergencies-go/dto"
	"hazards-emergencies-go/utils"
	"hazards-emergencies-go/utils/datesutil"
	"strconv"
)

type CvbFireEmergenciesCompositeService struct{}

func (s *CvbFireEmergenciesCompositeService) GetAllFires() ([]dto.FireDto, error) {
	fireEmergencyResponse, err := clients.GetLastFireEmergencies()
	fireEmergency := fireEmergencyResponse.Data

	if err != nil {
		logrus.Error("Error: ", err)
	}

	logrus.Println("In CvbFireEmergenciesCompositeService GetAllFires")
	return mapEmergencySliceToFireDto(fireEmergency)
}

func mapEmergencySliceToFireDto(emergencies []clients.Emergency) ([]dto.FireDto, error) {
	var fireDtos []dto.FireDto
	for _, emergency := range emergencies {
		longitude := emergency.Lat
		latitude := emergency.Lng
		district := "Región de Valparaíso"
		key := emergency.Sigla
		place := emergency.Dir
		dateTime, err := datesutil.ConvertDateFormat(emergency.Fecha)
		if err != nil {
			logrus.Error("Error parsing datetime.", err)
			return nil, err
		}
		date, time := datesutil.GetDateAndTimeSplit(dateTime)
		keyDescription := utils.GetKeyDescriptionByKey(key)
		latitudeFloat, err := strconv.ParseFloat(latitude, 64)
		longitudeFloat, err := strconv.ParseFloat(longitude, 64)
		if err != nil {
			logrus.Error("Error parsing string to float64.", err)
			return nil, err
		}
		fireId := utils.GenerateFireId(place, latitudeFloat, longitudeFloat, key, dateTime)

		fireDto := dto.FireDto{
			FireId:         fireId,
			Key:            key,
			KeyDescription: keyDescription,
			Longitude:      latitudeFloat,
			Latitude:       longitudeFloat,
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

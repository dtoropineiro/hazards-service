package services

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hazards-emergencies-go/clients"
	"hazards-emergencies-go/dto"
	"hazards-emergencies-go/handlers"
	"hazards-emergencies-go/utils/datesutil"
	"regexp"
	"strconv"
	"strings"
)

func GetEarthquakes() dto.EarthquakeResponseDto {
	earthquakesSismologia, err := clients.GetEarthquakes()
	if err != nil {
		logrus.Error(err)
	}
	earthquakes := make([]dto.EarthquakeDto, 0)
	for _, earthquakesSismologia := range earthquakesSismologia {

		dateAndTimeChile, place := getDateAndPlace(earthquakesSismologia.Location)
		latitudeStr, longitudeStr := getLatitudeAndLongitude(earthquakesSismologia.LatLong)
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		magnitude := earthquakesSismologia.Magnitude
		depth := earthquakesSismologia.Depth
		earthquakeDate, earthquakeTime := datesutil.GetDateAndTimeSplit(dateAndTimeChile)
		annotationSize := handlers.GetAnnotationSize(earthquakesSismologia.Magnitude)
		earthquakeId := generateEarthquakeId(place, latitude, longitude, depth, magnitude, dateAndTimeChile)
		if err != nil {
			logrus.Error("Error parsing string to float64:", err)
		}
		logrus.Info(earthquakesSismologia)

		earthquake := dto.EarthquakeDto{
			EarthquakeID:   earthquakeId,
			Magnitude:      magnitude,
			Longitude:      longitude,
			Latitude:       latitude,
			Depth:          depth,
			PlaceName:      place,
			DateTime:       dateAndTimeChile,
			Date:           earthquakeDate,
			Time:           earthquakeTime,
			AnnotationSize: annotationSize,
		}
		earthquakes = append(earthquakes, earthquake)
	}
	return dto.EarthquakeResponseDto{
		EarthquakesDto: earthquakes,
	}
}

func getDateAndPlace(dateTimeAndPlace string) (string, string) {
	re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})(\d+ .*)`)
	matches := re.FindStringSubmatch(dateTimeAndPlace)
	if matches != nil {
		timestamp := matches[1]
		location := matches[2]
		return timestamp, location
	} else {
		logrus.Error("Could not split dateTimeAndPlace", dateTimeAndPlace)
		return "", ""
	}
}

func getLatitudeAndLongitude(latitudeAndLongitude string) (string, string) {
	latitudeAndLongitudeSlice := strings.Split(latitudeAndLongitude, " ")
	return latitudeAndLongitudeSlice[0], latitudeAndLongitudeSlice[1]
}

func generateEarthquakeId(
	placeName string,
	latitude float64,
	longitude float64,
	depth string,
	magnitude string,
	dateTime string) string {

	stringEarthquake := placeName + strconv.FormatFloat(latitude, 'f', -1, 64) +
		strconv.FormatFloat(longitude, 'f', -1, 64) + depth + magnitude + dateTime

	earthquakeUUID := uuid.NewMD5(uuid.Nil, []byte(stringEarthquake))

	return earthquakeUUID.String()
}

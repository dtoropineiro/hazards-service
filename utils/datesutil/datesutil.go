package datesutil

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

func ConvertTimestampToIso8601(timestamp int64) string {
	instant := time.Unix(0, timestamp*int64(time.Millisecond))
	return instant.Format(time.RFC3339)
}

func GetDateTimeAndPlaceSplit(dateTimeAndPlace string) (string, string) {
	pattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (.*)`)
	matches := pattern.FindStringSubmatch(dateTimeAndPlace)

	if len(matches) != 3 {
		logrus.Info("Error in date time and place string")
		return "", ""
	}

	return matches[1], strings.TrimSpace(matches[2])
}

func GetDateAndTimeSplit(dateTime string) (string, string) {
	parts := strings.Fields(dateTime)
	date := parts[0]
	incidentTime := parts[len(parts)-1]

	regex := "^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/((\\d{4}))$"

	if matched, _ := regexp.MatchString(regex, date); matched {
		inputFormat := "02/01/2006"
		parsedDate, _ := time.Parse(inputFormat, date)
		outputFormat := "2006-01-02"
		formattedDate := parsedDate.Format(outputFormat)
		return formattedDate, incidentTime
	}

	return date, incidentTime
}

func ConvertDateFormat(dateStr string) (string, error) {
	inputDateTime, err := time.Parse("02/01/2006 - 15:04:05", dateStr)
	if err != nil {
		return "", err
	}
	formattedDate := inputDateTime.Format("2006-01-02 15:04:05")
	return formattedDate, nil
}

package utils

import (
	"fmt"
	"github.com/google/uuid"
	"hazards-emergencies-go/constants"
	"strings"
	"unicode"
)

func GetKeyDescriptionByKey(key string) string {
	keyFormatted := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(key), "clave"))
	keyDescription, ok := constants.GetFireMapCodes()[keyFormatted]

	if !ok {
		return ""
	}

	runes := []rune(keyDescription)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func GenerateFireId(placeName string, latitude float64, longitude float64, key string, dateTime string) string {
	stringFire := fmt.Sprintf("%s%f%f%s%s", placeName, latitude, longitude, key, dateTime)
	md5 := uuid.NewMD5(uuid.Nil, []byte(stringFire))

	return strings.ReplaceAll(md5.String(), "-", "")
}

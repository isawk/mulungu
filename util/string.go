package util

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

//StringDecode decodes base64 encoded string content
func StringDecode(encodedString string) ([]byte, error) {
	decoded, decodedError := base64.StdEncoding.DecodeString(encodedString)
	if decodedError != nil {
		return nil, decodedError
	}
	return decoded, nil
}

//StringContains checks if a string contains any of the pradicates send
func StringContains(subject string, pradicates []string) bool {
	for _, pradicate := range pradicates {
		if strings.Contains(subject, pradicate) {
			return true
		}
	}
	return false
}

//StringToInt converts a string to an integer
func StringToInt(subject string) int {
	if subject != "" {
		i, err := strconv.Atoi(subject)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

//StringToInt64 converts a string to an integer
func StringToInt64(subject string) int64 {
	if subject != "" {
		i, err := strconv.ParseInt(subject, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

//StringToFloat64 converts a string to a floating point 64 bit
func StringToFloat64(subject string) float64 {
	if subject != "" {
		i, err := strconv.ParseFloat(subject, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

//StringToFloat32 converts a string to a floating point 32 bit
func StringToFloat32(subject string) float64 {
	if subject != "" {
		i, err := strconv.ParseFloat(subject, 32)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

//ToString Converts interface to a string
func ToString(subject interface{}) string {

	if subject == nil {
		return ""
	}

	switch ReflectKind(subject) {
	case reflect.String:
		return subject.(string)
	case reflect.Int:
		return strconv.Itoa(subject.(int))
	case reflect.Int64:
		return strconv.FormatInt(subject.(int64), 10)
	case reflect.Float64:
		return strconv.FormatFloat(subject.(float64), 'E', -1, 64)
	}

	return ""
}

//NumberizeString converts string to either int or float
func NumberizeString(subject interface{}) interface{} {
	if govalidator.IsInt(subject.(string)) {
		return StringToInt(subject.(string))
	} else if govalidator.IsFloat(subject.(string)) {
		return StringToFloat64(subject.(string))
	}
	return subject
}

//StringTobyte converts string t byte
func StringTobyte(subject string) []byte {
	return []byte(subject)
}

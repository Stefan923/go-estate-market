package util

import (
	"encoding/json"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ConvertTo[T any](object any) (*T, error) {
	objectJson, err := json.Marshal(&object)
	if err != nil {
		return nil, err
	}

	var convertedObject T
	err = json.Unmarshal(objectJson, &convertedObject)
	if err != nil {
		return nil, err
	}

	return &convertedObject, nil
}

func ConvertToSnakeCaseMap(object any) (map[string]interface{}, error) {
	objectMap, err := ConvertTo[map[string]interface{}](object)
	if err != nil {
		return nil, err
	}

	convertedObjectMap := map[string]interface{}{}
	for key, value := range *objectMap {
		convertedObjectMap[convertToSnakeCase(key)] = value
	}

	return convertedObjectMap, nil
}

func convertToSnakeCase(string string) string {
	snakeString := matchFirstCap.ReplaceAllString(string, "${1}_${2}")
	snakeString = matchAllCap.ReplaceAllString(snakeString, "${1}_${2}")
	return strings.ToLower(snakeString)
}

package assert

import (
	"net/http"
)

func Equal(expected interface{}, actual interface{}) bool {
	return expected == actual
}

func NotEqual(expected interface{}, actual interface{}) bool {
	return expected != actual
}

func IsNil(actual interface{}) bool {
	return actual == nil
}

/*
Assertion for compare response HTTP code with the expected code
*/
func HTTPCodeEqual(response http.Response, expected int) bool {
	return response.StatusCode == expected
}

/*
Assertion fo check if the property name his is on JSON
*/
func HasJSONProperty(json map[string]interface{}, propertyName string) bool {
	_, ok := json[propertyName]
	return ok
}

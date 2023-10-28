package Request

import (
	"fmt"
	"net/http"
	"strings"

	Console "api_tester/Class/Console"
)

/*
Wrapper to send a HTTP request with the given Method
*/
func Send(url string, method string) *http.Response {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		return Get_Request(url)
	default:
		return Get_Request(url)
	}
}

func Get_Request(url string) *http.Response {
	// Create HTTP client
	client := &http.Client{}

	// Request execution
	response, err := client.Get(url)
	if err != nil {
		Console.Error("The HTTP request failed")
		fmt.Println(err)
	}

	// Return response
	return response
}

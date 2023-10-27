package Request

import (
	"api_tester/Class/Console"
	"fmt"
	"net/http"
)

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

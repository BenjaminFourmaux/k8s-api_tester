package Request

import (
	"api_tester/Entity"
	"fmt"
	"net/http"
	URL "net/url"
	"strings"

	Console "api_tester/Class/Console"
)

// Options Type
type Options struct {
	HeaderParam []string
	Parameter   []string
}

// Send
/*
Wrapper to send an HTTP request with the given Method
*/
func Send(url string, method string, HttpOptions Options) *http.Response {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		return GetRequest(url, HttpOptions)
	default:
		return GetRequest(url, HttpOptions)
	}
}

func GetRequest(url string, options Options) *http.Response {
	// Url Parameters
	if len(options.Parameter) > 0 {
		params := ParametersBuilder(options.Parameter)
		url = url + "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Console.Error("Error creating Request")
		panic(err)
	}

	// Add Headers
	req.Header = HeadersBuilder(options.HeaderParam)

	// Create HTTP client
	client := &http.Client{}

	// Request execution
	response, err := client.Do(req)
	if err != nil {
		Console.Error("The HTTP request failed")
		fmt.Println(err)
	}

	// Return response
	return response
}

// BuildOption
/*
Create Request Options from TestConfig
*/
func BuildOption(config Entity.Test) Options {
	options := Options{}

	if len(config.Headers) > 0 {
		options.HeaderParam = config.Headers
	}
	if len(config.Parameters) > 0 {
		options.Parameter = config.Parameters
	}

	return options
}

// HeadersBuilder
/*
Adding HTTP Headers from list of string provided by TestConfig Headers field
*/
func HeadersBuilder(headers []string) http.Header {
	httpHeaders := http.Header{}

	if headers != nil && len(headers) > 0 {
		for _, headerStr := range headers {
			keyValue := strings.Split(headerStr, "=")
			httpHeaders.Add(keyValue[0], keyValue[1])
		}
	}

	return httpHeaders
}

// ParametersBuilder
/*
Adding URL Parameters from list of string provided by TestConfig Parameters field
*/
func ParametersBuilder(params []string) URL.Values {
	urlParameters := URL.Values{}

	if params != nil && len(params) > 0 {
		for _, paramStr := range params {
			keyValue := strings.Split(paramStr, "=")
			urlParameters.Add(keyValue[0], keyValue[1])
		}
	}

	return urlParameters
}

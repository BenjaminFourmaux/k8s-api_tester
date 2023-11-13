package Pipeline

import (
	"fmt"
	"io"
	"strconv"

	Assert "api_tester/Class/Assert"
	Console "api_tester/Class/Console"
	Request "api_tester/Class/Request"
	Entity "api_tester/Entity"
)

/*
Create a test Pipeline
*/
func Create(testConfig Entity.TestConfig) int {
	numberOfSuccess := 0

	// Browse tests
	for index, test := range testConfig.Tests {
		isSuccess := false

		// Step
		PrintStep(index+1, test.Name)

		// Act
		PrintAct(test.Method, test.Url)
		Console.Write(test.Body)
		HttpOptions := Request.BuildOption(test) // Add Options like Header, Url parameter and Body to the request
		response := Request.Send(test.Url, test.Method, HttpOptions)

		b, _ := io.ReadAll(response.Body)
		fmt.Println(string(b))

		// Assert
		PrintAssert("HTTP code Equal to the expected code")
		if Assert.HTTPCodeEqual(*response, test.ExpectedHTTPCode) {
			Console.Success("HTTP code is " + strconv.Itoa(response.StatusCode))

			isSuccess = true
		} else {
			Console.Write(strconv.Itoa(response.StatusCode))
		}

		if isSuccess {
			numberOfSuccess++
		}
	}

	// Return the number of successful test
	return numberOfSuccess
}

// Resume
/*
Display a resume of tests
*/
func Resume(numberOfSuccess int, totalTest int) {
	successPercentage := (float64(numberOfSuccess) / float64(totalTest)) * 100.0

	if numberOfSuccess == totalTest { // Successful
		message := strconv.Itoa(numberOfSuccess) + "/" + strconv.Itoa(totalTest) + " Successful Tests (" + strconv.FormatFloat(successPercentage, 'f', 2, 64) + " %)"
		Console.Success(message)
	}
}

// PrintStep
/*
Info in the console for Step test
*/
func PrintStep(number int, name string) {
	message := "> Test Step  n° " + strconv.Itoa(number) + " - " + name
	Console.PrintWithColor(message, "blue")
}

// PrintAct
/*
Print in the console for Act test
*/
func PrintAct(method string, url string) {
	message := ">> Act - Send " + method + " " + url
	Console.PrintWithColor(message, "")
}

// PrintAssert
/*
Print in the console for Assertion test
*/
func PrintAssert(name string) {
	message := ">> Assert - " + name
	Console.PrintWithColor(message, "")
}

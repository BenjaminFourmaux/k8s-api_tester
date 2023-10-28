package Pipeline

import (
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
		response := Request.Send(test.Url, test.Method)

		// Assert
		PrintAssert("HTTP code Equal to the expected code")
		if Assert.HTTPCodeEqual(*response, test.ExpectedHTTPCode) {
			Console.Success(">>> HTTP code is " + strconv.Itoa(response.StatusCode))

			isSuccess = true
		} else {
			Console.Write(strconv.Itoa(response.StatusCode))
		}

		if isSuccess {
			numberOfSuccess++
		}
	}

	// Return the number of succesfull test
	return numberOfSuccess
}

func Resume(numberOfSuccess int, totalTest int) {
	successPercentage := (float64(numberOfSuccess) / float64(totalTest)) * 100.0

	if numberOfSuccess == totalTest { // Successful
		message := strconv.Itoa(numberOfSuccess) + "/" + strconv.Itoa(totalTest) + " Successful Tests (" + strconv.FormatFloat(successPercentage, 'f', 2, 64) + " %)"
		Console.Success(message)
	}
}

/*
Info in the console for Step test
*/
func PrintStep(number int, name string) {
	message := "> Test Step  n° " + strconv.Itoa(number) + " - " + name
	Console.PrintWithColor(message, "blue")
}

func PrintAct(method string, url string) {
	message := ">> Act - Send " + method + " " + url
	Console.PrintWithColor(message, "")
}

func PrintAssert(name string) {
	message := ">> Assert - " + name
	Console.PrintWithColor(message, "")
}

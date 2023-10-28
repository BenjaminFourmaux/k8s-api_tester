package main

import (
	"fmt"

	Asssert "api_tester/Class/Assert"
	Console "api_tester/Class/Console"
	Request "api_tester/Class/Request"
)

func main() {
	/* - Startup - */
	Console.Write("K8s - API Tester")
	Console.Write("version : 0.0.1")
	Console.Write("author : Benjamin Fourmaux -- Beruet  https://github.com/BenjaminFourmaux/")
	Console.Write("license : Apache 2.0")
	Console.Write("--------------------")

	/* --- Test Pipeline --- */

	// 1. Arrange - Send request
	response := Request.Get_Request("https://google.com")

	// 2. Assert
	if Asssert.HTTPCodeEqual(*response, 200) {
		Console.Success("HTTP code is " + fmt.Sprint(200))
	} else {
		Console.Error("HTTP code is " + fmt.Sprint(response.StatusCode) + " not " + fmt.Sprint(200))
	}
}

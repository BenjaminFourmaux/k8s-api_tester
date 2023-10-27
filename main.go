package main

import (
	Console "api_tester/Class/Console"
	Request "api_tester/Class/Request"
)

func main() {
	Console.Debug("Hello World")

	Request.Get_Request("https://google.com/exemple")
}

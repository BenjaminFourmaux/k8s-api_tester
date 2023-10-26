package Console

var colorBlue = "\033[34m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorRed = "\033[31m"
var colorReset = "\033[0m"

func Debug(message string) {
	println(string(colorBlue) + message + string(colorReset))
}

func Success(message string) {
	println(string(colorGreen) + "SUCCESS : " + message + string(colorReset))
}

func Warning(message string) {
	println(string(colorYellow) + "WARNING : " + message + string(colorReset))
}

func Error(message string) {
	println(string(colorRed) + "ERROR : " + message + string(colorReset))
}

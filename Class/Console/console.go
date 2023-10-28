package Console

var colorBlue = "\033[34m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorRed = "\033[31m"
var colorReset = "\033[0m"

func Write(message string) {
	println(message)
}

func PrintWithColor(message string, color string) {
	switch color {
	case "blue":
		println(string(colorBlue) + message + string(colorReset))
	case "green":
		println(string(colorGreen) + message + string(colorReset))
	case "yellow":
		println(string(colorYellow) + message + string(colorReset))
	case "red":
		println(string(colorRed) + message + string(colorReset))
	default:
		println(message)
	}
}

func Debug(message string) {
	println(string(colorReset) + message + string(colorReset))
}

func Info(message string) {
	println(string(colorBlue) + "INFO : " + message + string(colorReset))
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

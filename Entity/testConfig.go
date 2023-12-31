package Entity

type TestConfig struct {
	Tests []Test // List of test
}

// Describe a Test
type Test struct {
	Name             string // Represent the Name of the test. Ex: Internet connection
	Url              string // The HTTP url of the API
	Method           string // Wich HTTP Method do you want use send request
	ExpectedHTTPCode int    `yaml:"expectedHTTPCode"` // The expected HTTP code
}

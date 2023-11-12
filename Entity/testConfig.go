package Entity

type TestConfig struct {
	Tests []Test // List of test
}

// Test Describe a Test
type Test struct {
	Name             string   // Represent the Name of the test. Ex: Internet connection
	Url              string   // The HTTP Url of the API
	Method           string   // Which HTTP Method do you want use send request
	ExpectedHTTPCode int      `yaml:"expectedHTTPCode"` // The expected HTTP code
	Headers          []string `yaml:"headers,flow"`     // List of HTTP Headers
	Parameters       []string `yaml:"parameters,flow"`  // List of URL Parameters
}

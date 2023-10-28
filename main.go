package main

import (
	"os"

	"gopkg.in/yaml.v2"

	Console "api_tester/Class/Console"
	Pipeline "api_tester/Class/Pipeline"
	Entity "api_tester/Entity"
)

func main() {
	/* - Startup - */
	Console.Write("K8s - API Tester")
	Console.Write("version : 0.0.1")
	Console.Write("source : https://github.com/BenjaminFourmaux/k8s-api_tester")
	Console.Write("author : Benjamin Fourmaux -- Beruet")
	Console.Write("license : Apache 2.0")
	Console.Write("--------------------")

	/* -- Read Tests manifest YAML file  -- */
	//yamlFileTests, err := os.ReadFile("Resource/tests.yaml")
	yamlFileTests, err := os.ReadFile("/etc/api_tester/tests.yaml")
	if err != nil {
		Console.Error("Unable to read Test config YAML file")
		panic(err)
	}

	var testConfig Entity.TestConfig

	err = yaml.Unmarshal(yamlFileTests, &testConfig)

	if err != nil {
		Console.Error("Unable to parse in YAML. Check syntax")
		panic(err)
	}

	/* --- Test Pipeline --- */
	numberOfSuccess := Pipeline.Create(testConfig)

	Console.Write("--------------------")

	/* --- Resume --- */
	Pipeline.Resume(numberOfSuccess, len(testConfig.Tests))
}

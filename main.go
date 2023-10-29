package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"

	Console "api_tester/Class/Console"
	Pipeline "api_tester/Class/Pipeline"
	Entity "api_tester/Entity"
)

func main() {
	/* - Startup - */
	Console.Write("K8s - API Tester")
	Console.Write("version : 1.0.0")
	Console.Write("source : https://github.com/BenjaminFourmaux/k8s-api_tester")
	Console.Write("author : Benjamin Fourmaux -- Beruet")
	Console.Write("license : Apache 2.0")
	Console.Write("--------------------")

	/* -- Read ENV vars -- */
	ENV_ISCONTAINERIZED := os.Getenv("IS_CONTAINERIZED")
	ENV_DEPLOY := os.Getenv("DEPLOY")

	/* -- Read Tests manifest YAML file  -- */
	var yamlFileTests []byte
	var err error
	if ENV_ISCONTAINERIZED == "true" {
		yamlFileTests, err = os.ReadFile("/etc/api_tester/tests.yaml")
	} else {
		yamlFileTests, err = os.ReadFile("Resource/tests.yaml")
	}
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

	/* -- Infinit loop -- */
	if ENV_ISCONTAINERIZED == "true" && ENV_DEPLOY != "CronJob" {
		for {
			time.Sleep(time.Second)
		}
	}
}

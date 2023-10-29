# k8s API Tester
A small app in Go to test your APIs in Kubernetes enviromnent 

[![](https://badgen.net/badge/color/1.21.3/00aed8?label=GoLang)]()
[![](https://badgen.net/badge/icon/docker?icon=docker&label)]()
[![](https://badgen.net/github/tag/BenjaminFourmaux/k8s-api_tester?color=green&cache=600)](https://github.com/BenjaminFourmaux/k8s-api_tester/tags)

![Banner](/Resource/k8s-api_tester-banner.png)

## Features ✨
- [x] Send a HTTP request 
- [x] Declare tests using a YAML file who describe test
- [x] Dockerfile to build the Docker image 🐳
- [x] Wget and Curl package are installed !
- [ ] Kubernetes YAML manifest to deploy pod for test in your Kube env ⚓

## To start unsing API Tester 🚀
### Go
Install Go lang `1.21.3` on your system [Download it](https://go.dev/dl/).
Clone (or dezip) this project and run the following command:
```bash
go run main.go
```

### Docker
Build Docker image with  Dockerfile
```bash
docker build -t api_tester:1.0.0 -f Infrastructure/docker/dockerfile .
```

Or download the image from release files named `api_tester.tar`.
Load this image on your Docker
```bash
docker load -i api_tester.tar
```
The image is save on your local image registry.
To deploy a container, use the following command. And check env var if necessary.
```bash
docker run -e IS_CONTAINERIZED=true api_tester:1.0.0
``` 

### Kubernetes 
Chose a deployment type between `Pod`, `DaemonSet`, `Stateful` or `CronJob` depending of why you need.
To deploy a deployment, run the following command:
```bash
kubectl apply -f deploy-pod.yaml
```
You can modify files if necessary.

## Describe your Tests with YAML file
Tests are described in `tests.yaml` YAML file. It location in container is `/etc/api_tester/tests.yaml`. 
\
You can overwrite this file by binding your own file as a `volume` to the container location.

### Properties
| Key | Description | Type |
| --- | ----------- | ---- |
| `tests` | Array of test | `Array` |
| `name` | Name of the test | `String` |
| `url` | URL of the API to test | `String` |
| `method` | HTTP method to use | `String` |
| `expectedHTTPCode` | The expected HTTP code | `Int` |

### Exemple
```yaml
tests:
  - name: "Test Internet connection"
    url: "https://google.com"
    method: "GET"
    expectedHTTPCode: 200

  - name: "Back API health"
    url: "http://back-api/health"
    method: "GET"
    expectedHTTPCode: 200
```

## Contributors
[![](https://badgen.net/github/contributors/BenjaminFourmaux/k8s-api_tester)](https://github.com/BenjaminFourmaux/k8s-api_tester/graphs/contributors)
- :crown: [Benjamin Fourmaux](https://github.com/BenjaminFourmaux)

## Licence
All Service Down project's is under [**Apache License v2**](https://www.apache.org/licenses/LICENSE-2.0).
You can:
- Reuse the code 
- Modified the code
- Build the code

You must **Mention** the © Copyright if you use and modified code for your own profit. Thank you

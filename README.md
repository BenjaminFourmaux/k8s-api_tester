# k8s API Tester
A small app in Go to test your APIs in Kubernetes enviromnent 

[![](https://badgen.net/badge/color/1.21.3/00aed8?label=GoLang)]()
[![](https://badgen.net/badge/icon/docker?icon=docker&label)]()

![Banner](/Resource/k8s-api_tester-banner.png)

## Features ✨
- [x] Send a HTTP request 
- [X] Declare tests using a YAML file who describe test
- [ ] Dockerfile to build the Docker image 🐳
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
docker build -t api_tester:0.0.1 -f Infrastructure/docker/dockerfile .
```

Or download the image from release files named `api_tester.tar`.
Load this image on your Docker
```bash
docker load -i api_tester.tar
```
The image is save on your local image registry.
To deploy a container, use the following command. And check env var if necessary.
```bash
docker run -e IS_CONTAINERIZED=true api_tester:latest
``` 

### Kubernetes 
Chose a deployment type between `Pod`, `DaemonSet`, `Stateful` or `CronJob` depending of why you need.
To deploy a deployment, run the following command:
```bash
kubectl apply -f deploy-pod.yaml
```
You can modify files if necessary.

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

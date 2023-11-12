# Changelog
[![](https://badgen.net/github/tag/BenjaminFourmaux/k8s-api_tester?cache=600)]() [![](https://badgen.net/github/release/BenjaminFourmaux/k8s-api_tester?cache=600)]() [![](https://badgen.net/github/tags/BenjaminFourmaux/k8s-api_tester)]()

---
## v1.1.0 - HTTP Request
Add some stuff to use all HTTP
### Added
- Supporting HTTP Header on `Request` Class
- Some Func documentation
### Modified
- Tests YAML definition to add properties : `test.headers`
- Change some Func documentation

## v1.0.0 - First One
First version of k8s Api Tester, you can only send HTTP GET request and check response status code
### Added
- `Request` Class to send HTTP Request
- `Console` Class to have custom print with colors
- `Pipeline` Class to build tests
- `Assert` Class to have assertions like to compare response HTTP code with excepted code
- `TestConfig` Entity to unmarshall tests config YAML file
- `Dockerfile` to create Docker image of this app
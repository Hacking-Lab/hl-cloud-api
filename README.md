# HL-Cloud API

This repository contains the API for the HL-Cloud service. It is a RESTful API that can be used to mostly get information about the running service, but also contains a licence key checker.

## Installation

When starting the API it will listen on the port 1337 and can be accessible on http://localhost:1337.

### Docker (recommended)

You can launch the API by using [Golang's official docker images](https://hub.docker.com/_/golang).

```bash
$ docker compose up --build
```

### Building & Running

First make sure you've installed Go for your operating system from the [official website](https://go.dev/dl/).

Can also build the API and run it by running the following command:

```bash
$ go build -o hlcloud-api main.go
$ ./hlcloud-api
```

### Running

First make sure you've installed Go for your operating system from the [official website](https://go.dev/dl/).

Then you can directly launch the API by running the following command:

```bash
$ go run main.go
```

## Licence Key Checker

The license key checker is still in development and is not yet ready for use, it is recommended not to use it at the moment.

To check if a license you can use the following command:

```bash
$ curl -X GET http://localhost:1337/license/check?key=XXXX-XXXX-XXXX-XXXX-XXXX
```

#### Valid Key Response

```json
{
  "valid": true,
}
```

#### Invalid Key Response

```json
{
  "valid": false,
}
```

## Endpoints

The following endpoints are currently available:

| Endpoint         | Description                                                                        |
|------------------|------------------------------------------------------------------------------------|
| `/`              | Get the API version that is currently running along with the available endpoints   |
| `/info`          | Get information about the API with some runtime information                        |
| `/license/check` | Check if a license key is valid, the key must be provided as a `key` GET parameter |
| `/pricing`       | Get the pricing of the cloud service                                               |
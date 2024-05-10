🌍 [English](README.md)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

- [Currency Exchange Rate Application](#currency-exchange-rate-application)
    - [Direct Usage](#direct-usage)
    - [Docker](#docker)
  - [Endpoints](#endpoints)
  - [Environment Variables](#environment-variables)
  - [Getting Started](#getting-started)
  - [Acknowledgments](#acknowledgments)
  - [Swagger](#swagger)
  - [Testing](#testing)
  - [Issues](#issues)


# Currency Exchange Rate Application

This application provides access to currency exchange rates using data from the Central Bank of the Russian Federation (CBR). It exposes several endpoints for GET requests, allowing users to retrieve information about the application, get currency exchange rates, and access Swagger documentation.

### Direct Usage
1. Install go version 1.22 or higher. See [official docs](https://go.dev/doc/install).
2. Install dependencies. Run `go mod tidy`.
3. Run app. See [allowed usage](#getting-started).
4. Use [endpoints](#endpoints).

### Docker
1. Install docker engine. See [official docs](https://docs.docker.com/engine/install/).
2. Build image. Run `docker build -t exchange_rate_app:tag_name .` or use an existing image at [Dockerhub](https://hub.docker.com/r/sveboo/exchange-rate-app).
3. Run app container.
4. Use [endpoints](#endpoints).

## Endpoints
Sure, here's how you can represent the endpoints in a table format:

| Endpoint              | Method | Parameters               | Description                                                                                                            |
|-----------------------|--------|--------------------------|------------------------------------------------------------------------------------------------------------------------|
| /info/                |<span style="color:green">GET</span>| None                     | Returns information about the application.                                                                             |
| /info/currency/       |<span style="color:green">GET</span>| - date (optional): Date in the format YYYY-MM-DD<br>- currency (optional): Currency in ISO 4217 standard      | Returns currency exchange rates based on the provided date and currency.                                                |
| /info/docs            |<span style="color:green">GET</span>| None                     | Returns Swagger documentation for the API endpoints.                                                                   |

## Environment Variables
You can customize the behavior of the application by setting the following environment variables:

**VERSION**: Specifies the version of the application.

**PORT**: Specifies the port on which the application listens for incoming connections.
If these environment variables are not set, default values from config.json will be used.

## Getting Started

To run the Currency Exchange Rate Application locally, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/Sveboo/exchange-rate-app.git
cd repository
```

2. Build:

```bash
make
```

3. Set the required environment variables (optional):

```bash
export VERSION=1.0.0
export PORT=8080
```

4. Run the application:

```bash
./currency_service
```

5. Access the API endpoints using a web browser or tools like cURL or Postman:

- `http://localhost:8080/info/`
- `http://localhost:8080/info/currency/?date=2022-01-01&currency=USD`
- `http://localhost:8080/info/docs`


## Acknowledgments

- Central Bank of the Russian Federation (CBR) for providing currency exchange rate data.
- Gin framework for web server implementation.
- Swagger for API documentation.

## Swagger

Application also provides swagger documentation accessible with endpoint `/docs`. See [endpoints](#endpoints).

## Testing

Code is covered with unit tests. Current test coverage is just under **84.3%**. You can see detailed information in [cov.html](docs/cov.html)

## Issues

If you notice any bug or vulnerability, feel free to create an Issue with a full description of the problem found.
In addition, it will be great if you offer your solution to the problem through a Pull Request.
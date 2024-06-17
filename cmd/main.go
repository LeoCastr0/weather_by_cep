package main

import (
	"github.com/leocastr0/weather_by_cep/internal/infra/repository"

	"github.com/leocastr0/weather_by_cep/internal/infra/entrypoint/controller"

	"github.com/leocastr0/weather_by_cep/internal/infra/entrypoint"

	"github.com/leocastr0/weather_by_cep/internal/infra/client"

	"github.com/leocastr0/weather_by_cep/internal/application/service"

	"github.com/leocastr0/weather_by_cep/internal/application/config"
)

func main() {
	conf := config.NewConfig()
	httpClient := client.NewHTTPClient()

	zipCodeRepo := repository.NewZipCodeRepository(httpClient, conf.ViaCepURL)
	weatherRepo := repository.NewWeatherRepository(httpClient, conf.WeatherAPIURL, conf.WeatherAPIKey)

	zipCodeService := service.NewZipCodeService(zipCodeRepo)
	weatherService := service.NewWeatherService(weatherRepo)

	weatherController := controller.NewWeatherController(weatherService, zipCodeService)

	r := entrypoint.SetupRouter(weatherController)
	r.Run(":8080")
}

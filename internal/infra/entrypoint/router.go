package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/leocastr0/weather_by_cep/internal/infra/entrypoint/controller"
)

func SetupRouter(weatherController *controller.WeatherController) *gin.Engine {
	r := gin.Default()
	r.GET("/weather/:zipcode", weatherController.GetWeather)
	return r
}

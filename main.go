package main

import (
	"lmwn-go-test/covid-19-summary/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/covid/summary", controllers.GetCovidSummary)
	router.Run()
}

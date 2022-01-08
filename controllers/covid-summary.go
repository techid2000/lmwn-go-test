package controllers

import (
	"lmwn-go-test/covid-19-summary/models"
	"lmwn-go-test/covid-19-summary/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCovidSummary(ctx *gin.Context) {
	cases, err := models.GetCovidCases()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	summary := services.GenerateCovidSummary(cases)

	ctx.JSON(http.StatusOK, summary)
}

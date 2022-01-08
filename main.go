package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CovidCase struct {
	ConfirmDate    interface{} `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            interface{} `json:"Age"`
	Gender         interface{} `json:"Gender"`
	GenderEn       interface{} `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	NationEn       interface{} `json:"NationEn"`
	Province       interface{} `json:"Province"`
	ProvinceId     interface{} `json:"ProvinceId"`
	District       interface{} `json:"District"`
	ProvinceEn     interface{} `json:"ProvinceEn"`
	StatQuarantine interface{} `json:"StatQuarantine"`
}

type CovidCases struct {
	Data []CovidCase `json:"Data"`
}

type Summary struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

func main() {
	router := gin.Default()
	router.GET("/covid/summary", func(c *gin.Context) {
		response, error := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")

		if error != nil {
			panic(error.Error())
		}

		defer response.Body.Close()

		covidCases := new(CovidCases)
		json.NewDecoder(response.Body).Decode(covidCases)

		summary := new(Summary)
		summary.Province = make(map[string]int)
		summary.AgeGroup = make(map[string]int)

		summary.AgeGroup["0-30"] = 0
		summary.AgeGroup["31-60"] = 0
		summary.AgeGroup["60+"] = 0
		summary.AgeGroup["N/A"] = 0

		for _, _case := range covidCases.Data {
			provinceEn, assertionOk := _case.ProvinceEn.(string)
			if assertionOk {
				summary.Province[provinceEn]++
			} else {
				summary.Province["N/A"]++
			}

			age, assertionOk := _case.Age.(float64)
			if assertionOk {
				if age >= 0 && age <= 30 {
					summary.AgeGroup["0-30"]++
				} else if age >= 31 && age <= 60 {
					summary.AgeGroup["31-60"]++
				} else if age > 60 {
					summary.AgeGroup["60+"]++
				} else {
					summary.AgeGroup["N/A"]++
				}
			} else {
				summary.AgeGroup["N/A"]++
			}
		}
		c.JSON(http.StatusOK, summary)
	})
	router.Run()
}

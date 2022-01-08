package models

import (
	"encoding/json"
	"lmwn-go-test/covid-19-summary/types"
	"net/http"
)

func GetCovidCases() (*types.CovidCases, error) {
	// TODO: Use Redis caching
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	cases := new(types.CovidCases)
	json.NewDecoder(resp.Body).Decode(cases)

	return cases, err
}

package services

import "lmwn-go-test/covid-19-summary/types"

func GenerateCovidSummary(cases *types.CovidCases) *types.CovidSummary {
	summary := &types.CovidSummary{
		Province: make(map[string]int),
		AgeGroup: make(map[string]int),
	}

	// This is crucial since there are only 4 groups.
	summary.AgeGroup["0-30"] = 0
	summary.AgeGroup["31-60"] = 0
	summary.AgeGroup["60+"] = 0
	summary.AgeGroup["N/A"] = 0

	for _, _case := range cases.Data {
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

	return summary
}

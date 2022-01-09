package services

import (
	"fmt"
	"lmwn-go-test/covid-19-summary/types"
	"testing"
)

type ProvinceFact struct {
	name  string
	count int
}

type AgeGroupFact struct {
	name  string
	count int
}

type CovidSummaryTestCase struct {
	mockCases     types.CovidCases
	provincesFact []ProvinceFact
	ageGroupsFact []AgeGroupFact
}

func TestGenerateCovidSummary(t *testing.T) {
	testCases := []CovidSummaryTestCase{
		// Test case #1
		{
			mockCases: types.CovidCases{
				Data: []types.CovidCase{},
			},
			provincesFact: []ProvinceFact{},
			ageGroupsFact: []AgeGroupFact{
				{
					name:  "0-30",
					count: 0,
				},
				{
					name:  "31-60",
					count: 0,
				},
				{
					name:  "60+",
					count: 0,
				},
				{
					name:  "N/A",
					count: 0,
				},
			},
		},
		// Test case #2
		// NOTE: Go's JSON decoder will treat any number value to be a type float64.
		{
			mockCases: types.CovidCases{
				Data: []types.CovidCase{
					{
						ConfirmDate:    "2022-01-01",
						No:             1,
						Age:            float64(15),
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "กุรงเทพมหานคร",
						ProvinceId:     1,
						District:       "ปุทมวัน",
						ProvinceEn:     "Bangkok",
						StatQuarantine: 1,
					},
					{
						ConfirmDate:    "2022-01-01",
						No:             2,
						Age:            float64(45),
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "กรุงเทพมหานคร",
						ProvinceId:     1,
						District:       "ปทุมวัน",
						ProvinceEn:     "Bangkok",
						StatQuarantine: 1,
					},
					{
						ConfirmDate:    "2022-01-01",
						No:             3,
						Age:            float64(65),
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "ระยอง",
						ProvinceId:     2,
						District:       "เมืองระยอง",
						ProvinceEn:     "Rayong",
						StatQuarantine: 1,
					},
				},
			},
			provincesFact: []ProvinceFact{
				{
					name:  "Bangkok",
					count: 2,
				},
				{
					name:  "Rayong",
					count: 1,
				},
			},
			ageGroupsFact: []AgeGroupFact{
				{
					name:  "0-30",
					count: 1,
				},
				{
					name:  "31-60",
					count: 1,
				},
				{
					name:  "60+",
					count: 1,
				},
				{
					name:  "N/A",
					count: 0,
				},
			},
		},
		// Test case #3
		{
			mockCases: types.CovidCases{
				Data: []types.CovidCase{
					{
						ConfirmDate:    "2022-01-01",
						No:             1,
						Age:            float64(-1),
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "กุรงเทพมหานคร",
						ProvinceId:     1,
						District:       "ปุทมวัน",
						ProvinceEn:     1234,
						StatQuarantine: 1,
					},
					{
						ConfirmDate:    "2022-01-01",
						No:             2,
						Age:            false,
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "กรุงเทพมหานคร",
						ProvinceId:     1,
						District:       "ปทุมวัน",
						ProvinceEn:     false,
						StatQuarantine: 1,
					},
					{
						ConfirmDate:    "2022-01-01",
						No:             3,
						Age:            nil,
						Gender:         "ชาย",
						GenderEn:       "Male",
						Nation:         "ไทย",
						NationEn:       "Thai",
						Province:       "ระยอง",
						ProvinceId:     2,
						District:       "เมืองระยอง",
						ProvinceEn:     nil,
						StatQuarantine: 1,
					},
				},
			},
			provincesFact: []ProvinceFact{
				{
					name:  "N/A",
					count: 3,
				},
			},
			ageGroupsFact: []AgeGroupFact{
				{
					name:  "0-30",
					count: 0,
				},
				{
					name:  "31-60",
					count: 0,
				},
				{
					name:  "60+",
					count: 0,
				},
				{
					name:  "N/A",
					count: 3,
				},
			},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Testing case #%d\n", i+1)
		summary := GenerateCovidSummary(&testCase.mockCases)
		if l := len(summary.Province); l != len(testCase.provincesFact) {
			t.Errorf("Number of Provinces got %d; want %d", l, len(testCase.provincesFact))
		}
		for _, province := range testCase.provincesFact {
			if _, ok := summary.Province[province.name]; !ok {
				t.Errorf("Province should contains key %s", province.name)
			}
			if c := summary.Province[province.name]; c != province.count {
				t.Errorf("Number of Province %s got %d; want %d", province.name, c, province.count)
			}
		}
		for _, ageGroup := range testCase.ageGroupsFact {
			if _, ok := summary.AgeGroup[ageGroup.name]; !ok {
				t.Errorf("AgeGroup should contains key %s", ageGroup.name)
			}
			if c := summary.AgeGroup[ageGroup.name]; c != ageGroup.count {
				t.Errorf("Number of AgeGroup %s got %d; want %d", ageGroup.name, c, ageGroup.count)
			}
		}
	}
}

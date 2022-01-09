package types

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

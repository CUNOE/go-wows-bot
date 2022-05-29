package model

type ChangeBind struct {
	Code      int    `json:"code"`
	Data      string `json:"data"`
	Message   string `json:"message"`
	QueryTime int    `json:"queryTime"`
}

package models

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

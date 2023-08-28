package models

type Response struct {
	StatusCode    int32  `json:"status_code"`
	StatusMassage string `json:"status_massage,omitempty"`
}

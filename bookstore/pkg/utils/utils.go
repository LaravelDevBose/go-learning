package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func (res *Response) PrepareResponse(data any, status int, message string) ([]byte, error) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return json.Marshal(response)
}

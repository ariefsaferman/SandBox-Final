package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func ResponJson(w http.ResponseWriter, status int, err error) {
	msg := "success"
	if err != nil {
		msg = err.Error()
	}

	resp := Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Message:    msg,
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonResp)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func resWithErr(
	w http.ResponseWriter,
	code int,
	msg string,
) {
	if code > 499 {
		log.Println("Responsing with 5xx error: ", msg)
	}

	resWithJson(w, code, errResponse{
		Error: msg,
	})
}

func resWithJson(
	w http.ResponseWriter,
	code int,
	payload interface{},
) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

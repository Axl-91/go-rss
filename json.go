package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(wr http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		wr.WriteHeader(500)
		log.Fatal(err)
	}

	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(code)

	_, err = wr.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

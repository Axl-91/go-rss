package main

import (
	"net/http"
)

func handlerReadiness(wr http.ResponseWriter, req *http.Request) {
	respondWithJSON(wr, 200, struct{}{})
}

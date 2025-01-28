package main

import (
	"net/http"
)

func handlerError(wr http.ResponseWriter, req *http.Request) {
	respondWithError(wr, 400, "Something went wrong")
}

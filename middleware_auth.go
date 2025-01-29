package main

import (
	"fmt"
	"net/http"

	"github.com/Axl-91/go-rss/internal/auth"
	"github.com/Axl-91/go-rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(wr, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(wr, 403, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(wr, r, user)
	}
}

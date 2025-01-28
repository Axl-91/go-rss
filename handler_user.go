package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axl-91/go-rss/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(wr http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err :=
		apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			Name:      params.Name,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't create user: %v", err))
	}

	respondWithJSON(wr, 200, user)
}

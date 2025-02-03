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
		return
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
		return
	}

	respondWithJSON(wr, 201, user)
}

func (apiCfg *apiConfig) handlerGetUser(wr http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(wr, 200, user)
}

func (apiCfg *apiConfig) handlerGetPostsForUser(wr http.ResponseWriter, r *http.Request, user database.User) {
	posts, err :=
		apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  10,
		})
	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't fetch posts: %v", err))
		return
	}

	respondWithJSON(wr, 200, posts)
}

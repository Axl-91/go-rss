package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axl-91/go-rss/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(wr http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err :=
		apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			Name:      params.Name,
			Url:       params.Url,
			UserID:    user.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(wr, 201, feed)
}

func (apiCfg *apiConfig) handlerGetFeeds(wr http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(wr, 200, feeds)
}

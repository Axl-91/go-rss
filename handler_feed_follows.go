package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axl-91/go-rss/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(wr http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err :=
		apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
			ID:        uuid.New(),
			UserID:    user.ID,
			FeedID:    params.FeedID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(wr, 201, feedFollow)
}

// func (apiCfg *apiConfig) handlerGetFeedsFollow(wr http.ResponseWriter, r *http.Request) {
// 	feeds, err := apiCfg.DB.GetFeeds(r.Context())

// 	if err != nil {
// 		respondWithError(wr, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
// 		return
// 	}

// 	respondWithJSON(wr, 200, feeds)
// }

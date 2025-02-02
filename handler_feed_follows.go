package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axl-91/go-rss/internal/database"
	"github.com/go-chi/chi"
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
		respondWithError(wr, 400, fmt.Sprintf("Couldn't create feed follows: %v", err))
		return
	}

	respondWithJSON(wr, 201, feedFollow)
}

func (apiCfg *apiConfig) handlerGetFeedFollows(wr http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't get feed follows from user: %v", err))
		return
	}

	respondWithJSON(wr, 200, feedFollows)
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(wr http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "feedFollowsID")

	feedFollowID, err := uuid.Parse(feedFollowIDstr)

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't parse ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(wr, 400, fmt.Sprintf("Couldn't delete Feed Follows: %v", err))
		return
	}

	respondWithJSON(wr, 200, struct{}{})
}

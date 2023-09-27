package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/smrnjeet222/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	type Tparams struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Tparams{}
	err := decoder.Decode(&params)
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Err parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't create Feed follow: %v", err))
		return
	}

	resWithJson(w, 201, feed)
}

func (apiCfg *apiConfig) handlerGetFeedFollows(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't Get Feed follow: %v", err))
		return
	}
	resWithJson(w, 200, feeds)
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Feed Id not found: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't Get Feed follow: %v", err))
		return
	}
	resWithJson(w, 200, struct{}{})
}

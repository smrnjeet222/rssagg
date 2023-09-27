package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/smrnjeet222/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	type Tparams struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Tparams{}
	err := decoder.Decode(&params)
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Err parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't create User: %v", err))
		return
	}

	resWithJson(w, 201, feed)
}

func (apiCfg *apiConfig) handlerGetFeed(
	w http.ResponseWriter,
	r *http.Request,
) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}
	resWithJson(w, 200, feeds)
}

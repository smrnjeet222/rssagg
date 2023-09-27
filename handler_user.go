package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/smrnjeet222/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	type Tparams struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Tparams{}
	err := decoder.Decode(&params)
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Err parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't create User: %v", err))
		return
	}

	resWithJson(w, 201, user)
}

func (apiCfg *apiConfig) handlerGetUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	resWithJson(w, 200, user)
}

func (apiCfg *apiConfig) handlerGetPostsForUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.Gouser,
) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		resWithErr(w, 400, fmt.Sprintf("Couldn't get Users posts: %v", err))
		return
	}

	resWithJson(w, 200, posts)
}

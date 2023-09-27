package main

import (
	"fmt"
	"net/http"

	"github.com/smrnjeet222/rssagg/internal/auth"
	"github.com/smrnjeet222/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.Gouser)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			resWithErr(w, 403, fmt.Sprintf("Auth Err: %v", err))
			return
		}

		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			resWithErr(w, 404, fmt.Sprintf("Couldn't Find User: %v", err))
			return
		}

		handler(w, r, user)
	}
}

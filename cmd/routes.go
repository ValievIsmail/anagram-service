package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	log "github.com/sirupsen/logrus"
)

func createHTTPHandler(dict []string) (http.Handler, error) {
	mux := chi.NewMux()

	mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}).Handler)

	mux.Route("/", func(r chi.Router) {
		r.Post("/load", func(w http.ResponseWriter, r *http.Request) {
			if err := json.NewDecoder(r.Body).Decode(&dict); err != nil {
				http.Error(w, "invalid data, try again", http.StatusBadRequest)
				return
			}

			DictToLowerCase(dict)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("SUCCESS LOAD DICTIONARY, TRY GET ANAGRAM =)"))
		})

		r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
			if len(dict) == 0 {
				http.Error(w, "dictionary is empty, download the dictionary first", http.StatusBadRequest)
				return
			}

			word := strings.ToLower(r.URL.Query().Get("word"))

			if word == "" {
				http.Error(w, "Enter the word", http.StatusBadRequest)
				return
			}

			anagrams := SearchAnagrams(dict, word)

			if len(anagrams) == 0 {
				http.Error(w, "anagram not found :(\ntry another word", http.StatusNotFound)
				return
			}

			data, err := json.Marshal(&anagrams)
			if err != nil {
				log.Errorf("get handler write data: %v", err)
				http.Error(w, "failed write anagrams", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if _, err := w.Write(data); err != nil {
				log.Errorf("get handler write: %v", err)
				http.Error(w, "failed write anagrams", http.StatusInternalServerError)
				return
			}
		})
	})

	return mux, nil
}

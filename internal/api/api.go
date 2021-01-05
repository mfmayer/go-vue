package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mfmayer/go-vue/internal/log"
)

func titlePrefix(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"title": "let's",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func version(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"version": "0.0.1",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type user struct {
	Name string `json:"name"`
}

func setName(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	var u user
	if err := json.Unmarshal(body, &u); err != nil {
		log.Error.Printf("Error unmarshalling body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"message": "hello " + u.Name + ", let's go with vue!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// InstallAPI installs the api handler functions
func InstallAPI(r chi.Router) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // consider to allow specific origin hosts only
	}))
	r.Get("/titlePrefix", titlePrefix)
	r.Get("/version", version)
	r.Post("/setName", setName)
}

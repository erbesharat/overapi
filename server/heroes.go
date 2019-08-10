package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, err := h.db.GetHeroes()
	if err != nil {
		log.Printf("Coudln't fetch the heroes from database: %s", err.Error())
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(heroes)
}

func (h *Handler) GetHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hero, err := h.db.GetHero(params["id"])
	if err != nil {
		log.Printf("Coudln't fetch the heroes from database: %s", err.Error())
		w.WriteHeader(500)
	}
	if hero == nil {
		w.WriteHeader(404)
	}
	json.NewEncoder(w).Encode(hero)
}

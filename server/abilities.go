package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetHeroAbilities(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	abilities, err := h.db.GetHeroAbilities(params["id"])
	if err != nil {
		log.Printf("Coudln't fetch the heroes from database: %s", err.Error())
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(abilities)
}

func (h *Handler) GetAbilities(w http.ResponseWriter, r *http.Request) {
	abilities, err := h.db.GetAbilities()
	if err != nil {
		log.Printf("Coudln't fetch the abilities from database: %s", err.Error())
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(abilities)
}

func (h *Handler) GetAbility(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ability, err := h.db.GetAbility(params["id"])
	if err != nil {
		log.Printf("Coudln't fetch the abilities from database: %s", err.Error())
		w.WriteHeader(500)
	}
	if ability == nil {
		w.WriteHeader(404)
	}
	json.NewEncoder(w).Encode(ability)
}

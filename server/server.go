package server

import (
	"net/http"

	"github.com/erbesharat/goverapi/db"
	"github.com/gorilla/mux"
)

type Handler struct {
	db  *db.DB
	mux *mux.Router
}

func New(db *db.DB) *Handler {
	mux := mux.NewRouter()
	handler := Handler{db, mux}
	mux.HandleFunc("/api/heroes", handler.GetHeroes)
	mux.HandleFunc("/api/heroes/{id}/abilities", handler.GetHeroAbilities)
	mux.HandleFunc("/api/heroes/{id}", handler.GetHero)
	mux.HandleFunc("/api/abilities", handler.GetAbilities)
	mux.HandleFunc("/api/abilities/{id}", handler.GetAbility)
	return &handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

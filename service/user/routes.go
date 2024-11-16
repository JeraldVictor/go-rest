package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodGet)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodGet)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

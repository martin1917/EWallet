package handler

import (
	"net/http"

	"github.com/martin1917/EWallet/internal/app"
)

type Handler struct {
	Services *app.Services
}

func NewHandler(services *app.Services) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetLast(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {

}

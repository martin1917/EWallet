package handler

import "github.com/gorilla/mux"

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/send", h.Send)
	router.HandleFunc("/api/transactions", h.GetLast)
	router.HandleFunc("/api/wallet/{address}/balance", h.GetBalance)
	return router
}

package handler

import "github.com/gorilla/mux"

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/send", h.Send).Methods("POST")
	router.HandleFunc("/api/transactions", h.GetLast).Methods("GET")
	router.HandleFunc("/api/wallet/{address}/balance", h.GetBalance).Methods("GET")
	return router
}

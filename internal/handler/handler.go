package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/martin1917/EWallet/internal/app"
	"github.com/martin1917/EWallet/internal/app/commands"
	"github.com/martin1917/EWallet/internal/app/errors"
	"github.com/martin1917/EWallet/internal/app/queries"
)

type Handler struct {
	Services *app.Services
}

func NewHandler(services *app.Services) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	var sendMoneyRequest commands.SendModeyRequest
	err := json.NewDecoder(r.Body).Decode(&sendMoneyRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Services.Commands.SendModeyHandler.Handle(sendMoneyRequest)
	if err != nil {
		if serr, ok := err.(*errors.AppError); ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(serr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "transaction are failed",
			})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetLast(w http.ResponseWriter, r *http.Request) {
	countParam := r.URL.Query().Get("count")
	if countParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "'count' parameter is missing",
		})
		return
	}

	count, err := strconv.Atoi(countParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "'count' parameter must be integer",
		})
		return
	}

	request := queries.GetLastRequest{Count: count}
	response, _ := h.Services.Queries.GetLastHandler.Handle(request)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address, err := uuid.Parse(vars["address"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "'address' must be uuid",
		})
		return
	}

	request := queries.GetBalanceRequest{WalletAddress: address}
	response, err := h.Services.Queries.GetBalanceHandler.Handle(request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("wallet with'address' = %s does not exist", address),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

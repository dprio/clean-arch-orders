package orderhandler

import (
	"encoding/json"
	"net/http"

	"github.com/dprio/clean-arch-orders/internal/usecase/createorder"
)

type OrderHandler struct {
	createOrderUseCase createorder.UseCase
}

func New(useCase createorder.UseCase) *OrderHandler {
	return &OrderHandler{createOrderUseCase: useCase}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request createOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createOrderUseCase.Execute(ctx, request.ToCreateOrderInput())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewCreateOrderResponse(output)
	w.Header().Add("content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

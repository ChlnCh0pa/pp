package handler

import (
	"encoding/json"
	"net/http"
	"warehouse/internal/dto"
	"warehouse/internal/service"
)

type Handler struct {
	svc service.WarehouseService
}

func NewHandleFunc(svc service.WarehouseService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ReceiveHandler(w http.ResponseWriter, r *http.Request) {

	var req dto.ReceiveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return

	}
	if req.Quantity <= 0 {
		http.Error(w, "must be positive", http.StatusBadRequest)
		return
	}

	recive := req.ToRecRequest()
	err = h.svc.Receive(r.Context(), recive)
	if err != nil {
		http.Error(w, " bebebe", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

}

func (h *Handler) ShipHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.ShipRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if req.ItemID <= 0 || req.ShelfID <= 0 {
		http.Error(w, "item_id and shelf_id must be positive", http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		http.Error(w, "quantity must be positive", http.StatusBadRequest)
		return
	}

	shipI := req.ToShipRequest()

	err = h.svc.Ship(r.Context(), shipI)
	if err != nil {
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

}

func (h *Handler) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.GetAllItems(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var response []dto.ItemResponse
	for _, item := range items {
		response = append(response, dto.ItemResponse{

			ID:   item.ID,
			SKU:  item.SKU,
			Name: item.Name,
			Unit: item.Uint,
		})
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

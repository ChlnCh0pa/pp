package dto

import "warehouse/internal/model"

type ReceiveRequest struct {
	ItemID   int64 `json:"item_id"`
	ShelfID  int64 `json:"shelf_id"`
	Quantity int   `json:"quantity"`
}

type ShipRequest struct {
	ItemID   int64 `json:"item_id"`
	ShelfID  int64 `json:"shelf_id"`
	Quantity int   `json:"quantity"`
}

type ItemResponse struct {
	ID   int64  `json:"id"`
	SKU  string `json:"sku"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (r *ReceiveRequest) ToRecRequest() model.Receive {
	return model.Receive{
		ItemID:   r.ItemID,
		ShelfID:  r.ShelfID,
		Quantity: r.Quantity,
	}
}

func (r *ShipRequest) ToShipRequest() model.Ship {
	return model.Ship{
		ItemID:   r.ItemID,
		ShelfID:  r.ShelfID,
		Quantity: r.Quantity,
	}
}

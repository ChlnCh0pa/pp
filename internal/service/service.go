package service

import (
	"context"
	"errors"
	"warehouse/internal/model"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNotEnoughStock = errors.New("not enough stock")
)

type WarehouseService interface {
	Receive(ctx context.Context, shipment model.Receive) error
	Ship(ctx context.Context, shipment model.Ship) error
}

func (s *service) GetAllItems() ([]model.Shelf, error) {

}

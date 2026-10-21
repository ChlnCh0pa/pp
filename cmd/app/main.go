package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"warehouse/internal/handler"
)

type DummyService struct{}

func (d *DummyService) Receive(ctx context.Context, itemID, shelfID int64, qty int) error {
	fmt.Printf("[LOG] Receive: item=%d, shelf=%d, qty=%d\n", itemID, shelfID, qty)
	return nil
}

func (d *DummyService) Ship(ctx context.Context, itemID, shelfID int64, qty int) error {
	fmt.Printf("[LOG] Ship: item=%d, shelf=%d, qty=%d\n", itemID, shelfID, qty)
	return nil
}

func main() {

	svc := &DummyService{}

	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /receive", h.ReceiveHandler)
	mux.HandleFunc("POST /ship", h.ShipHandler)

	fmt.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

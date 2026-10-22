package model

import "time"

type Item struct {
	ID      int64
	SKU     string
	Name    string
	Unit    string
	ShelfID int64
}

type Shelf struct {
	ID       int64
	Code     string
	Capacity int
}

type Receive struct {
	ItemID   int64
	ShelfID  int64
	Quantity int
}

type Ship struct {
	ItemID   int64
	ShelfID  int64
	Quantity int
}

type EntryBook struct {
	ID        int64
	ItemID    int64
	ShelfID   int64
	Type      string
	Quantity  int
	CreatedAt time.Time
}

type Stock struct {
	ItemID   int64
	ShelfID  int64
	Quantity int
}

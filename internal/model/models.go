package model

type Item struct {
	ID   int64
	SKU  string
	Name string
	Unit string
}

type Shelf struct {
	ID        int64
	Code      string
	Capacity  int
	ItemCount int
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

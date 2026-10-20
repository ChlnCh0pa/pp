package model

import "time"

type Item struct {
	Name string
	ID   int
}

type Shelf struct {
	Capacity int
	ID       int
	Code     string
}

type Book struct {
	ID        int
	ItemID    int
	ShelfId   int
	Type      EntryType
	CreatedAt time.Time
	Quantity  int
}

type Stock struct {
	ItemID   int
	ShelfID  int
	Quantity int
}

type EntryType string

const (
	EntryIn       EntryType = "in"
	EntryOut      EntryType = "out"
	EntryMove     EntryType = "move"
	EntryWriteOff EntryType = "writeoff"
)

type stockKey struct {
	ItemID  int64
	ShelfID int64
}

package models

import "time"

type ShoppingListItem struct {
	ID      string
	Product Product
	Notes   string
	AddedBy User
	AddedAt time.Time
}

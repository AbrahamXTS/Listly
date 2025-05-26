package models

import "time"

type Purchase struct {
	ID               string
	Product          Product
	ProductUnitPrice *float64
	PurchasedBy      User
	PurchasedAt      time.Time
}

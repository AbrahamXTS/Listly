package models

type ShoppingList struct {
	ID              string
	Title           string
	Items           []ShoppingListItem
	CreatedBy       User
	Collaborators   []User
	PurchaseHistory []Purchase
}

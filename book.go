package main

type Books struct {
	Id             int    `db:"id" json:"id"`
	Title          string `db:"title" json:"title"`
	Description    string `db:"description" json:"description"`
	Price          int    `db:"price" json:"price"`
	Stock          int    `db:"stock" json:"stock"`
	PurchaseAmount int    `db:"purchase_amount" json:"purchase_amount"`
}

// type Books struct {
// 	Id             int    `db:"id" json:"id"`
// 	Title          string `db:"title" json:"title" binding:"required"`
// 	Description    string `db:"description" json:"description" binding:"required"`
// 	Price          int    `db:"price" json:"price" binding:"required"`
// 	Stock          int    `db:"stock" json:"stock" binding:"required"`
// 	PurchaseAmount int    `db:"purchase_amount" json:"purchase_amount" binding:"required"`
// }

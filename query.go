package main

const (
	INSERT_BOOK = `INSERT INTO mst_books
	(title, description, price, stock, purchase_amount)
	VALUES
	(:title, :description, :price, :stock, :purchase_amount)`

	GET_BOOKS = `SELECT * FROM mst_books`

	GET_BOOK = `SELECT * FROM mst_books WHERE id = :id`

	DELETE_BOOK = `DELETE FROM mst_books WHERE id = :id`

	UPDATE_BOOK = `UPDATE mst_books SET title = :title, description = :description, price = :price, stock = :stock, purchase_amount = :purchase_amount WHERE id = :id`
)

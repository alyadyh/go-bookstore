package bookmodel

import (
	"go-bookstore/config"
	"go-bookstore/entities"
)

func GetAll() []entities.Books {
	rows, err := config.DB.Query(`
		SELECT 
			books.id, 
			books.title, 
			authors.name as author_name,
			publishers.name as publisher_name,
			genres.name as genre_name,
			books.publication_date, 
			books.ISBN,
			books.price,
			books.stock_qty,
			books.created_at,
			books.updated_at FROM books
		JOIN authors ON books.author_id = authors.id
		JOIN publishers ON books.publisher_id = publishers.id
		JOIN genres ON books.genre_id = genres.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var books []entities.Books

	for rows.Next() {
		var book entities.Books
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author.Name,
			&book.Publisher.Name,
			&book.Genre.Name,
			&book.PublicationDate,
			&book.ISBN,
			&book.Price,
			&book.StockQty,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			panic(err)
		}

		books = append(books, book)
	}

	return books
}

func Create(book entities.Books) bool {

	result, err := config.DB.Exec(`
		INSERT INTO books(
			title, author_id, publisher_id, genre_id, publication_date, ISBN, price, stock_qty, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		book.Title,
		book.Author.Id,
		book.Publisher.Id,
		book.Genre.Id,
		book.PublicationDate,
		book.ISBN,
		book.Price,
		book.StockQty,
		book.CreatedAt,
		book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Books {
	row := config.DB.QueryRow(`
		SELECT 
			.id, 
			books.id, 
			books.title, 
			authors.name as author_name,
			publishers.name as publisher_name,
			genres.name as genre_name,
			books.publication_date, 
			books.ISBN,
			books.price,
			books.stock_qty,
			books.created_at, 
			books.updated_at FROM books
		JOIN authors ON books.author_id = authors.id,
		JOIN publishers ON books.publisher_id = publishers.id
		JOIN genres ON books.genre_id = genres.id
		WHERE books.id = ?
	`, id)

	var book entities.Books

	err := row.Scan(
		&book.Id,
		&book.Title,
		&book.Author.Name,
		&book.Publisher.Name,
		&book.Genre.Name,
		&book.PublicationDate,
		&book.ISBN,
		&book.Price,
		&book.StockQty,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return book
}

func Update(id int, book entities.Books) bool {
	query, err := config.DB.Exec(`
		UPDATE books SET 
			title = ?,
			genre_id = ?,
			price = ?,
			stock_qty = ?,
			updated_at = ?
		WHERE id = ?`,
		book.Title,
		book.Genre.Id,
		book.Price,
		book.StockQty,
		book.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

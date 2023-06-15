package authormodel

import (
	"go-bookstore/config"
	"go-bookstore/entities"
)

func GetAll() []entities.Authors {
	rows, err := config.DB.Query(`SELECT * FROM authors`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var authors []entities.Authors

	for rows.Next() {
		var author entities.Authors
		if err := rows.Scan(&author.Id, &author.Name); err != nil {
			panic(err)
		}

		authors = append(authors, author)
	}

	return authors
}

func Create(author entities.Authors) bool {
	result, err := config.DB.Exec(`
		INSERT INTO authors (name) VALUES (?)`,
		author.Name,
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

func Detail(id int) entities.Authors {
	row := config.DB.QueryRow(`SELECT id, name FROM authors WHERE id = ? `, id)

	var author entities.Authors

	if err := row.Scan(&author.Id, &author.Name); err != nil {
		panic(err.Error())
	}

	return author
}

func Update(id int, author entities.Authors) bool {
	query, err := config.DB.Exec(`UPDATE authors SET name = ? where id = ?`, author.Name, id)
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
	_, err := config.DB.Exec("DELETE FROM authors WHERE id = ?", id)
	return err
}

package genremodel

import (
	"go-bookstore/config"
	"go-bookstore/entities"
)

func GetAll() []entities.Genres {
	rows, err := config.DB.Query(`SELECT * FROM genres`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var genres []entities.Genres

	for rows.Next() {
		var genre entities.Genres
		if err := rows.Scan(&genre.Id, &genre.Name); err != nil {
			panic(err)
		}

		genres = append(genres, genre)
	}

	return genres
}

func Create(genre entities.Genres) bool {
	result, err := config.DB.Exec(`
		INSERT INTO genres (name) VALUES (?)`,
		genre.Name,
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

func Detail(id int) entities.Genres {
	row := config.DB.QueryRow(`SELECT id, name FROM genres WHERE id = ? `, id)

	var genre entities.Genres

	if err := row.Scan(&genre.Id, &genre.Name); err != nil {
		panic(err.Error())
	}

	return genre
}

func Update(id int, genre entities.Genres) bool {
	query, err := config.DB.Exec(`UPDATE genres SET name = ? where id = ?`, genre.Name, id)
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
	_, err := config.DB.Exec("DELETE FROM genres WHERE id = ?", id)
	return err
}

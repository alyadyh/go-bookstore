package publishermodel

import (
	"go-bookstore/config"
	"go-bookstore/entities"
)

func GetAll() []entities.Publishers {
	rows, err := config.DB.Query(`SELECT * FROM publishers`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var publishers []entities.Publishers

	for rows.Next() {
		var publisher entities.Publishers
		if err := rows.Scan(&publisher.Id, &publisher.Name); err != nil {
			panic(err)
		}

		publishers = append(publishers, publisher)
	}

	return publishers
}

func Create(publisher entities.Publishers) bool {
	result, err := config.DB.Exec(`
		INSERT INTO publishers (name) VALUES (?)`,
		publisher.Name,
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

func Detail(id int) entities.Publishers {
	row := config.DB.QueryRow(`SELECT id, name FROM publishers WHERE id = ? `, id)

	var publisher entities.Publishers

	if err := row.Scan(&publisher.Id, &publisher.Name); err != nil {
		panic(err.Error())
	}

	return publisher
}

func Update(id int, publisher entities.Publishers) bool {
	query, err := config.DB.Exec(`UPDATE publishers SET name = ? where id = ?`, publisher.Name, id)
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
	_, err := config.DB.Exec("DELETE FROM publishers WHERE id = ?", id)
	return err
}

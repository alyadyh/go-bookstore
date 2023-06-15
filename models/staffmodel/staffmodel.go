package staffmodel

import (
	"go-bookstore/config"
	"go-bookstore/entities"
)

func GetAll() []entities.Staff {
	rows, err := config.DB.Query(`SELECT * FROM staff`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var staffs []entities.Staff

	for rows.Next() {
		var staff entities.Staff
		if err := rows.Scan(&staff.Id, &staff.Name, &staff.Email); err != nil {
			panic(err)
		}

		staffs = append(staffs, staff)
	}

	return staffs
}

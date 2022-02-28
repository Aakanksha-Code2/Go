package stores

import (
	"database/sql"
	"errors"

	//	"fmt"

	"github.com/aakanksha/Crud/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{
		db: db,
	}
}
func Insert(emp1 models.Emp, db *sql.DB) (int64, error) {
	inserted, err := db.Exec("INSERT INTO employee (id, name, email, role) VALUES (?, ?, ?, ?)", emp1.Id, emp1.Name, emp1.Email, emp1.Role)
	if err != nil {
		return 0, errors.New("some error")
	}
	id, err := inserted.LastInsertId()
	// if err != nil {
	// 	return 0, errors.New("some error")
	// }
	return id, nil
}

func Update(emp1 models.Emp, db *sql.DB) error {
	_, err := db.Exec("UPDATE employee SET name = ?, email=?, role=? WHERE ID = ?",
		&emp1.Name, &emp1.Email, &emp1.Role, &emp1.Id)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func Delete(idupdated int, db *sql.DB) error {
	deleted, err := db.Exec("DELETE FROM employee where id=?", idupdated)
	if err != nil {
		return errors.New("got some error")
	}
	_, err = deleted.RowsAffected()
	// if err != nil {
	// 	return errors.New("got some error")
	// }
	return nil
}

func (s store) Getbyid(idget int, db *sql.DB) (models.Emp, error) {

	var emp models.Emp
	//get, err := db.Exec("select *from employee where id = ?", idget)
	row := s.db.QueryRow("SELECT * FROM employee WHERE id = ?", idget)
	err := row.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Role)
	if err != nil {
		return emp, sql.ErrNoRows
	}
	return emp, nil
}

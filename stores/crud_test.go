package stores

import (
	//	"errors"

	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aakanksha/Crud/models"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(2, "aakanksha2", "ak@gmail.com", "sde")
	testcases := []struct {
		id          int
		user        models.Emp
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          2,
			user:        models.Emp{2, "aakanksha2", "ak@gmail.com", "sde"},
			mockQuery:   mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(2, "aakanksha2", "ak@gmail.com", "sde")),
			expectError: nil,
		},
		{
			id:          44,
			mockQuery:   mock.ExpectQuery("SELECT * FROM employee WHERE id = ?").WithArgs(44).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
		},
	}

	s := New(db)

	for _, testCase := range testcases {
		t.Run("", func(t *testing.T) {
			user, err := s.Getbyid(testCase.id, db)
			log.Print(err)
			//if err != nil && err.Error() != testCase.expectError.Error() {
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got %v ", testCase.expectError, err)
			}
			if err == nil && !reflect.DeepEqual(user, testCase.user) {
				t.Errorf("expected user %v,got %v", testCase.user, user)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testcases := []struct {
		id          int
		user        models.Emp
		output      int64
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          3,
			user:        models.Emp{3, "aakanksha3", "ak3@gmail.com", "sde3"},
			mockQuery:   mock.ExpectExec("INSERT INTO employee (id, name, email, role) VALUES (?, ?, ?, ?)").WillReturnError(errors.New("some error")),
			expectError: errors.New("some error"),
		},
		{
			id:          4,
			user:        models.Emp{4, "aakanksha4", "ak4@gmail.com", "sde4"},
			output:      4,
			mockQuery:   mock.ExpectExec("INSERT INTO employee (id, name, email, role) VALUES (?, ?, ?, ?)").WithArgs(4, "aakanksha4", "ak4@gmail.com", "sde4").WillReturnResult(sqlmock.NewResult(4, 1)),
			expectError: nil,
		},
	}

	for _, testCase := range testcases {
		t.Run("", func(t *testing.T) {

			output, err := Insert(testCase.user, db)
			log.Print(err)

			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error :%v, got %v ", testCase.expectError, err)
			}
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("expected user %v,got %v", testCase.output, output)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testcases := []struct {
		id int
		//output      int64
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          3,
			mockQuery:   mock.ExpectExec("DELETE FROM employee where id=?").WillReturnError(errors.New("got some error")),
			expectError: errors.New("got some error"),
		},
		{
			id:          4,
			mockQuery:   mock.ExpectExec("DELETE FROM employee where id=?").WithArgs(4).WillReturnResult(sqlmock.NewResult(1, 1)),
			expectError: nil,
		},
	}
	for _, testCase := range testcases {
		t.Run("", func(t *testing.T) {

			err := Delete(testCase.id, db)
			log.Print(err)
			// if err != nil && err.Error() != testCase.expectError.Error() {
			// 	t.Errorf("expected error =%v, got %v ", testCase.expectError, err)
			// }
			if !reflect.DeepEqual(err, testCase.expectError) {
				t.Errorf("expected user %v, %v", testCase.expectError, err)
			}
		})
	}
}

// func TestUpdate(t *testing.T) {
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)

// 	}
// 	defer db.Close()
// 	updateinfo := errors.New("updated")
// 	//rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(3, "aakanksha3", "ak3@gmail.com", "sde3")
// 	testcases := []struct {
// 		id          int
// 		user        models.Emp
// 		output      int64
// 		mockQuery   interface{}
// 		expectError error
// 		//res         int
// 	}{
// 		{
// 			id:        3,
// 			user:      models.Emp{3, "aakanksha3", "ak3@gmail.com", "sde3"},
// 			mockQuery: mock.ExpectExec("UPDATE employee set  name = ?, email = ?, role = ? where id = ?").WithArgs(3, "aakanksha3", "ak3@gmail.com", "sde3").WillReturnError(updateinfo),
// 			//inserted, err := db.Exec("INSERT INTO employee (id, name, email, role) VALUES (?, ?, ?, ?)", emp1.Id, emp1.Name, emp1Email, emp1.Role)
// 			//if res==models.Emp.id
// 			expectError: nil,
// 		},
// 		// {
// 		// 	id:   4,
// 		// 	user: models.Emp{4, "aakanksha4", "ak4@gmail.com", "sde4"},
// 		// 	//output:      4,
// 		// 	mockQuery:   mock.ExpectQuery("INSERT INTO employee (id, name, email, role) VALUES (?, ?, ?, ?)").WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(4, "aakanksha4", "ak4@gmail.com", "sde4")),
// 		// 	expectError: nil,
// 		// },
// 	}

// 	//s := New(db)

// 	for _, testCase := range testcases {
// 		t.Run("", func(t *testing.T) {

// 			uid, err := Update(testCase.id, testCase.user, db)
// 			//log.Print(err)
// 			//if err != nil && err.Error() != testCase.expectError.Error() {
// 			// if err != nil && err.Error() != testCase.expectError.Error() {
// 			// 	t.Errorf("expected error :%v, got %v ", testCase.expectError, err)
// 			// }
// 			if !reflect.DeepEqual(uid, testCase.id) {
// 				t.Errorf("expected user %v,got %v", testCase.id, uid)
// 			}
// 			if !reflect.DeepEqual(err, testCase.output) {
// 				t.Errorf("expected user %v,got %v", testCase.expectError, err)
// 			}

// 			// if err != nil && err.Error() != testCase.expectError.Error() {
// 			// 	t.Errorf("expected error :%v, got %v ", testCase.expectError, err)
// 			// }
// 			// if !reflect.DeepEqual(user, testCase.user) {
// 			// 	t.Errorf("expected user %v,got %v", testCase.user, user)
// 			// }
// 		})
// 	}
// }

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}
	updaterr := errors.New("update failed")

	tests := []struct {
		desc      string
		expecterr error
		input_emp models.Emp
		mockCall  *sqlmock.ExpectedExec
	}{
		{
			desc:      "update succes",
			expecterr: nil,
			input_emp: models.Emp{1, "Aakanksha", "ak@gmail.com", "Java"},
			mockCall:  mock.ExpectExec("UPDATE employee SET name = ?, email=?, role=? WHERE ID = ?").WithArgs("Aakanksha", "ak@gmail.com", "Java", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:      "update fail",
			expecterr: updaterr,
			input_emp: models.Emp{2, "", "", ""},
			mockCall:  mock.ExpectExec("UPDATE employee SET name=?,email=?,role=? WHERE ID = ?").WithArgs("", "", "", 2).WillReturnError(updaterr),
		},
	}

	for _, tc := range tests {
		err := Update(tc.input_emp, db)

		if !reflect.DeepEqual(err, tc.expecterr) {
			t.Errorf("Expected: %v, Got: %v", tc.expecterr, err)
		}

	}

}

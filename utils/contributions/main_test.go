package contributions

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetContributions(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	cols := []string{"id", "user_id", "title"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, 1, "foo"))

	_, err := GetContributions(db, 1)

	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}

}

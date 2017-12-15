package tags

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetTags(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	cols := []string{"id", "user_contribution_id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, 1, "foo"))

	_, err := GetTags(db, 1)

	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
}

func TestMapTgas(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	cols := []string{"id", "user_contribution_id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, 1, "foo"))

	contributions := []types.UserContribution{
		{
			ID: 1,
		},
	}
	_, err := MapTgas(db, contributions)

	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
}

func TestMapTgasOnFailure(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	cols := []string{"id", "user_contribution_id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, 1, "foo")).WillReturnError(fmt.Errorf("some error"))

	contributions := []types.UserContribution{
		{
			ID: 1,
		},
	}
	_, err := MapTgas(db, contributions)

	if err == nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
}

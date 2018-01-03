package contributionsDetail

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	defer db.Close()

	cols := []string{"id", "user_contribution_id", "body"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, 1, "foo"))

	_, err := GetByID(db, 1)

	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
}

func TestGetBody(t *testing.T) {
	_, err := GetBody(`[{"priority":1,"body":"foo","iconType":1,"iconFace":1,"directionType":1,"talkType":1,"edit":false,"character":{"id":1,"fileName":"1.jpg","voiceType":1}}]`)

	if err != nil {
		t.Fatalf("An error '%s' was not expecting", err)
	}
}

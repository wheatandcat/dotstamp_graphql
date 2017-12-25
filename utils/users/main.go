package users

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
	date "github.com/wheatandcat/dotstamp_graphql/utils/date.go"
	"github.com/wheatandcat/dotstamp_graphql/utils/encryption"
)

// GetByEmail get user by email
func GetByEmail(DB *sqlx.DB, email string) (types.UserMaster, error) {
	u := types.UserMaster{}

	err := DB.Get(&u,
		`
		SELECT *
    FROM user_masters
    WHERE email = ?
    AND deleted_at IS NULL
		`, email)

	return u, err
}

// Create ceate user
func Create(DB *sqlx.DB, email string, password string, k string) (types.UserMaster, error) {
	u, err := GetByEmail(DB, email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return u, err
	}

	if u.ID != uint(0) {
		return u, errors.New("Is already registered e-mail address")
	}

	result, err := DB.Exec(
		`
		INSERT INTO user_masters
		(name, email, password, profile_image_id, created_at, updated_at)
		VALUES (?, ?, ?, 0, '`+date.Now()+`', '`+date.Now()+`')
		`, email, email, encryption.GetPassword(password, k))

	if err != nil {
		return u, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return u, err
	}

	r := types.UserMaster{
		ID:    uint(id),
		Email: email,
		Name:  email,
	}

	return r, nil
}

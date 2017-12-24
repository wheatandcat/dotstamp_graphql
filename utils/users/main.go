package users

import (
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
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

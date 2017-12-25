package login

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
	"github.com/wheatandcat/dotstamp_graphql/utils/encryption"
	"github.com/wheatandcat/dotstamp_graphql/utils/users"
)

// GetLogin get login user
func GetLogin(DB *sqlx.DB, email string, p string, k string) (types.UserMaster, error) {

	u, err := users.GetByEmail(DB, email)
	if err != nil {
		return u, err
	}

	if encryption.GetPassword(p, k) != u.Password {
		return u, errors.New("password diffrent")
	}

	return u, nil
}

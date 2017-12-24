package login

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
	"github.com/wheatandcat/dotstamp_graphql/utils/users"
	"github.com/wheatandcat/dotstamp_server/utils"
)

func getPassword(p string, k string) string {
	return utils.SrringToEncryption(p + k)
}

// GetLogin get login user
func GetLogin(DB *sqlx.DB, email string, p string, k string) (types.UserMaster, error) {

	u, err := users.GetByEmail(DB, email)
	if err != nil {
		return u, err
	}

	if getPassword(p, k) != u.Password {
		return u, errors.New("password diffrent")
	}

	return u, nil
}

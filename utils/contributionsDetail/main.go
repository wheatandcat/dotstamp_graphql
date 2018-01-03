package contributionsDetail

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
)

// GetByID get user_contribution_details by id
func GetByID(DB *sqlx.DB, id int) (types.UserContributionDetail, error) {
	u := types.UserContributionDetail{}

	err := DB.Get(&u,
		`
		SELECT *
    FROM user_contribution_details
    WHERE user_contribution_id = ?
    AND deleted_at IS NULL
		`, id)

	return u, err
}

// GetBody get body array
func GetBody(body string) (b []types.Body, err error) {
	bytes := []byte(body)
	err = json.Unmarshal(bytes, &b)

	return b, err
}

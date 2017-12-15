package follows

import (
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
)

// GetFollows get contribution follows count
func GetFollows(DB *sqlx.DB, id uint) (int, error) {
	r := 0

	err := DB.Get(&r,
		`
		SELECT COUNT(*) as count
    FROM user_contribution_follows
    WHERE user_contribution_id = ?
    AND deleted_at IS NULL
		`, id)

	return r, err
}

// MapTgas mapping follow
func MapTgas(DB *sqlx.DB, contributions []types.UserContribution) ([]types.UserContribution, error) {
	for k, c := range contributions {
		follow, err := GetFollows(DB, c.ID)
		if err != nil {
			return contributions, err
		}
		contributions[k].Follow = follow
	}

	return contributions, nil
}

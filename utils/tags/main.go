package tags

import (
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
)

// GetTags get contributions
func GetTags(DB *sqlx.DB, id uint) ([]types.UserContributionTag, error) {
	u := []types.UserContributionTag{}
	err := DB.Select(&u,
		`
		SELECT *
		FROM user_contribution_tags
		WHERE user_contribution_id = ? AND deleted_at IS NULL
		`, id)

	return u, err
}

// MapTgas mapping tgs
func MapTgas(DB *sqlx.DB, contributions []types.UserContribution) ([]types.UserContribution, error) {
	for k, c := range contributions {
		tags, err := GetTags(DB, c.ID)
		if err != nil {
			return contributions, err
		}
		contributions[k].Tags = tags
	}

	return contributions, nil
}

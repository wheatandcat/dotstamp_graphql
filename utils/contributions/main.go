package contributions

import (
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
)

// GetContributions get contributions
func GetContributions(DB *sqlx.DB, first int) ([]types.UserContribution, error) {
	u := []types.UserContribution{}
	err := DB.Select(&u,
		`
		SELECT
			user_contributions.id as id,
	    user_contributions.title as title,
	    user_contributions.view_status as view_status,
		  user_contributions.user_id as user_id,
	    user_masters.name as user_name,
	    user_masters.profile_image_id as profile_image_id,
	    user_contributions.created_at as created_at,
	    user_contributions.updated_at as updated_at
		FROM user_contributions
			INNER JOIN user_masters ON user_contributions.user_id = user_masters.id
		WHERE user_contributions.deleted_at IS NULL
		ORDER BY user_contributions.id DESC LIMIT ?
		`, first)

	return u, err
}

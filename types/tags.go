package types

import "github.com/graphql-go/graphql"

// UserContributionTag タグ
type UserContributionTag struct {
	ID                 uint    `db:"id" json:"id"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	Name               string  `db:"name" json:"name"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}

// ContributionTagType 投稿タグタイプ
var ContributionTagType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tag",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "name",
			},
		},
	},
)

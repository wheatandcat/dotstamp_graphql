package types

import "github.com/graphql-go/graphql"

// UserContributionFollow フォロー
type UserContributionFollow struct {
	ID                 uint    `db:"id" json:"id"`
	UserID             int     `db:"user_id" json:"userId"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}

// FollowType contribution follow Type
var FollowType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Follow",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"userContributionID": &graphql.Field{
				Type:        graphql.Int,
				Description: "contribution id",
			},
			"userId": &graphql.Field{
				Type:        graphql.Int,
				Description: "user id",
			},
			"createdAt": &graphql.Field{
				Type:        graphql.String,
				Description: "created date",
			},
			"updatedAt": &graphql.Field{
				Type:        graphql.String,
				Description: "update date",
			},
		},
	},
)

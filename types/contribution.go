package types

import (
	"github.com/graphql-go/graphql"
)

// UserContribution user contribution
type UserContribution struct {
	ID             uint                  `db:"id" json:"id"`
	Title          string                `db:"title" json:"title"`
	ViewStatus     int                   `db:"view_status" json:"viewStatus"`
	UserID         int                   `db:"user_id" json:"userId"`
	UserName       string                `db:"user_name" json:"userName"`
	ProfileImageID string                `db:"profile_image_id" json:"profileImageId"`
	CreatedAt      string                `db:"created_at" json:"createdAt"`
	UpdatedAt      string                `db:"updated_at" json:"updatedAt"`
	DeletedAt      *string               `db:"deleted_at" json:"deletedAt"`
	Tags           []UserContributionTag `json:"tags"`
	Follow         int                   `json:"follow"`
}

// ContributionType contribution Type
var ContributionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Contribution",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "title",
			},
			"viewStatus": &graphql.Field{
				Type:        graphql.Int,
				Description: "view status 0 ore 1",
			},
			"userId": &graphql.Field{
				Type:        graphql.Int,
				Description: "user id",
			},
			"userName": &graphql.Field{
				Type:        graphql.String,
				Description: "user name",
			},
			"profileImageId": &graphql.Field{
				Type:        graphql.Int,
				Description: "profile image id",
			},
			"createdAt": &graphql.Field{
				Type:        graphql.String,
				Description: "created date",
			},
			"updatedAt": &graphql.Field{
				Type:        graphql.String,
				Description: "update date",
			},
			"tags": &graphql.Field{
				Type:        graphql.NewList(TagType),
				Description: "tags",
			},
			"follow": &graphql.Field{
				Type:        graphql.Int,
				Description: "follow count",
			},
		},
	},
)

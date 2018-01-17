package types

import "github.com/graphql-go/graphql"

// UserContributionMovie 動画
type UserContributionMovie struct {
	ID                 uint    `db:"id" json:"id"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	MovieType          int     `db:"movie_type" json:"movieType"`
	MovieID            string  `db:"movie_id" json:"movieId"`
	MovieStatus        int     `db:"movie_status" json:"movieStatus"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}

// MovieType 動画タイプ
var MovieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"movieType": &graphql.Field{
				Type:        graphql.Int,
				Description: "movie type",
			},
			"movieId": &graphql.Field{
				Type:        graphql.String,
				Description: "movie id",
			},
			"movieStatus": &graphql.Field{
				Type:        graphql.Int,
				Description: "movie status",
			},
		},
	},
)

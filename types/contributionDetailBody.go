package types

import (
	"github.com/graphql-go/graphql"
)

// UserContributionDetail user contribution detail
type UserContributionDetail struct {
	ID                 uint    `db:"id" json:"id"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	Body               string  `db:"body" json:"body"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}

// BodyCharacter body character type
type BodyCharacter struct {
	ID        int    `json:"id"`
	FileName  string `json:"fileName"`
	VoiceType int    `json:"voiceType"`
}

// BodyCharacterType body character type
var BodyCharacterType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BodyCharacter",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "id",
			},
			"fileName": &graphql.Field{
				Type:        graphql.String,
				Description: "file name",
			},
			"voiceType": &graphql.Field{
				Type:        graphql.Int,
				Description: "voice type",
			},
		},
	},
)

// Body user contribution detail body
type Body struct {
	Priority      int           `json:"priority"`
	Body          string        `json:"body"`
	Character     BodyCharacter `json:"character"`
	DirectionType int           `json:"directionType"`
	TalkType      int           `json:"talkType"`
}

// BodyType contribution Type
var BodyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Body",
		Fields: graphql.Fields{
			"priority": &graphql.Field{
				Type:        graphql.Int,
				Description: "priority",
			},
			"body": &graphql.Field{
				Type:        graphql.String,
				Description: "body",
			},
			"character": &graphql.Field{
				Type:        BodyCharacterType,
				Description: "character",
			},
			"directionType": &graphql.Field{
				Type:        graphql.Int,
				Description: "direction",
			},
			"talkType": &graphql.Field{
				Type:        graphql.Int,
				Description: "talk",
			},
		},
	},
)

package types

import (
	"github.com/graphql-go/graphql"
)

// Login ログイン
type Login struct {
	Email    uint `json:"email"`
	Password uint `json:"password"`
}

// LoginType ログインタイプ
var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "email",
			},
			"password": &graphql.Field{
				Type:        graphql.String,
				Description: "password",
			},
		},
	},
)

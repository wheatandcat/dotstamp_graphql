package types

import (
	"github.com/graphql-go/graphql"
)

// Login ログイン
type Login struct {
	Email    uint `json:"email"`
	Password uint `json:"password"`
}

// AuthKey 認証キー
type AuthKey struct {
	Key string `json:"key"`
}

// LoginType ログインタイプ
var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"key": &graphql.Field{
				Type:        graphql.String,
				Description: "auth key",
			},
		},
	},
)

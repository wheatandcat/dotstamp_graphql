package authJwt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthID auth user yype
type AuthID struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

// CreateTokenString create token
func CreateTokenString(userID uint, key string) (string, error) {
	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &AuthID{
		ID: userID,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte(key))
	if err != nil {
		return tokenstring, err
	}

	return tokenstring, nil
}

// Auth auth by tokenstring
func Auth(tokenstring string, key string) (uint, error) {
	_, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	user := AuthID{}
	_, err = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	return user.ID, err
}

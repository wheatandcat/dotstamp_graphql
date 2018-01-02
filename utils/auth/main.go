package authJwt

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthID auth user yype
type AuthID struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

// CreateTokenString create token
func CreateTokenString(userID uint) (string, error) {
	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &AuthID{
		ID: userID,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		return tokenstring, err
	}

	return tokenstring, nil
}

// Auth auth by tokenstring
func Auth(tokenstring string) {
	// サーバだけが知り得るSecretでこれをParseする
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	// Parseメソッドを使うと、Claimsはmapとして得られる
	log.Println(token.Claims, err)

	// 別例, jwt.StandardClaimsを満たすstructに直接decodeさせることもできる
	user := AuthID{}
	token, err = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})

	log.Println(token.Valid, user, err)
}

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"time"
)

var jwtKey []byte

type Claims struct {
	Name string
	jwt.StandardClaims
}

func GenerateJWT(name string, uid int, audience string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.OneDay)

	claims := &Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Audience:  audience,
			IssuedAt:  nowTime.Unix(),
			Issuer:    consts.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

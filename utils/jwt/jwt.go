package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"strconv"
	"time"
)

// JWTKey TODO: Delete this and use config.yaml
var JWTKey = []byte("hdu-cxsj1JWTKEY")

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
			Id:        strconv.Itoa(uid),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

package util

import "github.com/golang-jwt/jwt/v4"

type JwtAuth struct {
	accessSecret string
	accessExpire int64
}

func NewJwtAuth(accessSecret string, accessExpire int64) *JwtAuth {
	return &JwtAuth{
		accessSecret: accessSecret,
		accessExpire: accessExpire,
	}
}
func (j *JwtAuth) GenerateToken(now int64, userId int64) (string, error) {
	claims := jwt.MapClaims{
		"iat":    now,
		"exp":    now + j.accessExpire,
		"userId": userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.accessSecret))
}

package common

import (
	"os"

	"github.com/bagusyanuar/go-deltomed/exception"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Unique   uuid.UUID `json:"unique"`
	Username string    `json:"username"`
	Role     string    `json:"roles"`
}

type JWTSignReturn struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}

func GenerateAccessToken(jwtSign *JWTSignReturn) (accessToken string, err error) {
	issuer := os.Getenv("JWT_ISSUER")
	jwtSignaturKey := os.Getenv("JWT_SIGNATURE_KEY")
	JWTSigninMethod := jwt.SigningMethodHS256
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: issuer,
		},
		Unique:   jwtSign.ID,
		Username: jwtSign.Username,
		Role:     jwtSign.Role,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	return token.SignedString([]byte(jwtSignaturKey))
}

func ValidateToken(authorization string) (interface{}, error) {
	if authorization == "" {
		return nil, exception.ErrorUnauthorized
	}

	bearer := string(authorization[0:7])
	tokenString := string(authorization[7:])
	JWTSignatureKey := os.Getenv("JWT_SIGNATURE_KEY")

	if bearer != "Bearer " {
		return nil, exception.ErrorUnauthorized
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTSignatureKey), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, exception.ErrorUnauthorized
}

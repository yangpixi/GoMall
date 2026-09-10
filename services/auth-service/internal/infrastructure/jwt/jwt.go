package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey         string
	expiration        time.Duration
	refreshExpiration time.Duration
}

type claims struct {
	RoleIDs []uint `json:"role_ids"`
	jwt.RegisteredClaims
}

func NewJWT(secretKey string, expiration, refreshExpiration time.Duration) (*JWT, error) {
	if secretKey == "" {
		return nil, errors.New("invalid secret key")
	}

	return &JWT{
		secretKey:         secretKey,
		expiration:        expiration,
		refreshExpiration: refreshExpiration,
	}, nil
}

func (J *JWT) Expire() time.Duration {
	return J.expiration
}

func (J *JWT) RefreshExpire() time.Duration {
	return J.refreshExpiration
}

func (J *JWT) Token(username string, roleIDs []uint) (string, error) {
	token, err := J.generateToken(username, roleIDs, J.expiration)
	return token, err
}

func (J *JWT) RefreshToken(username string, roleIDs []uint) (string, error) {
	token, err := J.generateToken(username, roleIDs, J.refreshExpiration)
	return token, err
}

// generateToken internal method to generate token
func (J *JWT) generateToken(username string, roleIDs []uint, expiration time.Duration) (string, error) {
	c := &claims{
		RoleIDs: roleIDs,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GoMall",
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString([]byte(J.secretKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

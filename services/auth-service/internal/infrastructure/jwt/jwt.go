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
	Kind    string `json:"kind"`
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

func (J *JWT) Token(username string, roleIDs []uint) (string, int, error) {
	token, err := J.generateToken(username, roleIDs, J.expiration, "access")
	return token, int(J.expiration.Seconds()), err
}

func (J *JWT) RefreshToken(username string, roleIDs []uint) (string, error) {
	token, err := J.generateToken(username, roleIDs, J.refreshExpiration, "refresh")
	return token, err
}

// generateToken internal method to generate token
func (J *JWT) generateToken(username string, roleIDs []uint, expiration time.Duration, kind string) (string, error) {
	c := &claims{
		RoleIDs: roleIDs,
		Kind:    kind,
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

func (J *JWT) ParseAndValidate(token string) (username, kind string, roleIDs []uint, err error) {
	res, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (any, error) {
		return []byte(J.secretKey), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithLeeway(5*time.Second),
	)

	if err != nil {
		return "", "", nil, err
	}

	if c, ok := res.Claims.(*claims); ok {
		return c.Subject, c.Kind, c.RoleIDs, nil
	}

	// this won't happen in most common ways
	return "", "", nil, nil
}

package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lakhan-purohit/net-http/internal/pkg/config"
)

type Claims struct {
	UserID int64  `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role,omitempty"`
	UUID   string `json:"uuid,omitempty"`
	Type   string `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

type Service struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWT() *Service {
	cfg := config.Get()

	// TTL = Time To Live (Expiration Duration)
	return &Service{
		secret:     []byte(cfg.JWT.Secret),
		accessTTL:  cfg.JWT.AccessExpiration,  // Taken from JWT_ACCESS_EXPIRES_IN (e.g., "1h")
		refreshTTL: cfg.JWT.RefreshExpiration, // Taken from JWT_REFRESH_EXPIRES_IN (e.g., "168h")
	}
}

func (s *Service) Generate(claims Claims) (accessToken string, refreshToken string, err error) {
	// now is the current server time used as a base for tokens
	now := time.Now()

	// --------------------
	// Access Token (short)
	// --------------------
	accessClaims := Claims{
		UserID: claims.UserID,
		Email:  claims.Email,
		Role:   claims.Role,
		UUID:   claims.UUID,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(s.secret)
	if err != nil {
		return
	}

	// --------------------
	// Refresh Token (long)
	// --------------------
	refreshClaims := Claims{
		UserID: claims.UserID,
		UUID:   claims.UUID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTTL)),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(s.secret)
	return
}

func (s *Service) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

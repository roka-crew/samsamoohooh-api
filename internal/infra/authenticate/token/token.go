package token

import (
	"fmt"
	"time"

	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/infra/config"
	"samsamoohooh-api/pkg/httperr"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	config *config.Config
}

func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}

func (s *TokenService) GenerateToken(userID int, per domain.Permission, kind domain.Kind) (string, error) {
	now := time.Now()
	var expTime time.Time

	// 토큰 종류에 따라 만료 시간 설정
	switch kind {
	case domain.KindAccess:
		expTime = now.Add(s.config.Token.AccessExp)
	case domain.KindRefresh:
		expTime = now.Add(s.config.Token.RefreshExp)
	default:
		return "", httperr.New().
			SetType(httperr.ServerInternalError).
			SetDetail("invalid token kind: %s", kind)
	}

	claims := domain.Token{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now.Add(s.config.Token.NotBefore)),
		},
		Per:    per,
		Kind:   kind,
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Token.SecretKey))
}

func (s *TokenService) ParseToken(tokenString string) (*domain.Token, error) {
	jwt, err := jwt.ParseWithClaims(tokenString, &domain.Token{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.config.Token.SecretKey, nil
	})
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.AuthFailed).
			SetDetail("failed to parse token, inspect: %v", err)
	}

	token, ok := jwt.Claims.(*domain.Token)
	if !ok {
		return nil, httperr.New().
			SetType(httperr.AuthFailed).
			SetDetail("unknown claims type, cannot proceed")
	}

	return token, nil
}

func (s *TokenService) RefreshToken(tokenString string) (string, error) {
	jwt, err := s.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// Refresh 토큰인지 확인
	if jwt.Kind != domain.KindRefresh {
		return "", httperr.New().
			SetType(httperr.AuthFailed).
			SetDetail("invalid token kind: expected refresh token")
	}

	// 새로운 Access 토큰 생성
	return s.GenerateToken(jwt.UserID, jwt.Per, domain.KindAccess)
}

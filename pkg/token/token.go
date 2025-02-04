package token

import (
	"samsamoohooh-api/pkg/config"
	"samsamoohooh-api/pkg/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Kind string

const (
	KindAccess  Kind = "ACCESS"
	KindRefresh Kind = "REFRESH"
)

type Permission string

const (
	PermissionStaff = "STAFF"
	PermissionUser  = "USER"
)

type GenerateTokenParams struct {
	Kind   Kind
	Per    Permission
	UserID int
}

type Token interface {
	GenerateToken(params GenerateTokenParams) (string, error)
	ParseToken(tokenString string) (Payload, error)
	RefreshToken(tokenString string) (string, error)
}

const (
	AccessExp  = 15 * time.Minute
	RefreshExp = 60 * 24 * time.Hour
	NotBefore  = 0
)

type token struct {
	cfg *config.Config
}

func New(cfg *config.Config) Token {
	return &token{cfg: cfg}
}

func (t token) GenerateToken(params GenerateTokenParams) (string, error) {
	now := time.Now()

	var expTime time.Time
	switch params.Kind {
	case KindAccess:
		expTime = now.Add(AccessExp)
	case KindRefresh:
		expTime = now.Add(RefreshExp)
	default:
		return "", errors.New("invalid kind")
	}

	payload := Payload{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now.Add(NotBefore)),
		},
		Per:    params.Per,
		Kind:   params.Kind,
		UserID: params.UserID,
	}

	secretKey := []byte(t.cfg.JWT.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(secretKey)
}

func (t token) ParseToken(tokenString string) (Payload, error) {
	secretKey := []byte(t.cfg.JWT.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return Payload{}, err
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return Payload{}, errors.New("invalid token")
	}

	return *payload, nil
}

func (t token) RefreshToken(tokenString string) (string, error) {
	payload, err := t.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	if payload.Kind != KindRefresh {
		return "", errors.New("token is not refresh token")
	}

	return t.GenerateToken(GenerateTokenParams{
		Kind:   KindAccess,
		Per:    payload.Per,
		UserID: payload.UserID,
	})
}

type Payload struct {
	jwt.RegisteredClaims

	// Kind (토큰의 종류)
	Kind Kind `json:"kind"`

	// 사용자의 ID
	UserID int `json:"userID"`

	// Permission (권한)
	Per Permission `json:"per"`
}

func (p Payload) Validate() error {
	switch p.Kind {
	case KindAccess, KindRefresh:
	default:
		return errors.New("invalid kind")
	}

	switch p.Per {
	case PermissionStaff, PermissionUser:
	default:
		return errors.New("invalid permission")
	}

	return nil
}

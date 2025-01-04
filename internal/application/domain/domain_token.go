package domain

import (
	"errors"

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

type Token struct {
	jwt.RegisteredClaims

	// Permission (권한)
	Per Permission `json:"per"`

	// Kind (토큰의 종류)
	Kind Kind `json:"kind"`

	// 사용자의 ID
	UserID int `json:"userID"`
}

func (t Token) Validate() error {
	// 3. Validate token kind
	switch t.Kind {
	case KindAccess, KindRefresh:
	default:
		return errors.New("invalid token kind")
	}

	// 4. Validate permission
	switch t.Per {
	case PermissionStaff, PermissionUser:
	default:
		return errors.New("invalid permission")
	}

	return nil
}

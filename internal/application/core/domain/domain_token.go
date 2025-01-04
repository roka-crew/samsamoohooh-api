package domain

import "time"

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
	// Expiration Time (만료시간)
	Exp time.Time

	// Issued At (발급시간)
	Ita time.Time

	// Not Before (토큰이 유효해지기 시작하는 시간)
	Nbf time.Time

	// Subject (토큰의 주체)
	Sub int

	// Permission (권한)
	Per Permission

	// Kind (토큰의 종류)
	Kind Kind
}

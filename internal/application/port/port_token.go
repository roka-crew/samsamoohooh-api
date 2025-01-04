package port

import "samsamoohooh-api/internal/application/domain"

type TokenService interface {
	GenerateToken(sub int, per domain.Permission, kind domain.Kind) // 토큰 생성
	ParseToken(tokenString string) (*domain.Token, error)           // 토큰 파싱(유효성 검증도 진행)
	RefreshToken(tokenString string) (string, error)                // 토큰 갱신
}

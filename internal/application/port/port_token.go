package port

import "samsamoohooh-api/internal/application/domain"

type TokenService interface {
	GenerateToken(sub int, per domain.Permission, kind domain.Kind) (string, error) // 토큰 생성
	ValidateToken(tokenString string) (bool, error)                                 // 토큰 유효성 검증
	ParseToken(tokenString string) (*domain.Token, error)                           // 토큰 파싱 및 클레임 추출
	RefreshToken(tokenString string) (string, error)                                // 토큰 갱신
}

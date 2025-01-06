package presenter

import "samsamoohooh-api/internal/application/domain"

type FindUserResponse struct {
	Nickname   string          `json:"nickname"`
	Resolution *string         `json:"resolution"`
	Provider   domain.Provider `json:"provider"`
}

func NewFindUserResponse(user *domain.User) *FindUserResponse {
	return &FindUserResponse{
		Nickname:   user.Nickname,
		Resolution: user.Resolution,
		Provider:   user.Provider,
	}
}

type PatchUserResponse struct {
	Nickname   string  `json:"nickname,omitempty"`
	Resolution *string `json:"resolution,omitempty"`
}

func NewPatchUserResponse(user *domain.User) *PatchUserResponse {
	return &PatchUserResponse{
		Nickname:   user.Nickname,
		Resolution: user.Resolution,
	}
}

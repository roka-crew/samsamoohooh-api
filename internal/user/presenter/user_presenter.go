package presenter

import "samsamoohooh-api/domain"

type FindUserByMeResponse struct {
	ID         int             `json:"id"`
	Nickname   string          `json:"nickname"`
	Resolution *string         `json:"resolution,omitemtpy"`
	Provider   domain.Provider `json:"provider"`
}

func (resp *FindUserByMeResponse) FromModel(user *domain.User) *FindUserByMeResponse {
	resp.ID = int(user.ID)
	resp.Nickname = user.Nickname
	resp.Resolution = user.Resolution
	resp.Provider = user.Provider
	return resp
}

package presenter

type FindUserRequest struct {
	ID int `params:"id" validate:"required,gte=1"`
}

func (r FindUserRequest) ToParams() *FoundUserParams {
	return &FoundUserParams{
		ID: r.ID,
	}
}

type PatchUserRequest struct {
	ID         int     `params:"id"       validate:"required,gte=1"`
	Nickname   *string `json:"nickname"   validate:"omitempty,min=3,max=12"`
	Resolution *string `json:"resolution" validate:"omitempty,min=0,max=18"`
}

func (r PatchUserRequest) ToParams() *PatchUserParams {
	return &PatchUserParams{
		ID:         r.ID,
		Nickname:   r.Nickname,
		Resolution: r.Resolution,
	}
}

package presenter

type FoundUserRequest struct {
	ID int `parmas:"id" validate:"required,gte=1"`
}

type PatchUserReqeust struct {
	ID         int    `params:"id"       validate:"required,gte=1"`
	Nickname   string `json:"nickname"   validate:"omitempty,min=3,max=12"`
	Resolution string `json:"resolutino" validate:"omitempty,min=0,max=18"`
}

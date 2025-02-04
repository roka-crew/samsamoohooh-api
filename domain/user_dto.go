package domain

type FindUserByMeRequest struct {
	RequestUserID int `json:"-" swaggerignore:"true"`
}

type PatchUserByMeRequest struct {
	RequestUserID int     `json:"-" swaggerignore:"true"`
	Nickname      *string `json:"nickname"`
	Resolution    *string `json:"resolution"`
}

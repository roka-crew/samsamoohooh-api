package presenter

import (
	"samsamoohooh-api/internal/application/domain"
)

type CreateUserParams struct {
	Nickname   string          `validate:"required,min=3,max=12"`
	Resolution *string         `validate:"omitempty,min=0,max=18"`
	Provider   domain.Provider `validate:"required,oneof=GOOGLE APPLE KAKAO"`
}

type FoundUserParams struct {
	UserID int `validate:"required,gte=1"`
}

type ListUsersParams struct {
	Offset *int `validate:"omitempty,gte=0"`
	Limit  int  `validate:"required,gte=0,lte=50"`
}

type PatchUserParams struct {
	UserID     int              `validate:"required,gte=1"`
	Nickname   *string          `validate:"omitempty,min=3,max=12"`
	Resolution *string          `validate:"omitempty,min=0,max=18"`
	Provider   *domain.Provider `validate:"omitempty,oneof=GOOGLE APPLE KAKAO"`
}

type DeleteUserParams struct {
	UserID int `validate:"required,gte=1"`
}

type GetUserGroupsParams struct {
	UserID int  `validate:"required,gte=1"`
	Offset *int `validate:"omitempty,gte=0"`
	Limit  int  `validate:"required,gte=0,lte=50"`
}

type GetUserTopicsParams struct {
	UserID int  `validate:"required,gte=1"`
	Offset *int `validate:"omitempty,gte=0"`
	Limit  int  `validate:"required,gte=0,lte=50"`
}

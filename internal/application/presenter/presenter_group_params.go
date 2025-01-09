package presenter

type CreateGroupParams struct {
	BookTitle        string  `validate:"required,min=1,max=255"`
	BookAuthor       string  `validate:"required,min=1,max=255"`
	BookPageMax      int     `validate:"required,gte=0"`
	BookPublisher    *string `validate:"omitempty,min=0,max=255"`
	BookIntroduction *string `validate:"omitempty,min=0,max=255"`
}

type FindGroupParams struct {
	GroupID int `validate:"required,gte=1"`
}

type ListGroupsParams struct {
	Offset *int `validate:"omitempty,gte=0"`
	Limit  int  `validate:"required,gte=0"`
}

type PatchGroupParams struct {
	GroupID          int     `validate:"required,gte=1"`
	BookTitle        *string `validate:"omitempty,min=1,max=255"`
	BookAuthor       *string `validate:"omitempty,min=1,max=255"`
	BookPageMax      *int    `validate:"omitempty,gte=0"`
	BookPublisher    *string `validate:"omitempty,min=0,max=255"`
	BookIntroduction *string `validate:"omitempty,min=0,max=255"`
}

type DeleteGroupParams struct {
	GroupID int `validate:"required,gte=1"`
}

type GetGroupUsersParams struct {
	GroupID int  `validate:"required,gte=1"`
	Offset  *int `validate:"omitempty,gte=0"`
	Limit   int  `validte:"required,gte=1"`
}

type GetGroupGoalsParams struct {
	GroupID int  `validate:"required,gte=1"`
	Offset  *int `validate:"omitempty,gte=0"`
	Limit   int  `validte:"required,gte=1"`
}

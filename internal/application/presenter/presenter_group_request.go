package presenter

type CreateGroupRequest struct {
	BookTitle        string  `json:"bookTitle" validate:"required,min=1,max=255"`
	BookAuthor       string  `json:"bookAuthor" validate:"required,min=1,max=255"`
	BookPageMax      int     `json:"bookPageMax" validate:"required,gte=0"`
	BookPublisher    *string `json:"bookPublisher" validate:"omitempty,min=0,max=255"`
	BookIntroduction *string `json:"bookIntroduction" validate:"omitempty,min=0,max=255"`
}

type FindGroupRequest struct {
	GroupID int `params:"id" validate:"required,gte=1"`
}

type ListGroupsRequest struct {
	Offset *int `json:"offset" validate:"omitempty,gte=0"`
	Limit  int  `json:"limit" validate:"required,gte=0"`
}

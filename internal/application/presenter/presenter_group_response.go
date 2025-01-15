package presenter

type CreateGroupResponse struct {
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPublisher    *string `json:"bookPublisher,omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty"`
}

type FindGroupResponse struct {
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPublisher    *string `json:"bookPublisher,omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty"`
}

type ListGroupsResponseItem struct {
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPublisher    *string `json:"bookPublisher,omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty"`
}

type ListGroupsResponse struct {
	Groups []ListGroupsResponseItem `json:"groups"`
}

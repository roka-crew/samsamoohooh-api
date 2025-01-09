package port

import (
	"context"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
)

type UserService interface {
	FindUser(ctx context.Context, req *presenter.FindUserRequest) (*domain.User, error)
	PatchUser(ctx context.Context, req *presenter.PatchUserRequest) (*domain.User, error)
}

type UserStore interface {
	CreateUser(ctx context.Context, params *presenter.CreateUserParams) (*domain.User, error)
	FindUser(ctx context.Context, params *presenter.FoundUserParams) (*domain.User, error)
	ListUsers(ctx context.Context, params *presenter.ListUsersParams) ([]domain.User, error)
	PatchUser(ctx context.Context, params *presenter.PatchUserParams) (*domain.User, error)
	DeleteUser(ctx context.Context, params *presenter.DeleteUserParams) error

	GetUserGroups(ctx context.Context, params *presenter.GetUserGroupsParams) ([]domain.Group, error)
	GetUserTopics(ctx context.Context, params *presenter.GetUserTopicsParams) ([]domain.Topic, error)
}

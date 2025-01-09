package port

import (
	"context"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
)

type GroupService interface {
}

type GroupStore interface {
	CreateGroup(ctx context.Context, params *presenter.CreateGroupParams) (*domain.Group, error)
	FindGroup(ctx context.Context, params *presenter.FindGroupParams) (*domain.Group, error)
	ListGroups(ctx context.Context, params *presenter.ListGroupsParams) ([]domain.Group, error)
	PatchGroup(ctx context.Context, params *presenter.PatchGroupParams) (*domain.Group, error)
	DeleteGroup(ctx context.Context, params *presenter.DeleteGroupParams) error

	GetGroupUsers(ctx context.Context, params *presenter.GetGroupUsersParams) ([]domain.User, error)
	GetGoals(ctx context.Context, params *presenter.GetGroupGoalsParams) ([]domain.Goal, error)
}

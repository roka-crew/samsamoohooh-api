package service

import (
	"context"
	"samsamoohooh-api/internal/application/core/user/store"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
)

type UserService struct {
	userStore store.UserStore
}

func NewUserService(
	userStore store.UserStore,
) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (s *UserService) FindUser(ctx context.Context, req *presenter.FindUserRequest) (*domain.User, error) {
	foundUser, err := s.userStore.FindUser(ctx, req.ToParams())
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (s *UserService) PatchUser(ctx context.Context, req *presenter.PatchUserRequest) (*domain.User, error) {
	patchedUser, err := s.userStore.PatchUser(ctx, req.ToParams())
	if err != nil {
		return nil, err
	}

	return patchedUser, nil
}

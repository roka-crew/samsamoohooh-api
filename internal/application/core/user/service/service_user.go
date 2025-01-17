package service

import (
	"context"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/port"
	"samsamoohooh-api/internal/application/presenter"
)

type UserService struct {
	userStore port.UserStore
}

func NewUserService(
	userStore port.UserStore,
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

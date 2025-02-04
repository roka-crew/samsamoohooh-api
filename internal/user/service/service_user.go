package service

import (
	"context"
	"samsamoohooh-api/domain"
)

type userService struct {
	userStore domain.UserStore
}

func NewUserService(userStore domain.UserStore) domain.UserService {
	return &userService{
		userStore: userStore,
	}
}

func (s userService) FindUserByMe(ctx context.Context, request domain.FindUserByMeRequest) (*domain.User, error) {
	foundUser, err := s.userStore.FindUser(ctx, domain.FindUserParams{
		UserID: request.RequestUserID,
	})
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (s userService) PatchByMeUser(ctx context.Context, request domain.PatchUserByMeRequest) error {
	err := s.userStore.PatchUser(ctx, domain.PatchUserParams{
		UserID:     request.RequestUserID,
		Nickname:   request.Nickname,
		Resolution: request.Resolution,
	})
	if err != nil {
		return err
	}

	return nil
}

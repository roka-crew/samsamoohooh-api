package store

import (
	"context"
	"samsamoohooh-api/domain"
	"samsamoohooh-api/pkg/errors"
	"samsamoohooh-api/pkg/mysql"
)

type userStore struct {
	db *mysql.Mysql
}

func NewUserStore(db *mysql.Mysql) domain.UserStore {
	return &userStore{db: db}
}

func (s userStore) CreateUser(ctx context.Context, params domain.CreateUserParams) (*domain.User, error) {
	db := s.db.WithContext(ctx)

	createUser := &domain.User{
		Nickname:   params.Nickname,
		Resolution: params.Resolution,
		Provider:   params.Provider,
	}
	if err := db.Create(createUser).Error; err != nil {
		return nil, errors.New(err)
	}

	return createUser, nil
}

func (s userStore) FindUser(ctx context.Context, params domain.FindUserParams) (*domain.User, error) {
	db := s.db.WithContext(ctx)

	findUser := &domain.User{}
	if err := db.First(findUser, params.UserID).Error; err != nil {
		return nil, errors.New(err)
	}

	return findUser, nil
}

func (s userStore) PatchUser(ctx context.Context, params domain.PatchUserParams) error {
	db := s.db.WithContext(ctx)

	patchUser := &domain.User{}

	if params.Nickname != nil {
		patchUser.Nickname = *params.Nickname
	}

	if params.Resolution != nil {
		patchUser.Resolution = params.Resolution
	}

	if params.Provider != nil {
		patchUser.Provider = *params.Provider
	}

	err := db.Model(&domain.User{}).
		Where("id = ?", params.UserID).
		Updates(patchUser).Error
	if err != nil {
		return errors.New(err)
	}

	return nil
}

func (s userStore) DeleteUser(ctx context.Context, params domain.DeleteUserParams) error {
	db := s.db.WithContext(ctx)

	if err := db.Delete(&domain.User{}, params.UserID).Error; err != nil {
		return errors.New(err)
	}

	return nil
}

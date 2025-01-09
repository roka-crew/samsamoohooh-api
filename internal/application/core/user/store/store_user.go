package store

import (
	"context"
	"errors"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/persistence/mysql"
	"samsamoohooh-api/internal/infra/validator"
	"samsamoohooh-api/pkg/httperr"

	"gorm.io/gorm"
)

type UserStore struct {
	db        *mysql.Mysql
	validator *validator.Validator
}

func NewUserStore(
	db *mysql.Mysql,
	validator *validator.Validator,
) *UserStore {
	return &UserStore{
		db:        db,
		validator: validator,
	}
}

func (s *UserStore) CreateUser(ctx context.Context, params *presenter.CreateUserParams) (*domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var createUser = &domain.User{
		Nickname:   params.Nickname,
		Resolution: params.Resolution,
		Provider:   params.Provider,
	}
	err = db.
		Create(createUser).
		Error
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed create user")
	}

	return createUser, nil
}

func (s *UserStore) FindUser(ctx context.Context, params *presenter.FoundUserParams) (*domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var foundUser = &domain.User{}
	err = db.
		First(foundUser, params.UserID).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperr.New(err).
			SetType(httperr.DBNotFound).
			SetDetail("failed retrieve %d user", params.UserID)
	}

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed retrieve %d user", params.UserID)
	}

	return foundUser, nil
}

func (s *UserStore) ListUsers(ctx context.Context, params *presenter.ListUsersParams) ([]domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	var listUsers []domain.User
	err = db.
		Limit(params.Limit).
		Find(&listUsers).
		Error
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed list user")
	}

	return listUsers, nil
}

func (s *UserStore) PatchUser(ctx context.Context, params *presenter.PatchUserParams) (*domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var patchUser = &domain.User{}
	if params.Nickname != nil {
		patchUser.Nickname = *params.Nickname
	}
	if params.Resolution != nil {
		patchUser.Resolution = params.Resolution
	}
	if params.Provider != nil {
		patchUser.Provider = *params.Provider
	}

	res := db.
		Model(&domain.User{ID: params.UserID}).
		Updates(patchUser)

	if res.RowsAffected == 0 {
		return nil, httperr.New().
			SetType(httperr.DBUpdateNotApplied).
			SetDetail("update operation did not affect any records")
	}

	if res.Error != nil {
		return nil, httperr.New(res.Error).
			SetType(httperr.DBFailed).
			SetDetail("failed update user")
	}

	return patchUser, nil
}

func (s *UserStore) DeleteUser(ctx context.Context, params *presenter.DeleteUserParams) error {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return err
	}

	db := s.db.WithContext(ctx)

	var deleteUser = &domain.User{ID: params.UserID}
	res := db.
		Delete(&deleteUser)

	if res.RowsAffected == 0 {
		return httperr.New().
			SetType(httperr.DBDeleteNotApplied).
			SetDetail("delete operation did not affect any records")
	}

	if res.Error != nil {
		return httperr.New(res.Error).
			SetType(httperr.DBFailed).
			SetDetail("failed delete user")
	}

	return nil
}

func (s *UserStore) GetUserGroups(ctx context.Context, params *presenter.GetUserGroupsParams) ([]domain.Group, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	var foundGroups []domain.Group
	err = db.
		Model(&domain.User{ID: params.UserID}).
		Limit(params.Limit).
		Association("Groups").
		Find(&foundGroups)

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed get groups")
	}

	return foundGroups, nil
}

func (s *UserStore) GetUserTopics(ctx context.Context, params *presenter.GetUserTopicsParams) ([]domain.Topic, error) {
	err := s.validator.Validate(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	var foundTopics []domain.Topic
	err = db.
		Model(&domain.User{ID: params.UserID}).
		Limit(params.Limit).
		Association("Topics").
		Find(&foundTopics)

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed get topics")
	}

	return foundTopics, nil
}

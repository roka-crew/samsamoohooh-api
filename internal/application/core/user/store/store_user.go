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

	var createUser = &domain.User{
		Nickname:   params.Nickname,
		Resolution: params.Resolution,
		Provider:   params.Provider,
	}
	err = s.db.
		WithContext(ctx).
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

	var foundUser = &domain.User{}
	err = s.db.
		WithContext(ctx).
		First(foundUser, params.ID).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperr.New(err).
			SetType(httperr.DBNotFound).
			SetDetail("failed retrieve %d user", params.ID)
	}

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed retrieve %d user", params.ID)
	}

	return foundUser, nil
}

func (s *UserStore) ListUsers(ctx context.Context, params *presenter.ListUsersParams) (*domain.Paginator[domain.User], error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	var listUsers []domain.User
	err = s.db.WithContext(ctx).
		Where("id >= ?", params.Cursor).
		Limit(params.Limit + 1).
		Find(&listUsers).
		Error
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed list user")
	}

	hasNext := len(listUsers) > params.Limit
	if hasNext {
		listUsers = listUsers[:params.Limit]
	}

	var nextCursor int
	if hasNext {
		nextCursor = params.Cursor + params.Limit
	}

	paginator := &domain.Paginator[domain.User]{
		Items:      listUsers,
		HasNext:    hasNext,
		NextCursor: nextCursor,
	}

	return paginator, nil
}

func (s *UserStore) PatchUser(ctx context.Context, params *presenter.PatchUserParams) (*domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

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

	res := s.db.WithContext(ctx).
		Model(&domain.User{ID: params.ID}).
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

	var deleteUser = &domain.User{ID: params.ID}
	res := s.db.WithContext(ctx).
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

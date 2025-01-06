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

func (s *UserStore) FoundUser(ctx context.Context, params *presenter.FoundUserParams) (*domain.User, error) {
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
			SetType(httperr.DBRecordNotFound).
			SetDetail("failed retrieve %d user", params.ID)
	}

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed retrieve %d user", params.ID)
	}

	return foundUser, nil
}

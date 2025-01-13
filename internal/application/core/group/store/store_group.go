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

type GroupStore struct {
	db        *mysql.Mysql
	validator *validator.Validator
}

func NewGroupStore(
	db *mysql.Mysql,
	validator *validator.Validator,
) *GroupStore {
	return &GroupStore{
		db:        db,
		validator: validator,
	}
}

func (s *GroupStore) CreateGroup(ctx context.Context, params *presenter.CreateGroupParams) (*domain.Group, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var createGroup = &domain.Group{
		BookTitle:        params.BookTitle,
		BookAuthor:       params.BookAuthor,
		BookPageMax:      params.BookPageMax,
		BookPublisher:    params.BookPublisher,
		BookIntroduction: params.BookIntroduction,
	}

	err = db.
		Create(createGroup).
		Error
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed create group")
	}

	return createGroup, nil
}

func (s *GroupStore) FindGroup(ctx context.Context, params *presenter.FindGroupParams) (*domain.Group, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var findGroup = &domain.Group{}

	err = db.
		First(findGroup, params.GroupID).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, httperr.New(err).
			SetType(httperr.DBNotFound).
			SetDetail("failed retrieve %d group", params.GroupID)
	}

	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed retrieve %d group", params.GroupID)
	}

	return findGroup, nil
}

func (s *GroupStore) ListGroups(ctx context.Context, params *presenter.ListGroupsParams) ([]domain.Group, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	var listGroups []domain.Group
	err = db.
		Limit(params.Limit).
		Find(&listGroups).
		Error
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed list groups")
	}

	return listGroups, nil
}

func (s *GroupStore) PatchGroup(ctx context.Context, params *presenter.PatchGroupParams) (*domain.Group, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	var patchGroup = &domain.Group{ID: params.GroupID}
	if params.BookTitle != nil {
		patchGroup.BookTitle = *params.BookTitle
	}

	if params.BookAuthor != nil {
		patchGroup.BookAuthor = *params.BookAuthor
	}

	if params.BookPageMax != nil {
		patchGroup.BookPageMax = *params.BookPageMax
	}

	if params.BookPublisher != nil {
		patchGroup.BookPublisher = params.BookPublisher
	}

	if params.BookIntroduction != nil {
		patchGroup.BookIntroduction = params.BookIntroduction
	}

	res := db.
		Updates(patchGroup)

	if res.RowsAffected == 0 {
		return nil, httperr.New(nil).
			SetType(httperr.DBUpdateNotApplied).
			SetDetail("update operation did not affect any records")
	}

	if res.Error != nil {
		return nil, httperr.New(res.Error).
			SetType(httperr.DBFailed).
			SetDetail("failed patch %d group", params.GroupID)

	}

	return patchGroup, nil
}

func (s *GroupStore) DeleteGroup(ctx context.Context, params *presenter.DeleteGroupParams) error {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return err
	}

	db := s.db.WithContext(ctx)

	res := db.
		Delete(&domain.Group{}, params.GroupID)

	if res.RowsAffected == 0 {
		return httperr.New(nil).
			SetType(httperr.DBDeleteNotApplied).
			SetDetail("delete operation did not affect any records")
	}

	if res.Error != nil {
		return httperr.New(res.Error).
			SetType(httperr.DBFailed).
			SetDetail("failed delete %d group", params.GroupID)
	}

	return nil
}

func (s *GroupStore) GetGroupUsers(ctx context.Context, params *presenter.GetGroupUsersParams) ([]domain.User, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db = db.Offset(*params.Offset)
	}

	var getUsers []domain.User
	err = db.
		Model(&domain.Group{ID: params.GroupID}).
		Limit(params.Limit).
		Association("Users").
		Find(getUsers)
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed get users")
	}

	return getUsers, nil
}

func (s *GroupStore) AddGroupUser(ctx context.Context, params *presenter.AddGroupUserParams) error {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return err
	}

	db := s.db.WithContext(ctx)

	var findGroup = &domain.Group{}
	err = db.
		First(findGroup, params.GroupID).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return httperr.New(err).
			SetType(httperr.DBNotFound).
			SetDetail("not found group")
	}
	if err != nil {
		return httperr.New(err).
			SetType(httperr.DBFailed)
	}

	var findUser = &domain.User{}
	err = db.
		First(findUser, params.UserID).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return httperr.New(err).
			SetType(httperr.DBNotFound).
			SetDetail("not found user, check user id")
	}

	if err != nil {
		return httperr.New(err).
			SetType(httperr.DBFailed)
	}

	err = db.
		Model(findGroup).
		Association("Users").
		Append(&findUser)
	if err != nil {
		return httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("not append relation")
	}

	return nil
}

func (s *GroupStore) GetGoals(ctx context.Context, params *presenter.GetGroupGoalsParams) ([]domain.Goal, error) {
	err := s.validator.ValidateParams(params)
	if err != nil {
		return nil, err
	}

	db := s.db.WithContext(ctx)

	if params.Offset != nil {
		db.Offset(*params.Offset)
	}

	var getGoals []domain.Goal
	err = db.
		Model(&domain.Group{ID: params.GroupID}).
		Limit(params.Limit).
		Association("Goals").
		Find(&getGoals)
	if err != nil {
		return nil, httperr.New(err).
			SetType(httperr.DBFailed).
			SetDetail("failed get goals")
	}

	return getGoals, nil
}

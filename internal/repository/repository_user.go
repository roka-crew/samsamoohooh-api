package repository

import (
	"context"
	"fmt"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *sqlite.Sqlite
}

func NewUserRepository(
	db *sqlite.Sqlite,
) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, param domain.CreateUserParam) (*domain.User, error) {
	user := domain.User{
		Name:     param.Name,
		DeviceID: param.DeviceID,
	}

	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) ListUser(ctx context.Context, param domain.ListUserParam) (*domain.User, error) {
	query := u.db.WithContext(ctx)
	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.WithBooks {
		query = query.Preload("UserBooks")
	}
	if param.WithThoughts {
		query = query.Preload("Thoughts")
	}

	var user domain.User
	if err := query.Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) ListUsers(ctx context.Context, param domain.ListUsersParam) (domain.Users, error) {
	query := u.db.WithContext(ctx).Model(&domain.User{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.Name) > 0 {
		query = query.Where("name IN ?", param.Name)
	}
	if len(param.DeviceIDs) > 0 {
		query = query.Where("biography IN ?", param.DeviceIDs)
	}

	if param.WithUserBooks {
		query = query.Preload("UserBooks")
	}
	if param.WithThoughts {
		query = query.Preload("Thoughts")
	}

	for _, o := range param.OrderBy {
		query = query.Order(o)
	}

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}
	if param.Offset > 0 {
		query = query.Offset(param.Offset)
	}

	var users domain.Users
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) PatchUser(ctx context.Context, param domain.PatchUserParam) error {
	updates := make(map[string]any)

	if param.Name != nil {
		updates["name"] = *param.Name
	}
	if param.DeviceID != nil {
		updates["device_id"] = *param.DeviceID
	}

	if len(updates) == 0 {
		return nil
	}

	result := u.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (u *UserRepository) PatchUsers(ctx context.Context, param domain.PatchUsersParam) error {
	if len(param) == 0 {
		return nil
	}

	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			updates := make(map[string]any)

			if p.Name != nil {
				updates["name"] = *p.Name
			}
			if p.DeviceID != nil {
				updates["device_id"] = *p.DeviceID
			}

			if len(updates) == 0 {
				continue
			}

			result := tx.Model(&domain.User{}).Where("id = ?", p.ID).Updates(updates)
			if result.Error != nil {
				return fmt.Errorf("patch user(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("patch user(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

func (u *UserRepository) DeleteUser(ctx context.Context, param domain.DeleteUserParam) error {
	query := u.db.WithContext(ctx)
	if param.IsHardDelete {
		query = query.Unscoped()
	}

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}

	if param.DeviceID != nil {
		query = query.Where("device_id = ?", param.DeviceID)
	}
	result := query.Delete(&domain.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (u *UserRepository) DeleteUsers(ctx context.Context, param domain.DeleteUsersParam) error {
	if len(param) == 0 {
		return nil
	}

	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			query := tx
			if p.IsHardDelete {
				query = query.Unscoped()
			}

			result := query.Delete(&domain.User{}, "id = ?", p.ID)
			if result.Error != nil {
				return fmt.Errorf("delete user(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("delete user(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

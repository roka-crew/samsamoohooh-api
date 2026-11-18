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

func (r *UserRepository) CreateUser(ctx context.Context, param domain.CreateUserParam) (*domain.User, error) {
	user := domain.User{
		Name:     param.Name,
		DeviceID: param.DeviceID,
	}

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) ListUser(ctx context.Context, param domain.ListUserParam) (*domain.User, error) {
	query := r.db.WithContext(ctx)
	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.DeviceID != nil {
		query = query.Where("device_id = ?", *param.DeviceID)
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

func (r *UserRepository) ListUsers(ctx context.Context, param domain.ListUsersParam) (domain.Users, error) {
	query := r.db.WithContext(ctx).Model(&domain.User{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.Names) > 0 {
		query = query.Where("name IN ?", param.Names)
	}
	if len(param.DeviceIDs) > 0 {

		query = query.Where("device_id IN ?", param.DeviceIDs)
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

func (r *UserRepository) PatchUser(ctx context.Context, param domain.PatchUserParam) error {
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

	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserRepository) PatchUsers(ctx context.Context, param domain.PatchUsersParam) error {
	if len(param) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (r *UserRepository) DeleteUser(ctx context.Context, param domain.DeleteUserParam) error {
	query := r.db.WithContext(ctx)
	if param.IsHardDelete {
		query = query.Unscoped()
	}

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}

	if param.DeviceID != nil {
		query = query.Where("device_id = ?", *param.DeviceID)
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

func (r *UserRepository) DeleteUsers(ctx context.Context, param domain.DeleteUsersParam) error {
	if len(param) == 0 {
		return nil
	}

	var softDeleteIDs []uint
	var hardDeleteIDs []uint

	for _, p := range param {
		if p.ID == 0 {
			continue
		}
		if p.IsHardDelete {
			hardDeleteIDs = append(hardDeleteIDs, p.ID)
		} else {
			softDeleteIDs = append(softDeleteIDs, p.ID)
		}
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(softDeleteIDs) > 0 {
			if err := tx.Where("id IN ?", softDeleteIDs).Delete(&domain.User{}).Error; err != nil {
				return fmt.Errorf("batch soft delete failed: %w", err)
			}
		}

		if len(hardDeleteIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", hardDeleteIDs).Delete(&domain.User{}).Error; err != nil {
				return fmt.Errorf("batch hard delete failed: %w", err)
			}
		}

		return nil
	})
}

package repository

import (
	"context"
	"fmt"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"gorm.io/gorm"
)

type UserBookRepository struct {
	db *sqlite.Sqlite
}

func NewUserBookRepository(
	db *sqlite.Sqlite,
) domain.UserBookRepository {
	return &UserBookRepository{
		db: db,
	}
}

func (r *UserBookRepository) CreateUserBook(ctx context.Context, param domain.CreateUserBookParam) (domain.UserBook, error) {
	userBook := domain.UserBook{
		UserID: param.UserID,
		BookID: param.BookID,
	}

	if err := r.db.WithContext(ctx).Create(&userBook).Error; err != nil {
		return domain.UserBook{}, err
	}

	return userBook, nil
}

func (r *UserBookRepository) CreateUserBooks(ctx context.Context, param domain.CreateUserBooksParam) (domain.UserBooks, error) {
	if len(param) == 0 {
		return nil, nil
	}

	userBooks := make(domain.UserBooks, 0, len(param))
	for _, p := range param {
		userBooks = append(userBooks, domain.UserBook{
			UserID: p.UserID,
			BookID: p.BookID,
		})
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userBooks).Error; err != nil {
			return fmt.Errorf("batch create user_books failed: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return userBooks, nil
}

func (r *UserBookRepository) ListUserBook(ctx context.Context, param domain.ListUserBookParam) (domain.UserBook, error) {
	query := r.db.WithContext(ctx).Model(&domain.UserBook{})

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.UserID != nil {
		query = query.Where("user_id = ?", *param.UserID)
	}
	if param.BookID != nil {
		query = query.Where("book_id = ?", *param.BookID)
	}

	if param.WithUser {
		query = query.Preload("User")
	}
	if param.WithBook {
		query = query.Preload("Book")
	}

	var userBook domain.UserBook
	if err := query.First(&userBook).Error; err != nil {
		return domain.UserBook{}, err
	}

	return userBook, nil
}

func (r *UserBookRepository) ListUserBooks(ctx context.Context, param domain.ListUserBooksParam) (domain.UserBooks, error) {
	query := r.db.WithContext(ctx).Model(&domain.UserBook{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.UserIDs) > 0 {
		query = query.Where("user_id IN ?", param.UserIDs)
	}
	if len(param.BookIDs) > 0 {
		query = query.Where("book_id IN ?", param.BookIDs)
	}

	if param.WithUser {
		query = query.Preload("User")
	}
	if param.WithBook {
		query = query.Preload("Book")
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

	var userBooks domain.UserBooks
	if err := query.Find(&userBooks).Error; err != nil {
		return nil, err
	}

	return userBooks, nil
}

func (r *UserBookRepository) PatchUserBook(ctx context.Context, param domain.PatchUserBookParam) error {
	updates := make(map[string]any)

	if param.UserID != nil {
		updates["user_id"] = *param.UserID
	}
	if param.BookID != nil {
		updates["book_id"] = *param.BookID
	}

	if len(updates) == 0 {
		return nil
	}

	result := r.db.WithContext(ctx).Model(&domain.UserBook{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserBookRepository) PatchUserBooks(ctx context.Context, param domain.PatchUserBooksParam) error {
	if len(param) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			updates := make(map[string]any)

			if p.UserID != nil {
				updates["user_id"] = *p.UserID
			}
			if p.BookID != nil {
				updates["book_id"] = *p.BookID
			}

			if len(updates) == 0 {
				continue
			}

			result := tx.Model(&domain.UserBook{}).Where("id = ?", p.ID).Updates(updates)
			if result.Error != nil {
				return fmt.Errorf("patch user_book(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("patch user_book(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

func (r *UserBookRepository) DeleteUserBook(ctx context.Context, param domain.DeleteUserBookParam) error {
	query := r.db.WithContext(ctx)
	if param.IsHardDelete {
		query = query.Unscoped()
	}

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.UserID != nil {
		query = query.Where("user_id = ?", *param.UserID)
	}
	if param.BookID != nil {
		query = query.Where("book_id = ?", *param.BookID)
	}

	result := query.Delete(&domain.UserBook{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *UserBookRepository) DeleteUserBooks(ctx context.Context, param domain.DeleteUserBooksParam) error {
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
			if err := tx.Where("id IN ?", softDeleteIDs).Delete(&domain.UserBook{}).Error; err != nil {
				return fmt.Errorf("batch soft delete failed: %w", err)
			}
		}

		if len(hardDeleteIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", hardDeleteIDs).Delete(&domain.UserBook{}).Error; err != nil {
				return fmt.Errorf("batch hard delete failed: %w", err)
			}
		}

		return nil
	})
}

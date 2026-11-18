package repository

import (
	"context"
	"fmt"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *sqlite.Sqlite
}

func NewBookRepository(
	db *sqlite.Sqlite,
) domain.BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, param domain.CreateBookParam) (domain.Book, error) {
	book := domain.Book{
		Title:     param.Title,
		PageCount: param.PageCount,
	}

	if err := r.db.WithContext(ctx).Create(&book).Error; err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) ListBook(ctx context.Context, param domain.ListBookParam) (domain.Book, error) {
	query := r.db.WithContext(ctx).Model(&domain.Book{})

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Title != nil {
		query = query.Where("title = ?", *param.Title)
	}
	if param.PageCount > 0 {
		query = query.Where("page_count = ?", param.PageCount)
	}

	if param.WithUserBooks {
		query = query.Preload("UserBooks")
	}
	if param.WithMilestones {
		query = query.Preload("Milestones")
	}

	var book domain.Book
	if err := query.First(&book).Error; err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (r *BookRepository) ListBooks(ctx context.Context, param domain.ListBooksParam) (domain.Books, error) {
	query := r.db.WithContext(ctx).Model(&domain.Book{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.Titles) > 0 {
		query = query.Where("title IN ?", param.Titles)
	}
	if len(param.PageCounts) > 0 {
		query = query.Where("page_count IN ?", param.PageCounts)
	}

	if param.WithUserBooks {
		query = query.Preload("UserBooks")
	}
	if param.WithMilestones {
		query = query.Preload("Milestones")
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

	var books domain.Books
	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) PatchBook(ctx context.Context, param domain.PatchBookParam) error {
	updates := make(map[string]any)

	if param.Title != nil {
		updates["title"] = *param.Title
	}
	if param.PageCount != nil {
		updates["page_count"] = *param.PageCount
	}

	if len(updates) == 0 {
		return nil
	}

	result := r.db.WithContext(ctx).Model(&domain.Book{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *BookRepository) PatchBooks(ctx context.Context, param domain.PatchBooksParam) error {
	if len(param) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			updates := make(map[string]any)

			if p.Title != nil {
				updates["title"] = *p.Title
			}
			if p.PageCount != nil {
				updates["page_count"] = *p.PageCount
			}

			if len(updates) == 0 {
				continue
			}

			result := tx.Model(&domain.Book{}).Where("id = ?", p.ID).Updates(updates)
			if result.Error != nil {
				return fmt.Errorf("patch book(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("patch book(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

func (r *BookRepository) DeleteBook(ctx context.Context, param domain.DeleteBookParam) error {
	query := r.db.WithContext(ctx)
	if param.IsHardDelete {
		query = query.Unscoped()
	}

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Title != nil {
		query = query.Where("title = ?", *param.Title)
	}

	result := query.Delete(&domain.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *BookRepository) DeleteBooks(ctx context.Context, param domain.DeleteBooksParam) error {
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
			if err := tx.Where("id IN ?", softDeleteIDs).Delete(&domain.Book{}).Error; err != nil {
				return fmt.Errorf("batch soft delete failed: %w", err)
			}
		}

		if len(hardDeleteIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", hardDeleteIDs).Delete(&domain.Book{}).Error; err != nil {
				return fmt.Errorf("batch hard delete failed: %w", err)
			}
		}

		return nil
	})
}

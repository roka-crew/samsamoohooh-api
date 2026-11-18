package repository

import (
	"context"
	"fmt"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"gorm.io/gorm"
)

type MilestoneRepository struct {
	db *sqlite.Sqlite
}

func NewMilestoneRepository(
	db *sqlite.Sqlite,
) domain.MilestoneRepository {
	return &MilestoneRepository{
		db: db,
	}
}

func (r *MilestoneRepository) CreateMilestone(ctx context.Context, param domain.CreateMilestoneParam) (domain.Milestone, error) {
	milestone := domain.Milestone{
		BookID:  param.BookID,
		Content: param.Content,
		DueDate: param.DueDate,
	}

	if err := r.db.WithContext(ctx).Create(&milestone).Error; err != nil {
		return domain.Milestone{}, err
	}

	return milestone, nil
}

func (r *MilestoneRepository) ListMilestone(ctx context.Context, param domain.ListMilestoneParam) (domain.Milestone, error) {
	query := r.db.WithContext(ctx).Model(&domain.Milestone{})

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.BookID != nil {
		query = query.Where("book_id = ?", *param.BookID)
	}

	if param.WithBook {
		query = query.Preload("Book")
	}
	if param.WithThoughts {
		query = query.Preload("Thoughts")
	}

	var milestone domain.Milestone
	if err := query.First(&milestone).Error; err != nil {
		return domain.Milestone{}, err
	}

	return milestone, nil
}

func (r *MilestoneRepository) ListMilestones(ctx context.Context, param domain.ListMilestonesParam) (domain.Milestones, error) {
	query := r.db.WithContext(ctx).Model(&domain.Milestone{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.BookIDs) > 0 {
		query = query.Where("book_id IN ?", param.BookIDs)
	}
	if len(param.Contents) > 0 {
		query = query.Where("content IN ?", param.Contents)
	}
	if param.DueDateFrom != nil {
		query = query.Where("due_date >= ?", *param.DueDateFrom)
	}
	if param.DueDateTo != nil {
		query = query.Where("due_date <= ?", *param.DueDateTo)
	}

	if param.WithBook {
		query = query.Preload("Book")
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

	var milestones domain.Milestones
	if err := query.Find(&milestones).Error; err != nil {
		return nil, err
	}

	return milestones, nil
}

func (r *MilestoneRepository) PatchMilestone(ctx context.Context, param domain.PatchMilestoneParam) error {
	updates := make(map[string]any)

	if param.BookID != nil {
		updates["book_id"] = *param.BookID
	}
	if param.Content != nil {
		updates["content"] = *param.Content
	}
	if param.DueDate != nil {
		updates["due_date"] = *param.DueDate
	}

	if len(updates) == 0 {
		return nil
	}

	result := r.db.WithContext(ctx).Model(&domain.Milestone{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *MilestoneRepository) PatchMilestones(ctx context.Context, param domain.PatchMilestonesParam) error {
	if len(param) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			updates := make(map[string]any)

			if p.BookID != nil {
				updates["book_id"] = *p.BookID
			}
			if p.Content != nil {
				updates["content"] = *p.Content
			}
			if p.DueDate != nil {
				updates["due_date"] = *p.DueDate
			}

			if len(updates) == 0 {
				continue
			}

			result := tx.Model(&domain.Milestone{}).Where("id = ?", p.ID).Updates(updates)
			if result.Error != nil {
				return fmt.Errorf("patch milestone(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("patch milestone(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

func (r *MilestoneRepository) DeleteMilestone(ctx context.Context, param domain.DeleteMilestoneParam) error {
	query := r.db.WithContext(ctx)
	if param.IsHardDelete {
		query = query.Unscoped()
	}

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.BookID != nil {
		query = query.Where("book_id = ?", *param.BookID)
	}

	result := query.Delete(&domain.Milestone{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *MilestoneRepository) DeleteMilestones(ctx context.Context, param domain.DeleteMilestonesParam) error {
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
			if err := tx.Where("id IN ?", softDeleteIDs).Delete(&domain.Milestone{}).Error; err != nil {
				return fmt.Errorf("batch soft delete failed: %w", err)
			}
		}

		if len(hardDeleteIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", hardDeleteIDs).Delete(&domain.Milestone{}).Error; err != nil {
				return fmt.Errorf("batch hard delete failed: %w", err)
			}
		}

		return nil
	})
}

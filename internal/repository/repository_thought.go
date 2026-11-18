package repository

import (
	"context"
	"fmt"

	"github.com/roka-crew/samsamoohooh-api/internal/database/sqlite"
	"github.com/roka-crew/samsamoohooh-api/internal/domain"
	"gorm.io/gorm"
)

type ThoughtRepository struct {
	db *sqlite.Sqlite
}

func NewThoughtRepository(
	db *sqlite.Sqlite,
) domain.ThoughtRepository {
	return &ThoughtRepository{
		db: db,
	}
}

func (r *ThoughtRepository) CreateThought(ctx context.Context, param domain.CreateThoughtParam) (*domain.Thought, error) {
	thought := domain.Thought{
		UserID:      param.UserID,
		MilestoneID: param.MilestoneID,
		Content:     param.Content,
		IsAnonymous: param.IsAnonymous,
	}

	if err := r.db.WithContext(ctx).Create(&thought).Error; err != nil {
		return nil, err
	}

	return &thought, nil
}

func (r *ThoughtRepository) ListThought(ctx context.Context, param domain.ListThoughtParam) (*domain.Thought, error) {
	query := r.db.WithContext(ctx)

	if param.ID > 0 {
		query = query.Where("id = ?", param.ID)
	}

	if param.UserID > 0 {
		query = query.Where("user_id = ?", param.UserID)
	}

	if param.MilestoneID > 0 {
		query = query.Where("milestone_id = ?", param.MilestoneID)
	}

	if param.Content != nil {
		query = query.Where("content = ?", *param.Content)
	}

	if param.IsAnonymous != nil {
		query = query.Where("is_anonymous = ?", *param.IsAnonymous)
	}

	if param.WithUser {
		query = query.Preload("User")
	}

	if param.WithMileston {
		query = query.Preload("Milestone")
	}

	var thought domain.Thought
	if err := query.Find(&thought).Error; err != nil {
		return nil, err
	}

	return &thought, nil
}

func (r *ThoughtRepository) ListThoughts(ctx context.Context, param domain.ListThoughtsParam) (domain.Thoughts, error) {
	query := r.db.WithContext(ctx).Model(&domain.Thought{})

	if len(param.IDs) > 0 {
		query = query.Where("id IN ?", param.IDs)
	}
	if len(param.UserIDs) > 0 {
		query = query.Where("user_id IN ?", param.UserIDs)
	}
	if len(param.MilestoneIDs) > 0 {
		query = query.Where("milestone_id IN ?", param.MilestoneIDs)
	}
	if len(param.Contents) > 0 {
		query = query.Where("content IN ?", param.Contents)
	}
	if len(param.IsAnonymouses) > 0 {
		query = query.Where("is_anonymous IN ?", param.IsAnonymouses)
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

	var thoughts domain.Thoughts
	if err := query.Find(&thoughts).Error; err != nil {
		return nil, err
	}

	return thoughts, nil
}

func (r *ThoughtRepository) PatchThought(ctx context.Context, param domain.PatchThoughtParam) error {
	updates := make(map[string]any)

	if param.UserID != nil {
		updates["user_id"] = *param.UserID
	}
	if param.MilestoneID != nil {
		updates["milestone_id"] = *param.MilestoneID
	}
	if param.Content != nil {
		updates["content"] = *param.Content
	}
	if param.IsAnonymous != nil {
		updates["is_anonymous"] = *param.IsAnonymous
	}

	if len(updates) == 0 {
		return nil
	}

	result := r.db.WithContext(ctx).Model(&domain.Thought{}).Where("id = ?", param.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ThoughtRepository) PatchThoughts(ctx context.Context, param domain.PatchThoughtsParam) error {
	if len(param) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, p := range param {
			updates := make(map[string]any)

			if p.UserID != nil {
				updates["user_id"] = *p.UserID
			}
			if p.MilestoneID != nil {
				updates["milestone_id"] = *p.MilestoneID
			}
			if p.Content != nil {
				updates["content"] = *p.Content
			}
			if p.IsAnonymous != nil {
				updates["is_anonymous"] = *p.IsAnonymous
			}

			if len(updates) == 0 {
				continue
			}

			result := tx.Model(&domain.Thought{}).Where("id = ?", p.ID).Updates(updates)
			if result.Error != nil {
				return fmt.Errorf("patch thought(id=%d) failed: %w", p.ID, result.Error)
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("patch thought(id=%d) failed: %w", p.ID, gorm.ErrRecordNotFound)
			}
		}
		return nil
	})
}

func (r *ThoughtRepository) DeleteThought(ctx context.Context, param domain.DeleteThoughtParam) error {
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
	if param.MilestoneID != nil {
		query = query.Where("milestone_id = ?", *param.MilestoneID)
	}

	result := query.Delete(&domain.Thought{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ThoughtRepository) DeleteThoughts(ctx context.Context, param domain.DeleteThoughtsParam) error {
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
			if err := tx.Where("id IN ?", softDeleteIDs).Delete(&domain.Thought{}).Error; err != nil {
				return fmt.Errorf("batch soft delete failed: %w", err)
			}
		}

		if len(hardDeleteIDs) > 0 {
			if err := tx.Unscoped().Where("id IN ?", hardDeleteIDs).Delete(&domain.Thought{}).Error; err != nil {
				return fmt.Errorf("batch hard delete failed: %w", err)
			}
		}

		return nil
	})
}

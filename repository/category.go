package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	tx := r.db.WithContext(ctx)
	resValue := []entity.Category{}
	if res := tx.Model(&entity.Category{}).Where("user_id = ?", id).Scan(&resValue); res.Error != nil {
		return nil, res.Error
	}

	// if len(resValue) == 0 {
	// 	return []entity.Category{}, nil
	// }
	return resValue, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	resValue := entity.Category{}
	tx := r.db.WithContext(ctx)
	if res := tx.Create(&category).Scan(&resValue); res.Error != nil {
		return 0, res.Error
	}

	return resValue.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	resValue := []entity.Category{}
	tx := r.db.WithContext(ctx)
	if res := tx.Create(&categories).Scan(&resValue); res.Error != nil {
		return res.Error
	}

	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	resValue := entity.Category{}

	// resVal := entity.Category{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.Category{}).Where("id = ?", id).Scan(&resValue); res.Error != nil {
		return entity.Category{}, res.Error
	}

	// if resValue == resVal {
	// 	return entity.Category{}, nil
	// }

	return resValue, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.Category{}).Updates(entity.Category{
		ID:     category.ID,
		Type:   category.Type,
		UserID: category.UserID,
	}); res.Error != nil {
		return res.Error
	}

	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx)
	if res := tx.Delete(&entity.Category{}, id); res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

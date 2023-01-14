package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	tx := r.db.WithContext(ctx)
	resValue := []entity.Task{}
	if res := tx.Model(&entity.Task{}).Where("user_id = ?", id).Scan(&resValue); res.Error != nil {
		return []entity.Task{}, res.Error
	}

	// if len(resValue) == 0 {
	// 	return []entity.Task{}, nil
	// }
	return resValue, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	resValue := entity.Task{}
	tx := r.db.WithContext(ctx)
	if res := tx.Create(&task).Scan(&resValue); res.Error != nil {
		return 0, res.Error
	}

	return resValue.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	resValue := entity.Task{}

	// resVal := entity.Task{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.Task{}).Where("id = ?", id).Scan(&resValue); res.Error != nil {
		return entity.Task{}, res.Error
	}

	// if resValue == resVal {
	// 	return entity.Task{}, nil
	// }

	return resValue, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	resValue := []entity.Task{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.Task{}).Where("category_id = ?", catId).Scan(&resValue); res.Error != nil {
		return []entity.Task{}, res.Error
	}

	// if len(resValue) == 0 {
	// 	return []entity.Task{}, nil
	// }

	return resValue, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      task.UserID,
	}); res.Error != nil {
		return res.Error
	}

	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&entity.Task{}).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

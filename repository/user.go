package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	// r.db.WithContext(ctx).Find(&entity.User{})
	resValue := entity.User{}

	resVal := entity.User{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.User{}).Where("id = ?", id).Scan(&resValue); res.Error != nil {
		return entity.User{}, res.Error
	}

	if resValue == resVal {
		return entity.User{}, nil
	}

	return resValue, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	resValue := entity.User{}

	resVal := entity.User{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.User{}).Where("email = ?", email).Scan(&resValue); res.Error != nil {
		return entity.User{}, res.Error
	}

	if resValue == resVal {
		return entity.User{}, nil
	}

	return resValue, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	resValue := entity.User{}
	tx := r.db.WithContext(ctx)
	if res := tx.Create(&user).Scan(&resValue); res.Error != nil {
		return entity.User{}, res.Error
	}
	return resValue, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	resValue := entity.User{}
	tx := r.db.WithContext(ctx)
	if res := tx.Model(&entity.User{}).Where("id = ?", user.ID).Updates(entity.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}).Scan(&resValue); res.Error != nil {
		return entity.User{}, res.Error
	}
	return resValue, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

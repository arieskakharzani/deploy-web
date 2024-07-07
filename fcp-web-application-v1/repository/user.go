package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var userTaskCategorys []model.UserTaskCategory

	err := r.db.Table("tasks").Select("users.id as id, users.fullname as fullname, users.email as email, tasks.title as task, tasks.deadline as deadline, tasks.priority as priority, tasks.status as status, categories.name as category").Joins("JOIN users ON tasks.user_id = users.id JOIN categories ON tasks.category_id = categories.id").Scan(&userTaskCategorys).Error
	if err != nil {
		return nil, err
	}
	return userTaskCategorys, nil // TODO: replace this
}

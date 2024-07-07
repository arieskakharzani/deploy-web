package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	tasks := &model.Task{}
	data := t.db.Debug().Model(tasks).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"title":       task.Title,
			"deadline":    task.Deadline,
			"priority":    task.Priority,
			"status":      task.Status,
			"category_id": task.CategoryID,
			"user_id":     task.UserID,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	task := model.Task{}
	if result := t.db.Where("id = ?", id).Delete(&task); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var task []model.Task
	rows, err := t.db.Table("tasks").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t.db.ScanRows(rows, &task)
	}
	return task, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var tasks []model.Task
	var categorys []model.Category
	var taskCategorys []model.TaskCategory

	if err := t.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return []model.TaskCategory{}, nil
	}

	if err := t.db.Find(&categorys).Error; err != nil {
		return nil, err
	}
	err := t.db.Table("tasks").Select("tasks.id AS id, tasks.title AS title, categories.name AS category").Joins("JOIN categories ON tasks.category_id = categories.id").Where("tasks.id = ?", id).Scan(&taskCategorys).Error
	if err != nil {
		return nil, err
	}

	return taskCategorys, nil // TODO: replace this
}

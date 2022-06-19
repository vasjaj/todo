package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User        User
	UserID      int
	Title       string
	Description string
	DueDate     *time.Time `gorm:"index"`
	CompletedAt *time.Time `gorm:"index"`
}

func (d *Database) GetTask(id int) (*Task, error) {
	task := &Task{}

	return task, d.DB.Where("id = ?", id).First(task).Error
}

func (d *Database) GetTasks(userID int) ([]Task, error) {
	var tasks []Task

	return tasks, d.DB.Where("user_id = ?", userID).Find(&tasks).Error
}

func (d *Database) GetCompletedTasks(userID int) ([]Task, error) {
	var tasks []Task

	return tasks, d.DB.Where("user_id = ? AND completed_at IS NOT NULL", userID).Find(&tasks).Error
}

func (d *Database) GetUncompletedTasks(userID int) ([]Task, error) {
	var tasks []Task

	return tasks, d.DB.Where("user_id = ? AND completed_at IS NULL", userID).Find(&tasks).Error
}

func (d *Database) CreateTask(userID int, title string, description string, dueDate *time.Time) error {
	return d.DB.Create(&Task{UserID: userID, Title: title, Description: description, DueDate: dueDate}).Error
}

func (d *Database) UpdateTask(taskID int, title string, description string, dueDate *time.Time) error {
	return d.Model(&Task{}).Where("id = ?", taskID).Updates(map[string]interface{}{
		"title": title, "description": description, "due_date": dueDate,
	}).Error
}

func (d *Database) DeleteTask(taskID int) error {
	return d.Model(&Task{}).Where("id = ?", taskID).Delete(&Task{}).Error
}

func (d *Database) CompleteTask(taskID int) error {
	return d.Model(&Task{}).Where("id = ?", taskID).Update("completed_at", time.Now()).Error
}

func (d *Database) UncompleteTask(taskID int) error {
	return d.Model(&Task{}).Where("id = ?", taskID).Update("completed_at", nil).Error
}

type LabelTask struct {
	gorm.Model
	Label   Label
	LabelID int `gorm:"index:label_id_task_id,unique"`
	Task    Task
	TaskID  int `gorm:"index:label_id_task_id,unique"`
}

func (d *Database) GetTasksByLabel(labelID int) ([]Task, error) {
	var tasks []Task

	return tasks, d.DB.Joins("JOIN label_tasks ON label_tasks.task_id = tasks.id").Where("label_tasks.label_id = ?", labelID).Find(&tasks).Error
}

func (d *Database) AddLabelToTask(labelID int, taskID int) error {
	return d.DB.Create(&LabelTask{LabelID: labelID, TaskID: taskID}).Error
}

func (d *Database) RemoveLabelFromTask(labelID int, taskID int) error {
	return d.Model(&LabelTask{}).Where("label_id = ? AND task_id = ?", labelID, taskID).Delete(&LabelTask{}).Error
}

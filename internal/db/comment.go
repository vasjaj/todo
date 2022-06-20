package db

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User        User
	UserID      int
	Task        Task
	TaskID      int
	Title       string
	Description string
}

func (d *db) GetComment(id int) (*Comment, error) {
	comment := &Comment{}

	return comment, d.DB.Where("id = ?", id).First(comment).Error
}

func (d *db) GetComments(taskID int) ([]Comment, error) {
	var comments []Comment

	return comments, d.DB.Where("task_id = ?", taskID).Find(&comments).Error
}

func (d *db) CreateComment(userID int, taskID int, title string, description string) error {
	return d.DB.Create(&Comment{
		UserID:      userID,
		TaskID:      taskID,
		Title:       title,
		Description: description,
	}).Error
}

func (d *db) UpdateComment(id int, title string, description string) error {
	return d.Model(&Comment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title": title, "description": description,
	}).Error
}

func (d *db) DeleteComment(id int) error {
	return d.Model(&Comment{}).Where("id = ?", id).Delete(&Comment{}).Error
}

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

func (d *Database) GetComment(id int) (*Comment, error) {
	comment := &Comment{}

	return comment, d.DB.Where("id = ?", id).First(comment).Error
}

func (d *Database) GetComments(taskID int) ([]Comment, error) {
	var comments []Comment

	return comments, d.DB.Where("task_id = ?", taskID).Find(&comments).Error
}

func (d *Database) CreateComment(userID int, taskID int, title string, description string) error {
	return d.DB.Create(&Comment{
		UserID:      userID,
		TaskID:      taskID,
		Title:       title,
		Description: description,
	}).Error
}

func (d *Database) UpdateComment(id int, title string, description string) error {
	return d.Model(&Comment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title": title, "description": description,
	}).Error
}

func (d *Database) DeleteComment(id int) error {
	return d.Model(&Comment{}).Where("id = ?", id).Delete(&Comment{}).Error
}

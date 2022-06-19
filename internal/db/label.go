package db

import (
	"errors"

	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	User   User
	UserID int
	Title  string `gorm:"unique"`
}

func (d *Database) GetLabel(id int) (*Label, error) {
	label := &Label{}

	return label, d.DB.Where("id = ?", id).First(label).Error
}

func (d *Database) GetLabels(userID int) ([]Label, error) {
	var labels []Label

	return labels, d.DB.Where("user_id = ?", userID).Find(&labels).Error
}

func (d *Database) GetTasksForLabel(labelID int) ([]Task, error) {
	var tasks []Task

	return tasks, d.DB.Joins("JOIN label_tasks ON label_tasks.task_id = tasks.id").Where("label_tasks.label_id = ?", labelID).Find(&tasks).Error
}

func (d *Database) CreateLabel(userID int, title string) error {
	return d.DB.Create(&Label{UserID: userID, Title: title}).Error
}

func (d *Database) UpdateLabel(labelID int, title string) error {
	return d.Model(&Label{}).Where("id = ?", labelID).Updates(map[string]interface{}{"title": title}).Error
}

var ErrLabelHasTasks = errors.New("label has tasks")

func (d *Database) DeleteLabel(labelID int) error {
	var count int64
	if err := d.DB.Model(&Task{}).Where("label_id = ?", labelID).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return ErrLabelHasTasks
	}

	return d.Model(&Label{}).Where("id = ?", labelID).Delete(&Label{}).Error
}

type LabelTask struct {
	gorm.Model
	Label   Label
	LabelID int `gorm:"index:label_id_task_id,unique"`
	Task    Task
	TaskID  int `gorm:"index:label_id_task_id,unique"`
}

func (d *Database) AddLabelToTask(labelID int, taskID int) error {
	return d.DB.Create(&LabelTask{LabelID: labelID, TaskID: taskID}).Error
}

func (d *Database) RemoveLabelFromTask(labelID int, taskID int) error {
	return d.Model(&LabelTask{}).Where("label_id = ? AND task_id = ?", labelID, taskID).Delete(&LabelTask{}).Error
}

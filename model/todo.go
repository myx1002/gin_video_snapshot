package model

import (
	"time"

	"gin_video_snapshot/dao"
)

// model
type TODO struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	Title     string     `json:"title" gorm:"not null;size:32;default:''"`
	Status    bool       `json:"status" gorm:"default:false;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:datetime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"type:datetime"`
}

func InitModel() {
	// 创建table
	err := dao.DB.AutoMigrate(&TODO{})
	if err != nil {
		panic(err.Error())
	}
}

func CreateTodo(todo *TODO) error {
	if err := dao.DB.Debug().Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoList(todoList *[]TODO) error {
	if err := dao.DB.Debug().Find(&todoList).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(id string) (todo *TODO, err error) {
	if err := dao.DB.Debug().Where("id = ?", id).Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTodo(todo *TODO) error {
	if err := dao.DB.Model(todo).Update("status", todo.Status).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(id string, todo *TODO) error {
	if err := dao.DB.Where("id=?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}

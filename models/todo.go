package models

import (
	_ "github.com/jinzhu/gorm"
)

type Todo struct {
	ID          int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Description string `json:"description"`
	Resolved    bool   `json:"resolved"`
}

func GetTodos() ([]Todo, error) {
	var todos []Todo
	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodo(des string) (*Todo, error) {
	todo := Todo{
		Description: des,
		Resolved:    false,
	}

	if err := db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(id int, des string) (*Todo, error) {
	var todo Todo

	err := db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}

	todo.Description = des

	if err := db.Model(&todo).Update(todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func ResolveTodo(id int) (int64, error) {
	var todo Todo

	err := db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return 0, err
	}

	todo.Resolved = true

	if err := db.Model(&todo).Update(todo).Error; err != nil {
		return 0, err
	}

	return todo.ID, nil
}

func DeleteTodo(id int) (int64, error) {
	var todo Todo

	err := db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return 0, err
	}

	if err := db.Delete(&todo).Error; err != nil {
		return 0, err
	}

	return todo.ID, nil
}

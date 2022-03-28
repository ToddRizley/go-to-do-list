package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-to-do-list/config"
)

func GetAllTodoItems(todo *[]TodoItem) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodoItem(todo *TodoItem, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodoItem(todo *TodoItem) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodoItem(todo *TodoItem, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

func DeleteATodoItem(todo *TodoItem, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

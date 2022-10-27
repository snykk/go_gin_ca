package handlers

import (
	"errors"
	"fmt"
	"log"

	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/models"
)

func (m *module) InsertTodo(user_id uint, newTodo datatransfers.TodoInsert) (todo models.Todo, err error) {
	if todo, err = m.db.todoRepository.InsertTodo(
		models.Todo{
			UserID:   user_id,
			Activity: newTodo.Activity,
			Priority: newTodo.Priority,
			IsDone:   false,
		},
	); err != nil {
		log.Print(err)
		return models.Todo{}, fmt.Errorf("error inserting todo. %v", err)
	}
	return
}

func (m *module) GetAllUserTodo(user_id uint) (todo []models.Todo, err error) {
	if todo, err = m.db.todoRepository.GetAllUserTodo(user_id); err != nil {
		return []models.Todo{}, fmt.Errorf("cannot find todos of todo with id %d", user_id)
	}
	return
}

func (m *module) GETTodoByID(user_id uint, id uint, isAdmin bool) (todo models.Todo, err error) {
	if todo, err = m.db.todoRepository.GetTodoByID(id); err != nil {
		return models.Todo{}, fmt.Errorf("cannot find todo with id %d", id)
	}

	if todo.UserID != user_id && !isAdmin {
		return models.Todo{}, fmt.Errorf("you dont have access to read todo with id %d", id)
	}

	return
}

func (m *module) UpdateTodoUser(user_id uint, id uint, isAdmin bool, todo datatransfers.TodoUpdate) (err error) {
	// Get Todo by id
	var t models.Todo
	if t, err = m.db.todoRepository.GetTodoByID(id); err != nil {
		return fmt.Errorf("cannot find todo with id %d", id)
	}

	if t.UserID != user_id && !isAdmin {
		return fmt.Errorf("you dont have access to update todo with id %d", id)
	}

	if err = m.db.todoRepository.UpdateTodo(models.Todo{
		ID:       id,
		Activity: todo.Activity,
		Priority: todo.Priority,
		IsDone:   todo.IsDone,
	}); err != nil {
		return errors.New("cannot update todo")
	}

	return
}

func (m *module) DeleteTodoById(user_id uint, id uint, isAdmin bool) (todo models.Todo, err error) {
	// Get Todo by id
	if todo, err = m.db.todoRepository.GetTodoByID(id); err != nil {
		fmt.Println("nih errornya", err)
		return models.Todo{}, fmt.Errorf("cannot find todo with id %d", id)
	}

	if todo.UserID != user_id && !isAdmin {
		return models.Todo{}, fmt.Errorf("you dont have access to update todo with id %d", id)
	}

	if err = m.db.todoRepository.DeleteTodo(id); err != nil {
		return models.Todo{}, fmt.Errorf("cannot delete todo with id %d", id)
	}

	return
}

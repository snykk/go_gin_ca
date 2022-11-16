package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/snykk/go_gin_ca/constants"
	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/handlers"
	"github.com/snykk/go_gin_ca/models"
	"github.com/snykk/go_gin_ca/utils"
)

func POSTTodo(c *gin.Context) {
	var err error
	var newTodo datatransfers.TodoInsert

	if err = c.ShouldBind(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	// cek if priority is valid or not
	if err := utils.IsPriorityValid(newTodo.Priority); err != nil {
		c.JSON(http.StatusCreated, datatransfers.Response{Status: true, Message: "todo created successfuly", Data: err.Error()})
		return
	}

	var todo models.Todo
	if todo, err = handlers.Handler.InsertTodo(c.GetUint(constants.UserIDKey), newTodo); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Status: false, Message: "failed create new todo"})
		return
	}

	c.JSON(http.StatusCreated, datatransfers.Response{Status: true, Message: "todo created successfuly", Data: todo})
}

func GETAllUserTodo(c *gin.Context) {
	todos, err := handlers.Handler.GetAllUserTodo(c.GetUint(constants.UserIDKey))

	if err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}
	if len(todos) == 0 {
		c.JSON(http.StatusNotFound, datatransfers.Response{Status: false, Message: "book data is empty"})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: "list of user todo fetched successfully", Data: todos})
}

func GETTodoByID(c *gin.Context) {
	var err error
	idParam, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	if todo, err = handlers.Handler.GETTodoByID(c.GetUint(constants.UserIDKey), uint(idParam), c.GetBool(constants.IsAdmin)); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: fmt.Sprintf("user todo with id %d fetched successfully", idParam), Data: todo})
}

func PUTTodoUser(c *gin.Context) {
	var err error
	var todo datatransfers.TodoUpdate
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err = c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	// cek if priority is valid or not
	if err := utils.IsPriorityValid(todo.Priority); err != nil {
		c.JSON(http.StatusCreated, datatransfers.Response{Status: true, Message: "todo created successfuly", Data: err.Error()})
		return
	}

	if err = handlers.Handler.UpdateTodoUser(c.GetUint(constants.UserIDKey), uint(idParam), c.GetBool(constants.IsAdmin), todo); err != nil {
		fmt.Println("ini errornya", err)
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: "todo data updated successfully", Data: todo})
}

func DELETETodo(c *gin.Context) {
	var err error
	idParam, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	if todo, err = handlers.Handler.DeleteTodoById(c.GetUint(constants.UserIDKey), uint(idParam), c.GetBool(constants.IsAdmin)); err != nil {
		fmt.Println("pasti ke sini", err.Error())
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: fmt.Sprintf("todo with id %d deleted successfully", idParam), Data: todo})
}

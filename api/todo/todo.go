package todo

import (
	"github.com/cohack-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTodoRequest struct {
	Description string `json:"description" binding:"required"`
}

type UpdateTodoRequest struct {
	ID          int    `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ModifyTodoRequest struct {
	ID int `json:"id" binding:"required"`
}

type ModifyTodoResponse struct {
	ID int64 `json:"id"`
}

func GetTodo(c *gin.Context) {
	ts, err := models.GetTodos()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, ts)
}

func CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// create the todo
	t, err := models.CreateTodo(req.Description)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, t)
}

func UpdateTodo(c *gin.Context) {
	var req UpdateTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//update todo
	t, err := models.UpdateTodo(req.ID, req.Description)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, t)
}

func ResolveTodo(c *gin.Context) {
	var req ModifyTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// resolve todo
	id, err := models.ResolveTodo(req.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ModifyTodoResponse{ID: id})
}

func DeleteTodo(c *gin.Context) {
	var req ModifyTodoRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// delete todo
	id, err := models.DeleteTodo(req.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ModifyTodoResponse{ID: id})
}

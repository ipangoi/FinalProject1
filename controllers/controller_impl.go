package controllers

import (
	"finalProject1/database"
	"finalProject1/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HandlerImpl struct{}

func NewHandlerImpl() TodoHandler {
	return &HandlerImpl{}
}

// GetAllTodo godoc
// @Summary Get details
// @Description Get details of all todos
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 404 {object} models.Todo
// @Router /todo [get]
func (s *HandlerImpl) GetAllTodo(c *gin.Context) {
	var db = database.GetDB()

	var todo []models.Todo
	err := db.Find(&todo).Error

	if err != nil {
		fmt.Println("Error getting order datas : ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GetOneTodo godoc
// @Summary Get details for a given ID
// @Description Get details of todo corresponding to the input ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "ID of the todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 404 {object} models.Todo
// @Router /todo/:id [get]
func (s *HandlerImpl) GetOneTodo(c *gin.Context) {
	var db = database.GetDB()

	var todoOne models.Todo

	err := db.First(&todoOne, "ID = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": todoOne})
}

// CreateTodo godoc
// @Summary Post details for a given ID
// @Description Post details of todo corresponding to the input ID
// @Tags todo
// @Accept json
// @Produce json
// @Param models.Todo body models.Todo true "create todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 404 {object} models.Todo
// @Router /todo [post]
func (s *HandlerImpl) CreateTodo(c *gin.Context) {
	var db = database.GetDB()

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	now := time.Now()

	todoInput := models.Todo{
		NamaTodo:    input.NamaTodo,
		Description: input.Description,
		Tanggal:     input.Tanggal,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	db.Create(&todoInput)

	c.JSON(http.StatusOK, gin.H{"data": todoInput})
}

// UpdateTodo godoc
// @Summary Update todo identified by the given ID
// @Description Update the todo corresponding to the input ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "ID of the todo to be updated"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Todo
// @Failure 404 {object} models.Todo
// @Router /todo/:id [patch]
func (s *HandlerImpl) UpdateTodo(c *gin.Context) {
	var db = database.GetDB()

	var input models.Todo
	err := db.First(&input, "ID = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	now := time.Now()

	todoInput := models.Todo{
		NamaTodo:    input.NamaTodo,
		Description: input.Description,
		Tanggal:     input.Tanggal,
		UpdatedAt:   now,
	}
	db.Model(&input).Updates(todoInput)
	if err := c.ShouldBindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// DeleteTodo godoc
// @Summary Delete todo identified by the given ID
// @Description Delete the todo corresponding to the input ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "ID of the todo to be deleted"
// @Success 200 {object} models.Todo
// @Router /todo/:id [delete]
func (s *HandlerImpl) DeleteTodo(c *gin.Context) {
	var db = database.GetDB()

	var todoDelete models.Todo
	err := db.First(&todoDelete, "ID = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&todoDelete)

	c.JSON(http.StatusOK, gin.H{"data": "Data deleted"})
}

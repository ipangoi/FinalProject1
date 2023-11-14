package controllers

import "github.com/gin-gonic/gin"

type TodoHandler interface {
	GetAllTodo(*gin.Context)
	GetOneTodo(*gin.Context)
	CreateTodo(*gin.Context)
	UpdateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}

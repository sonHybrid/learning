package handlers

import (
	"net/http"
	"todovue/dbase"
	"todovue/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoInterface interface {
	TemplateIndex(c *gin.Context)
	GetListTodo(c *gin.Context)
	AddTodo(c *gin.Context)
}

type todoHandler struct {
	dba *gorm.DB
}

func NewTodoHandler() TodoInterface {

	return &todoHandler{
		dba: dbase.DB(),
	}
}

func (t *todoHandler) TemplateIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Todos",
	})
}

func (t *todoHandler) GetListTodo(c *gin.Context) {
	var todos = []model.Todo{}
	t.dba.Table("public.todos as u").
		Select("*").
		Order("u.id  Desc").Scan(&todos)
	c.JSON(200, gin.H{
		"listdata": todos,
	})
}

func (t *todoHandler) AddTodo(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	var todo = model.Todo{Name: name, Description: description, Done: false}
	t.dba.Create(&todo)
	c.JSON(200, gin.H{})
}

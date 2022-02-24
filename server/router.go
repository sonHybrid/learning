package server

import (
	Home "todovue/handlers"

	"github.com/gin-gonic/gin"
)

var (
	home Home.TodoInterface
)

func router(app *gin.Engine) {
	home = Home.NewTodoHandler()

	app.GET("/ping2", home.TemplateIndex)
	app.GET("/listtodo", home.GetListTodo)
	app.POST("/addTodo", home.AddTodo)

}

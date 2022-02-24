package server

import "github.com/gin-gonic/gin"

func Router() error {
	app := setup()
	return app.Run(":3000")
}

func setup() *gin.Engine {
	app := gin.Default()
	app.Delims("[[", "]]")
	app.Static("/assets", "./assets")
	app.LoadHTMLGlob("templates/*.html")
	router(app)
	return app
}

package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/stephstephg/learn-go/golang-angular/handlers"
)

/*
add a new to-do item : curl localhost:3000/todo -d '{"message": "finish writing the article"}'
get all to-do items : curl localhost:3000/todo

*/
func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.GET("/todo", handlers.GetTodoHandler)
	r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

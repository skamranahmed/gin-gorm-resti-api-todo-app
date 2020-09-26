package Controllers

import (
	"Volumes/Kamran/go-lang-todo-app/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all todos
func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateATodo Function
func CreateATodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)

	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get a particular Todo with id
func GetATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")

	err := Models.GetATodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

// Update an existing todo
func UpdateATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")

	err := Models.GetATodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// Delete a todo
func DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")

	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "deleted"})
	}

}

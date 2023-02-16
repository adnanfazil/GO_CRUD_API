package main

import "strconv"

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
    router := gin.Default()

    router.GET("/todos", getTodos)
    router.GET("/todos/:id", getTodo)
    router.POST("/todos", addTodo)
    router.PUT("/todos/:id", updateTodo)
    router.DELETE("/todos/:id", deleteTodo)

    router.Run(":8080")
}

func getTodos(c *gin.Context) {
    c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    for _, todo := range todos {
        if todo.ID == id {
            c.JSON(http.StatusOK, todo)
            return
        }
    }

    c.AbortWithStatus(http.StatusNotFound)
}

func addTodo(c *gin.Context) {
    var todo Todo
    if err := c.BindJSON(&todo); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    todo.ID = len(todos) + 1
    todos = append(todos, todo)

    c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    var updatedTodo Todo
    if err := c.BindJSON(&updatedTodo); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    for i, todo := range todos {
        if todo.ID == id {
            updatedTodo.ID = id
            todos[i] = updatedTodo
            c.JSON(http.StatusOK, updatedTodo)
            return
        }
    }

    c.AbortWithStatus(http.StatusNotFound)
}

func deleteTodo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    for i, todo := range todos {
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }

    c.AbortWithStatus(http.StatusNotFound)
}

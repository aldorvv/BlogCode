package views

import (
	"net/http"
	"strconv"

	"github.com/BlogCode/MultiContainer/internal/db"
	"github.com/BlogCode/MultiContainer/internal/db/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	const query = "SELECT * FROM todos WHERE isActive=true;"
	repo := db.GetRepo()

	result, err := repo.Query(query)
	var todos []models.Todo

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var todo models.Todo
		err = result.Scan(&todo.ID, &todo.Description, &todo.IsComplete, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	const query = "SELECT * FROM todos WHERE id=? AND isActive=true;"
	repo := db.GetRepo()
	id := c.Param("id")
	var todo models.Todo

	result, err := repo.Query(query, id)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err = result.Scan(&todo.ID, &todo.Description, &todo.IsComplete, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func PostTodo(c *gin.Context) {
	var newTodo models.Todo
	repo := db.GetRepo()
	const query = "INSERT INTO todos (`description`) VALUES (?);"

	if err := c.BindJSON(&newTodo); err != nil {
		panic(err.Error())
	}

	_, err := repo.Query(query, newTodo.Description)

	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusCreated, newTodo)

}

func UpdateTodo(c *gin.Context) {
	repo := db.GetRepo()
	const query = "UPDATE todos SET isComplete = ? WHERE id=?;"
	id := c.Param("id")
	statusParam := c.DefaultQuery("completed", "0")

	status, err := strconv.Atoi(statusParam)

	if err != nil {
		panic(err.Error())
	}

	_, err = repo.Query(query, status, id)

	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, nil)

}

func SoftDeleteTodo(c *gin.Context) {
	repo := db.GetRepo()
	const query = "UPDATE todos SET isActive=false WHERE id=?;"
	id := c.Param("id")

	_, err := repo.Query(query, id)

	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, nil)

}

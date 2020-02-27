package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

/*
	Create User
	curl -X POST \
	  -H 'Content-Type: application/json' \
	  -d '{"name":"Joe Smith"}' \
	  localhost:1323/users
	Response

	{
	  "id": 1,
	  "name": "Joe Smith"
	}

	Get User
	curl localhost:1323/users/1
	Response

	{
	  "id": 1,
	  "name": "Joe Smith"
	}

	Update User
	curl -X PUT \
	  -H 'Content-Type: application/json' \
	  -d '{"name":"Joe"}' \
	  localhost:1323/users/1
	Response

	{
	  "id": 1,
	  "name": "Joe"
	}

	Delete User
	curl -X DELETE localhost:1323/users/1
*/

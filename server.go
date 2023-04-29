package main

import (
	"echo-practice/animal"
	"echo-practice/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	u := user.New(1, "testUser", 20)

	e.GET("/user", u.Get)
	e.POST("/user", user.Post)

	dog := animal.New("dog")

	e.GET("/animal/dog", dog.Get)
	e.POST("/animal/dog", dog.Post)

	cat := animal.New("cat")

	e.GET("/animal/cat", cat.Get)
	e.POST("/animal/cat", cat.Post)

	e.Logger.Fatal(e.Start(":1323"))
}

package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Age      int64  `json:"age"`
}

func New(id int64, userName string, age int64) *User {
	return &User{id, userName, age}
}

func (u *User) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, u)
}

func Post(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}

	log.Printf("User:%s", string(userData))

	jsonData, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	})
	return c.String(http.StatusCreated, string(jsonData))
}

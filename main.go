package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

type User struct {
    Name  string `json:"name"`
}




func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        user := User{Name: "hello eotld"}
        return c.JSON(http.StatusCreated, user)
    })

    e.Run(standard.New(":5000"))
}

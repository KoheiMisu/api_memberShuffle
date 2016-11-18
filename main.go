package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "api_memberShuffle/controllers"
    "api_memberShuffle/db"
)

func main() {
    e := echo.New()
    e.Use(middleware.CORS())

    defer db.Mongo.Close()

    e.POST("/api/v1/members", controllers.PostMember)

    e.GET("/api/v1/members", controllers.FetchMembers)

    e.PUT("/api/v1/members", controllers.PutMember)

    e.Logger.Fatal(e.Start(":5000"))
}

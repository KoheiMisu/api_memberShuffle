package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"

    "ReceitAnalysisApi/controllers"
    "ReceitAnalysisApi/db"
)

type Member struct {
    Name  string `json:"name" bson:"name"`
    Present bool `json:"present, string" bson:"present"`
}

func main() {
    e := echo.New()
    e.Use(middleware.CORS())

    defer db.Mongo.Close()

    e.Post("/api/v1/members", controllers.PostMember)

    e.Get("/api/v1/members", controllers.FetchMembers)

    e.Put("/api/v1/members", controllers.PutMember)

    e.Run(standard.New(":5000"))
}

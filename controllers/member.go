package controllers

import (
    "net/http"
    "github.com/labstack/echo"
    "labix.org/v2/mgo/bson"

    "ReceitAnalysisApi/db"
    "strconv"
)

type Member struct {
    Name  string `json:"name"`
    Present bool `json:"present"`
}

func FetchMembers(c echo.Context) error {

    cId := c.QueryParam("cId")

    var results []Member

    err := db.Mongo.DB("member_shuffles").C(cId).Find(nil).All(&results)
    if err != nil {
        panic(err)
    }

    return c.JSON(http.StatusCreated, results)
}

func PostMember(c echo.Context) error {

    collectionId := c.FormValue("cId")

    con := db.Mongo.DB("member_shuffles").C(collectionId)

    member := Member{Name: c.FormValue("name"), Present: true}

    err := con.Insert(member)
    if err != nil {
        panic(err)
    }

    return c.JSON(http.StatusCreated, member)
}

func PutMember(c echo.Context) error {

    cId := c.FormValue("cId")

    con := db.Mongo.DB("member_shuffles").C(cId)

    name := c.FormValue("name")

    var present bool
    present, _ = strconv.ParseBool(c.FormValue("present"))

    member := Member{Name: name, Present: present}

    query := bson.M{"name": name}

    err := con.Update(query, member)
    if err != nil {
        panic(err)
    }

    return c.JSON(http.StatusCreated, member)
}
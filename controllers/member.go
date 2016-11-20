package controllers

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "labix.org/v2/mgo/bson"

    "api_memberShuffle/db"
    "api_memberShuffle/util"
)

type PostResult struct {
    Member `json:"member"`
    Validate
}

type Member struct {
    Name  string `json:"name" bson:"name"`
    Present bool `json:"present" bson:"present"`
}

type Validate struct {
    Result bool `json:"result"`
    Message string `json:"message"`
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

    name := c.FormValue("name");

    member := Member{
        Name: name,
        Present: true,
    }

    result, message := util.ValidateMember(name)

    if !result {
        validate := Validate {
            Result: result,
            Message: message,
        }
        return c.JSON(http.StatusCreated, validate)
    }

    isUnique := con.Find(bson.M{"name": name}).One(&member)
    if isUnique == nil {
        validate := Validate {
            Result: false,
            Message: name+" is already registered",
        }
        return c.JSON(http.StatusCreated, validate)
    }

    err := con.Insert(member)
    if err != nil {
        panic(err)
    }

    postResult := PostResult{
        Member: member,
        Validate: Validate {
            Result: true,
            Message: "",
        },
    }

    return c.JSON(http.StatusCreated, postResult)
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

func DeleteMember(c echo.Context) error {

    cId := c.FormValue("cId")

    con := db.Mongo.DB("member_shuffles").C(cId)

    name := c.FormValue("name")

    var present bool
    present, _ = strconv.ParseBool(c.FormValue("present"))

    member := Member{Name: name, Present: present}

    query := bson.M{"name": name}

    err := con.Remove(query)
    if err != nil {
        panic(err)
    }

    return c.JSON(http.StatusCreated, member)
}


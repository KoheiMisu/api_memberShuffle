package db

import mgo "gopkg.in/mgo.v2"

var Mongo, _ = mgo.Dial("mongo")
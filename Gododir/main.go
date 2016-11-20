package main

import (
    godo "gopkg.in/godo.v2"
)

func Tasks(p *godo.Project) {
    p.Task("build", nil, func(c *godo.Context) {
        c.Run("go build main.go")
    }).Src("*/*.go")

    p.Task("server", nil, func(c *godo.Context) {
        // rebuilds and restarts when a watched file changes
        c.Start("main.go", godo.M{"$in": "./"})
    }).Src("*.go", "**/*.go").
        Debounce(3000)
}

func main() {
    godo.Godo(Tasks)
}

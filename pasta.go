package main

import (
	"github.com/codegangsta/martini" // DEP
	"fmt"
)

var (
	version string
)


func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return fmt.Sprintf("Pasta %s", version)
	})

	m.Run()

}

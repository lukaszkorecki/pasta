package main

import (
	"fmt"
	"github.com/codegangsta/martini" // DEP
	es "github.com/lukaszkorecki/pasta/expiring_store"
	"net/http"
)

var (
	version string
	store = es.New(0) // uses default 5 minute time
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return fmt.Sprintf("Pasta %s", version)
	})

	m.Get("/:key", func(params martini.Params) (int, string) {
		text, exists := store.Get(params["key"])

		if exists {
			return 200, text
		} else {
			return 404, "Gone"
		}
	})

	m.Post("/paste", func(_, r *http.Request) (int, string) {
		body := r.FormValue("body")

		key := store.Set(body)

		return 201, key
	})

	m.Run()

}

package main

import (
	"bluecanvas.io/auth0-api/api"
	"github.com/rcrowley/go-tigertonic"
)

func main() {
	c, err := api.NewContext()
	if err != nil {
		panic(err)
	}

	tigertonic.NewServer(":8000", tigertonic.Logged(api.Routes(c), nil)).ListenAndServe()
}

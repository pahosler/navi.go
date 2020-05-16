package main

import "github.com/noriah/navi.go/app"

func main() {
	if err := app.Go(); err != nil {
		panic(err)
	}
}

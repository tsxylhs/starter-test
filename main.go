package main

import (
	"code_test/app"

	starter "github.com/tsxylhs/go-starter/app"
)

func main() {
	starter.RegisterStarter(app.Server)
	if err := starter.Start(); err != nil {
		panic("error")
	}
}

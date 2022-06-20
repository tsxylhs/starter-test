package main

import (
	"code_test/app"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	starter "github.com/tsxylhs/go-starter/starter"

	"golang.org/x/sync/errgroup"
)

func main() {
	app.WebApi.Mount(app.Server)
	starter.RegisterStarter(app.WebApi)
	err := starter.Start()
	if err != nil {
		panic(err)
	}
	apiServer := gin.New()
	apiServer.Use(gin.Logger())
	apiServer.Use(gin.Recovery())
	var g errgroup.Group
	fmt.Println("app.WebApi.Port", app.WebApi.Port)
	g.Go(func() error {
		return apiServer.Run(":" + app.WebApi.Port)
	})

	if err := g.Wait(); err != nil {
		panic(err)
	} else {
		fmt.Println("started success")
		fmt.Println("dbDesc", app.WebApi.DB.Desc())
	}

}

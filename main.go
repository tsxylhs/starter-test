package main

import (
	"code_test/app"
	"code_test/rest"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tsxylhs/go-starter/starter"
	"github.com/tsxylhs/go-starter/web"

	"golang.org/x/sync/errgroup"
)

func main() {
	app.WebApi.Mount(app.Server)
	starter.RegisterStarter(app.WebApi)
	err := starter.Start()
	if err != nil {
		panic(err)
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	log.Logger.Logger.Info("测试线程")
	// 	time.Sleep(2000022)
	// 	wg.Done()

	// }(&wg)
	// wg.Wait()
	apiServer := gin.New()
	apiServer.Use(gin.Logger())
	apiServer.Use(gin.Recovery())
	//跨域过滤器
	apiServer.Use(web.CorsHandler(app.WebApi))

	pcApi := apiServer.Group("/api")
	rest.RegisterAPIs(pcApi)
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

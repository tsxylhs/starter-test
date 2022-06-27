package main

import (
	"bytes"
	"code_test/app"
	"code_test/rest"
	"fmt"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tsxylhs/go-starter/log"
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
	// var result = make(chan interface{}, 5)
	// var task = task.Task{
	// 	CountPool:   5,
	// 	Name:        "测试",
	// 	Description: []string{"简单并发模型"},
	// 	Do:          dosomething,
	// 	Results:     result,
	// }
	//task.Task3()

	// for {
	// 	v, ok := <-task.Results

	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("task 读到数据=%v\n", v)
	// }
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
		log.Logger.Logger.Info("started success")
		//log.Logger.Logger.Info("dbDesc {}", app.WebApi.DB.Desc().to)
	}

}
func dosomething() interface{} {
	for i := 1; i < 10; i++ {
		fmt.Println("gid", goID())
		fmt.Println("i", i)

	}
	return "你好"

}

func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

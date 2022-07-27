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
	app.WebApi.Mount(app.Server, app.MqttClient, app.TcpClient)

	starter.RegisterStarter(app.WebApi)
	err := starter.Start()
	if err != nil {
		panic(err)
	}
	apiServer := gin.New()
	apiServer.Use(gin.Logger())
	apiServer.Use(gin.Recovery())
	//跨域过滤器
	apiServer.Use(web.CorsHandler(app.WebApi))

	pcApi := apiServer.Group("/api")
	rest.RegisterAPIs(pcApi)
	var g errgroup.Group
	client := app.MqttClient.SharedBroker
	if client.IsConnected() {
		log.Logger.Logger.Info("mqtt is connect success")
	}
	tcpClient := app.TcpClient.SharedBroker.ConnTcp
	tcpClient.Write([]byte("1212"))
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

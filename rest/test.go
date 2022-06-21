package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/tsxylhs/go-starter/web"
)

type test struct {
	*web.RestHandler
}

func NewTestAPI() *test {
	return &test{
		RestHandler: web.DefaultRestHandler,
	}
}
func (api *test) Get(c *gin.Context) {
	c.JSON(200, "haahah")
}
func (api *test) Register(router gin.IRouter) {
	r := router.Group("/v1")
	r.GET("/get", api.Get) // 获取单条记录

}

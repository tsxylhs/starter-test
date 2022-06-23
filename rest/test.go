package rest

import (
	"code_test/module"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsxylhs/go-starter/log"
	"github.com/tsxylhs/go-starter/validator"
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
func (api *test) PostTest(c *gin.Context) {
	tests := &module.Test{}
	if err := c.ShouldBind(tests); err != nil {
		log.Logger.Logger.Info("test")
	}
	code, errmessage := validator.BindAndValid(c, tests)
	if code != http.StatusOK {
		c.JSON(code, errmessage)
		return
	}
	c.JSON(200, tests)
}
func (api *test) Register(router gin.IRouter) {
	r := router.Group("/v1")
	r.GET("/get", api.Get) // 获取单条记录
	r.POST("/POST", api.PostTest)

}

package rest

import "github.com/gin-gonic/gin"

func RegisterAPIs(router gin.IRouter) {
	NewTestAPI().Register(router)
}

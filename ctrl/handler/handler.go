package handler

import (
	"fmt"
	"irita-api/common"
	"irita-api/seelog"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Handler 接口处理接口
type Handler interface {
	BindReq(c *gin.Context) error
	AfterBindReq() error
	GetResponse() interface{}
	SetError(code int, message string)
	Process()
}

// Handle 接口处理
func Handle(c *gin.Context, hd Handler) {
	defer func() {
		if err := recover(); err != nil {
			hd.SetError(common.ErrorPanic, fmt.Sprintf("%v", err))
			seelog.Errorf("panic: %v", err)
			stack := debug.Stack()
			seelog.Errorf("panic stack: %s", string(stack))
		}
		c.JSON(http.StatusOK, hd.GetResponse())
	}()
	if err := hd.BindReq(c); err != nil {
		return
	}
	if err := hd.AfterBindReq(); err != nil {
		return
	}
	hd.Process()
}

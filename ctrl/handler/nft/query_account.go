package nft

import (
	"fmt"
	"irita-api/common"
	"irita-api/ctrl/handler"
	"irita-api/irita"
	"irita-api/seelog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// QueryAccountHandle 查询账户处理方法
func QueryAccountHandle(c *gin.Context) {
	hd := &QueryAccountHandler{}
	handler.Handle(c, hd)
}

// QueryAccountHandler 查询账户处理对象
type QueryAccountHandler struct {
	Req  QueryAccountRequest
	Resp QueryAccountResponse
}

// QueryAccountRequest 查询账户请求
type QueryAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// QueryAccountResponse 查询账户返回
type QueryAccountResponse struct {
	common.BaseResponse
	Address string `json:"address"`
}

// BindReq 绑定参数
func (h *QueryAccountHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *QueryAccountHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *QueryAccountHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *QueryAccountHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *QueryAccountHandler) Process() {
	address, err := irita.ClientInstance.QueryAccount(h.Req.Name, h.Req.Password)
	if err != nil {
		msg := fmt.Sprintf("query account error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.Address = address
}

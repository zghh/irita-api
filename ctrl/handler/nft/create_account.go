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

// CreateAccountHandle 创建账户处理方法
func CreateAccountHandle(c *gin.Context) {
	hd := &CreateAccountHandler{}
	handler.Handle(c, hd)
}

// CreateAccountHandler 创建账户处理对象
type CreateAccountHandler struct {
	Req  CreateAccountRequest
	Resp CreateAccountResponse
}

// CreateAccountRequest 创建账户请求
type CreateAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mnemonic string `json:"mnemonic"`
}

// CreateAccountResponse 创建账户返回
type CreateAccountResponse struct {
	common.BaseResponse
	Address  string `json:"address"`
	Mnemonic string `json:"mnemonic"`
}

// BindReq 绑定参数
func (h *CreateAccountHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *CreateAccountHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *CreateAccountHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *CreateAccountHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *CreateAccountHandler) Process() {
	adress, mnemonic, err := irita.ClientInstance.CreateAccount(h.Req.Name, h.Req.Password, h.Req.Mnemonic)
	if err != nil {
		msg := fmt.Sprintf("create account error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.Address = adress
	h.Resp.Mnemonic = mnemonic
}

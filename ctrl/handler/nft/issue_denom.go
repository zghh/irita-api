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

// IssueDenomHandle 创建NFT类别处理方法
func IssueDenomHandle(c *gin.Context) {
	hd := &IssueDenomHandler{}
	handler.Handle(c, hd)
}

// IssueDenomHandler 创建NFT类别处理对象
type IssueDenomHandler struct {
	Req  IssueDenomRequest
	Resp IssueDenomResponse
}

// IssueDenomRequest 创建NFT类别请求
type IssueDenomRequest struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Schema string `json:"schema" binding:"required"`
}

// IssueDenomResponse 创建NFT类别返回
type IssueDenomResponse struct {
	common.BaseResponse
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Hash      string `json:"hash"`
	Height    int64  `json:"height"`
}

// BindReq 绑定参数
func (h *IssueDenomHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *IssueDenomHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *IssueDenomHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *IssueDenomHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *IssueDenomHandler) Process() {
	tx, err := irita.ClientInstance.IssueDenom(h.Req.ID, h.Req.Name, h.Req.Schema)
	if err != nil {
		msg := fmt.Sprintf("issue denom error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.GasWanted = tx.GasWanted
	h.Resp.GasUsed = tx.GasUsed
	h.Resp.Hash = tx.Hash
	h.Resp.Height = tx.Height
}

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

// QueryDenomHandle 查询NFT类别处理方法
func QueryDenomHandle(c *gin.Context) {
	hd := &QueryDenomHandler{}
	handler.Handle(c, hd)
}

// QueryDenomHandler 查询NFT类别处理对象
type QueryDenomHandler struct {
	Req  QueryDenomRequest
	Resp QueryDenomResponse
}

// QueryDenomRequest 查询NFT类别请求
type QueryDenomRequest struct {
	ID string `json:"id" binding:"required"`
}

// QueryDenomResponse 查询NFT类别返回
type QueryDenomResponse struct {
	common.BaseResponse
	ID      string `json:"id"`
	Name    string `json:"name"`
	Schema  string `json:"schema"`
	Creator string `json:"creator"`
}

// BindReq 绑定参数
func (h *QueryDenomHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *QueryDenomHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *QueryDenomHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *QueryDenomHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *QueryDenomHandler) Process() {
	denom, err := irita.ClientInstance.QueryDenom(h.Req.ID)
	if err != nil {
		msg := fmt.Sprintf("query denom error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.ID = denom.ID
	h.Resp.Name = denom.Name
	h.Resp.Schema = denom.Schema
	h.Resp.Creator = denom.Creator
}

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

// QueryNFTHandle 查询NFT处理方法
func QueryNFTHandle(c *gin.Context) {
	hd := &QueryNFTHandler{}
	handler.Handle(c, hd)
}

// QueryNFTHandler 查询NFT处理对象
type QueryNFTHandler struct {
	Req  QueryNFTRequest
	Resp QueryNFTResponse
}

// QueryNFTRequest 查询NFT请求
type QueryNFTRequest struct {
	ID      string `json:"id" binding:"required"`
	DenomID string `json:"denomId" binding:"required"`
}

// QueryNFTResponse 查询NFT返回
type QueryNFTResponse struct {
	common.BaseResponse
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
}

// BindReq 绑定参数
func (h *QueryNFTHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *QueryNFTHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *QueryNFTHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *QueryNFTHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *QueryNFTHandler) Process() {
	nft, err := irita.ClientInstance.QueryNFT(h.Req.DenomID, h.Req.ID)
	if err != nil {
		msg := fmt.Sprintf("query nft error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.ID = nft.ID
	h.Resp.Name = nft.Name
	h.Resp.URI = nft.URI
	h.Resp.Data = nft.Data
	h.Resp.Creator = nft.Creator
}

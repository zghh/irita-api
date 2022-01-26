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

// QueryOwnNFTsHandle 查询拥有的NFT列表处理方法
func QueryOwnNFTsHandle(c *gin.Context) {
	hd := &QueryOwnNFTsHandler{}
	handler.Handle(c, hd)
}

// QueryOwnNFTsHandler 查询拥有的NFT列表处理对象
type QueryOwnNFTsHandler struct {
	Req  QueryOwnNFTsRequest
	Resp QueryOwnNFTsResponse
}

// QueryOwnNFTsRequest 查询拥有的NFT列表请求
type QueryOwnNFTsRequest struct {
	Address string `json:"address" binding:"required"`
	DenomID string `json:"denomId" binding:"required"`
}

// QueryOwnNFTsResponse 查询拥有的NFT列表返回
type QueryOwnNFTsResponse struct {
	common.BaseResponse
	IDs []string `json:"ids"`
}

// BindReq 绑定参数
func (h *QueryOwnNFTsHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *QueryOwnNFTsHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *QueryOwnNFTsHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *QueryOwnNFTsHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *QueryOwnNFTsHandler) Process() {
	resp, err := irita.ClientInstance.QueryOwner(h.Req.Address, h.Req.DenomID)
	if err != nil {
		msg := fmt.Sprintf("query owner error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.IDs = make([]string, 0)
	for _, idc := range resp.IDCs {
		for _, id := range idc.TokenIDs {
			h.Resp.IDs = append(h.Resp.IDs, id)
		}
	}
}

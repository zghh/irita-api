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

// MintNFTHandle 创建NFT处理方法
func MintNFTHandle(c *gin.Context) {
	hd := &MintNFTHandler{}
	handler.Handle(c, hd)
}

// MintNFTHandler 创建NFT处理对象
type MintNFTHandler struct {
	Req  MintNFTRequest
	Resp MintNFTResponse
}

// MintNFTRequest 创建NFT请求
type MintNFTRequest struct {
	DenomID string `json:"denomId" binding:"required"`
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	URI     string `json:"uri" binding:"required"`
	Data    string `json:"data" binding:"required"`
}

// MintNFTResponse 创建NFT返回
type MintNFTResponse struct {
	common.BaseResponse
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Hash      string `json:"hash"`
	Height    int64  `json:"height"`
}

// BindReq 绑定参数
func (h *MintNFTHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *MintNFTHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *MintNFTHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *MintNFTHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *MintNFTHandler) Process() {
	tx, err := irita.ClientInstance.MintNFT(h.Req.DenomID, h.Req.ID, h.Req.Name, h.Req.URI, h.Req.Data)
	if err != nil {
		msg := fmt.Sprintf("mint nft error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.GasWanted = tx.GasWanted
	h.Resp.GasUsed = tx.GasUsed
	h.Resp.Hash = tx.Hash
	h.Resp.Height = tx.Height
}

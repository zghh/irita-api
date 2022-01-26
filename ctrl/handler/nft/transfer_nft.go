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

// TransferNFTHandle 转移NFT处理方法
func TransferNFTHandle(c *gin.Context) {
	hd := &TransferNFTHandler{}
	handler.Handle(c, hd)
}

// TransferNFTHandler 转移NFT处理对象
type TransferNFTHandler struct {
	Req  TransferNFTRequest
	Resp TransferNFTResponse
}

// TransferNFTRequest 转移NFT请求
type TransferNFTRequest struct {
	DenomID   string `json:"denomId" binding:"required"`
	ID        string `json:"id" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
}

// TransferNFTResponse 转移NFT返回
type TransferNFTResponse struct {
	common.BaseResponse
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Hash      string `json:"hash"`
	Height    int64  `json:"height"`
}

// BindReq 绑定参数
func (h *TransferNFTHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *TransferNFTHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *TransferNFTHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *TransferNFTHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *TransferNFTHandler) Process() {
	tx, err := irita.ClientInstance.TransferNFT(h.Req.DenomID, h.Req.ID, h.Req.Recipient)
	if err != nil {
		msg := fmt.Sprintf("transfer nft error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.GasWanted = tx.GasWanted
	h.Resp.GasUsed = tx.GasUsed
	h.Resp.Hash = tx.Hash
	h.Resp.Height = tx.Height
}

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

// QueryNFTsHandle 查询NFT列表处理方法
func QueryNFTsHandle(c *gin.Context) {
	hd := &QueryNFTsHandler{}
	handler.Handle(c, hd)
}

// QueryNFTsHandler 查询NFT列表处理对象
type QueryNFTsHandler struct {
	Req  QueryNFTsRequest
	Resp QueryNFTsResponse
}

// QueryNFTsRequest 查询NFT列表请求
type QueryNFTsRequest struct {
	DenomID string `json:"denomId" binding:"required"`
}

// QueryNFTsResponse 查询NFT列表返回
type QueryNFTsResponse struct {
	common.BaseResponse
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Schema  string    `json:"schema"`
	Creator string    `json:"creator"`
	NFTs    []NFTData `json:"nfts"`
}

// NFTData nft数据
type NFTData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
}

// BindReq 绑定参数
func (h *QueryNFTsHandler) BindReq(c *gin.Context) error {
	if err := c.ShouldBindBodyWith(&h.Req, binding.JSON); err != nil {
		msg := fmt.Sprintf("invalid request, bind error: %v", err)
		seelog.Errorf(msg)
		h.SetError(common.ErrorPanic, msg)
		return err
	}
	return nil
}

// AfterBindReq 绑定参数后校验参数
func (h *QueryNFTsHandler) AfterBindReq() error {
	return nil
}

// GetResponse 获取返回信息
func (h *QueryNFTsHandler) GetResponse() interface{} {
	return h.Resp
}

// SetError 设置错误信息
func (h *QueryNFTsHandler) SetError(code int, message string) {
	h.Resp.Code = code
	h.Resp.Message = message
}

// Process 接口处理
func (h *QueryNFTsHandler) Process() {
	resp, err := irita.ClientInstance.QueryNFTs(h.Req.DenomID)
	if err != nil {
		msg := fmt.Sprintf("query nfts error, %v", err)
		h.SetError(common.ErrorInner, msg)
		seelog.Errorf(msg)
		return
	}
	h.Resp.ID = resp.Denom.ID
	h.Resp.Name = resp.Denom.Name
	h.Resp.Schema = resp.Denom.Schema
	h.Resp.Creator = resp.Denom.Creator
	h.Resp.NFTs = make([]NFTData, len(resp.NFTs))
	for i, nft := range resp.NFTs {
		h.Resp.NFTs[i].ID = nft.ID
		h.Resp.NFTs[i].Name = nft.Name
		h.Resp.NFTs[i].URI = nft.URI
		h.Resp.NFTs[i].Data = nft.Data
		h.Resp.NFTs[i].Creator = nft.Creator
	}
}

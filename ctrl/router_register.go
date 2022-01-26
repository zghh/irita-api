package ctrl

import (
	"irita-api/common"
	"irita-api/ctrl/handler/nft"
)

func init() {
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/createAccount",
		Method:  "POST",
		Handler: nft.CreateAccountHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/issueDenom",
		Method:  "POST",
		Handler: nft.IssueDenomHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/mintNFT",
		Method:  "POST",
		Handler: nft.MintNFTHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/transferNFT",
		Method:  "POST",
		Handler: nft.TransferNFTHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/queryAccount",
		Method:  "POST",
		Handler: nft.QueryAccountHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/queryDenom",
		Method:  "POST",
		Handler: nft.QueryDenomHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/queryNFT",
		Method:  "POST",
		Handler: nft.QueryNFTHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/queryNFTs",
		Method:  "POST",
		Handler: nft.QueryNFTsHandle,
	})
	common.RouterRegister.RegisterRouterHandler(common.RouterHandler{
		Path:    "/api/v1/queryOwnNFTs",
		Method:  "POST",
		Handler: nft.QueryOwnNFTsHandle,
	})
}

package irita

import (
	"fmt"
	"irita-api/conf"
	"irita-api/seelog"
	"sync"

	iritasdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/modules/nft"
	"github.com/bianjieai/irita-sdk-go/types"
	"github.com/bianjieai/irita-sdk-go/types/store"
	opb "github.com/bianjieai/opb-sdk-go/pkg/app/sdk"
	"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/model"
)

// ClientInstance 客户端实例
var ClientInstance *Client

var initOnce sync.Once
var initError error

// Client 客户端
type Client struct {
	c *iritasdk.IRITAClient
}

// InitClient 客户端初始化
func InitClient() error {
	initOnce.Do(func() {
		initError = initClient()
	})
	return initError
}

func initClient() error {
	fee, err := types.ParseCoin(conf.Conf.IritaConf.Coin)
	if err != nil {
		return fmt.Errorf("parse coin error, %v", err)
	}
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.FeeOption(types.NewDecCoinsFromCoins(fee)),
	}
	cfg, err := types.NewClientConfig(conf.Conf.IritaConf.RPCAddress, conf.Conf.IritaConf.WSAddress, conf.Conf.IritaConf.GRPCAddress, conf.Conf.IritaConf.ChainID, options...)
	if err != nil {
		return fmt.Errorf("client config error, %v", err)
	}
	authToken := model.NewAuthToken(conf.Conf.IritaConf.ProjectID, conf.Conf.IritaConf.ProjectKey, conf.Conf.IritaConf.ChainAccountAddress)
	authToken.SetRequireTransportSecurity(conf.Conf.IritaConf.TLSEnable)
	client := opb.NewClient(cfg, &authToken)
	client.SetLogger(&Logger{})
	if _, err := client.Key.Recover(conf.Conf.IritaConf.Username, conf.Conf.IritaConf.Password, conf.Conf.IritaConf.Mnemonic); err != nil {
		seelog.Errorf("recover key error, %v", err)
	}
	ClientInstance = &Client{&client}
	return nil
}

// CreateAccount 创建账户
func (c *Client) CreateAccount(name, password, mnemonic string) (string, string, error) {
	if mnemonic != "" {
		address, err := c.c.Key.Recover(name, password, mnemonic)
		if err != nil {
			return "", "", fmt.Errorf("recover key error, %v", err)
		}
		return address, mnemonic, nil
	}

	address, mnemonic, err := c.c.Key.Add(name, password)
	if err != nil {
		return "", "", fmt.Errorf("recover key error, %v", err)
	}
	return address, mnemonic, nil
}

// IssueDenom 创建NFT类别
func (c *Client) IssueDenom(
	ID, name, schema string,
) (*types.ResultTx, error) {
	baseTx := types.BaseTx{
		From:     conf.Conf.IritaConf.Username,
		Password: conf.Conf.IritaConf.Password,
		Gas:      conf.Conf.IritaConf.GasLimit,
		Memo:     "",
		Mode:     types.Commit,
	}
	tx, err := c.c.NFT.IssueDenom(nft.IssueDenomRequest{ID: ID, Name: name, Schema: schema}, baseTx)
	if err != nil {
		return nil, fmt.Errorf("issue denom error, %v", err)
	}
	return &tx, nil
}

// MintNFT 创建NFT
func (c *Client) MintNFT(
	denomID, ID, name, uri, data string,
) (*types.ResultTx, error) {
	baseTx := types.BaseTx{
		From:     conf.Conf.IritaConf.Username,
		Password: conf.Conf.IritaConf.Password,
		Gas:      conf.Conf.IritaConf.GasLimit,
		Memo:     "",
		Mode:     types.Commit,
	}
	req := nft.MintNFTRequest{
		Denom: denomID,
		ID:    ID,
		Name:  name,
		URI:   uri,
		Data:  data,
	}
	tx, err := c.c.NFT.MintNFT(req, baseTx)
	if err != nil {
		return nil, fmt.Errorf("mint nft error, %v", err)
	}
	return &tx, nil
}

// TransferNFT 转移NFT
func (c *Client) TransferNFT(
	denomID, ID, recipient string,
) (*types.ResultTx, error) {
	baseTx := types.BaseTx{
		From:     conf.Conf.IritaConf.Username,
		Password: conf.Conf.IritaConf.Password,
		Gas:      conf.Conf.IritaConf.GasLimit,
		Memo:     "",
		Mode:     types.Commit,
	}
	req := nft.TransferNFTRequest{
		Denom:     denomID,
		ID:        ID,
		Recipient: recipient,
	}
	tx, err := c.c.NFT.TransferNFT(req, baseTx)
	if err != nil {
		return nil, fmt.Errorf("transfer nft error, %v", err)
	}
	return &tx, nil
}

// QueryAccount 查询账户信息
func (c *Client) QueryAccount(name, password string) (string, error) {
	_, address, err := c.c.BaseClient.Find(name, password)
	if err != nil {
		return "", fmt.Errorf("find error, %v", err)
	}
	return address.String(), nil
}

// QueryDenom 查询NFT类别信息
func (c *Client) QueryDenom(ID string) (*nft.QueryDenomResp, error) {
	resp, err := c.c.NFT.QueryDenom(ID)
	if err != nil {
		return nil, fmt.Errorf("query denom error, %v", err)
	}
	return &resp, nil
}

// QueryNFT 查询NFT信息
func (c *Client) QueryNFT(denomID, ID string) (*nft.QueryNFTResp, error) {
	resp, err := c.c.NFT.QueryNFT(denomID, ID)
	if err != nil {
		return nil, fmt.Errorf("query nft error, %v", err)
	}
	return &resp, nil
}

// QueryOwner 查询账户拥有的NFT
func (c *Client) QueryOwner(creator, denomID string) (*nft.QueryOwnerResp, error) {
	resp, err := c.c.NFT.QueryOwner(creator, denomID, nil)
	if err != nil {
		return nil, fmt.Errorf("query owner error, %v", err)
	}
	return &resp, nil
}

// QueryNFTs 查询NFT列表
func (c *Client) QueryNFTs(denomID string) (*nft.QueryCollectionResp, error) {
	resp, err := c.c.NFT.QueryCollection(denomID, nil)
	if err != nil {
		return nil, fmt.Errorf("query nfts error, %v", err)
	}
	return &resp, nil
}

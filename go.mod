module irita-api

require (
	github.com/bianjieai/irita-sdk-go v1.1.1-0.20211214032850-7c9cd100e6bd
	github.com/bianjieai/opb-sdk-go v0.1.0
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/gin-gonic/gin v1.7.4
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.3
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.14

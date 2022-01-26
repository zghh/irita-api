package conf

// Conf 配置对象
var Conf *Config

// Config 配置信息
type Config struct {
	ServerConf ServerConfig `mapstructure:"server"`
	LoggerConf LoggerConfig `mapstructure:"logger"`
	IritaConf  IritaConfig  `mapstructure:"irita"`
}

// ServerConfig 服务端配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

// IritaConfig itita配置
type IritaConfig struct {
	Username            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	Mnemonic            string `mapstructure:"mnemonic"`
	Coin                string `mapstructure:"coin"`
	GasLimit            uint64 `mapstructure:"gasLimit"`
	TLSEnable           bool   `mapstructure:"tlsEnable"`
	RPCAddress          string `mapstructure:"rpcAddress"`
	WSAddress           string `mapstructure:"wsAddress"`
	GRPCAddress         string `mapstructure:"grpcAddress"`
	ChainID             string `mapstructure:"chainId"`
	ProjectID           string `mapstructure:"projectId"`
	ProjectKey          string `mapstructure:"projectKey"`
	ChainAccountAddress string `mapstructure:"chainAccountAddress"`
}

const (
	// ConfigPrefix 配置前缀
	ConfigPrefix = "IRITA"
	// DefaultConfigFile 默认配置文件路径
	DefaultConfigFile = "config/irita_config.yaml"
)

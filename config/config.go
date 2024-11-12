package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Ethereum EthereumConfig `mapstructure:"ethereum"`
	VRF      VRFConfig      `mapstructure:"vrf"`
	Log      LogConfig      `mapstructure:"log"`
}

type EthereumConfig struct {
	RPCURL          string `mapstructure:"rpc_url"`
	ContractAddress string `mapstructure:"contract_address"`
	PrivateKey      string `mapstructure:"private_key"`
}

type VRFConfig struct {
	NumWords      uint32        `mapstructure:"num_words"`
	GasLimit      uint32        `mapstructure:"gas_limit"`
	Confirmations uint16        `mapstructure:"confirmations"`
	Timeout       time.Duration `mapstructure:"timeout"`
	PollInterval  time.Duration `mapstructure:"poll_interval"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

func LoadConfig(path string) (*Config, error) {
	// 设置配置文件路径
	viper.SetConfigFile(path)
	// 启用环境变量支持
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	// 从环境变量获取私钥
	if privateKey := os.Getenv("PRIVATE_KEY"); privateKey != "" {
		config.Ethereum.PrivateKey = privateKey
	}

	return &config, nil
}

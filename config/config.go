package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Networks NetworksConfig `mapstructure:"networks"`
	VRF      VRFConfig      `mapstructure:"vrf"`
	Log      LogConfig      `mapstructure:"log"`
}

type NetworksConfig struct {
	Primary  NetworkConfig `mapstructure:"primary"`
	Fallback NetworkConfig `mapstructure:"fallback"`
}

type NetworkConfig struct {
	Name            string `mapstructure:"name"`
	RPCURL          string `mapstructure:"rpc_url"`
	ContractAddress string `mapstructure:"contract_address"`
}

type VRFConfig struct {
	NumWords             uint32            `mapstructure:"num_words"`
	GasLimit             uint32            `mapstructure:"gas_limit"`
	Confirmations        uint16            `mapstructure:"confirmations"`
	Timeout              time.Duration     `mapstructure:"timeout"`
	PollInterval         time.Duration     `mapstructure:"poll_interval"`
	NetworkSwitchTimeout time.Duration     `mapstructure:"network_switch_timeout"`
	Retry                RetryConfig       `mapstructure:"retry"`
	HealthCheck          HealthCheckConfig `mapstructure:"health_check"`
}

type RetryConfig struct {
	MaxRetries int           `mapstructure:"max_retries"`
	Interval   time.Duration `mapstructure:"interval"`
}

type HealthCheckConfig struct {
	Interval time.Duration `mapstructure:"interval"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Output   string `mapstructure:"output"`
	FilePath string `mapstructure:"file_path"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)

	// 启用环境变量替换
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 处理环境变量替换
	for _, k := range viper.AllKeys() {
		val := viper.GetString(k)
		if strings.HasPrefix(val, "${") && strings.HasSuffix(val, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(val, "${"), "}")
			if envVal := os.Getenv(envVar); envVal != "" {
				viper.Set(k, envVal)
			}
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &config, nil
}

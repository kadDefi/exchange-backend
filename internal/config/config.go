package config

import (
	"github.com/spf13/viper"
	"strings"
)

type apiConfig struct {
	Addr string `mapstructure:"addr"`
}

type redisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Contract struct {
	SkipInit             bool   `mapstructure:"skip_init"`
	Address              string `mapstructure:"address"`
	Schema               string `mapstructure:"schema"`
	CreatedAtBlockNumber uint64 `mapstructure:"created_at_block_number"`
}

type Config struct {
	Database struct {
		DriverName     string `mapstructure:"driver_name"`
		DataSourceName string `mapstructure:"data_source_name"`
		AutoMigrate    bool   `mapstructure:"auto_migrate"`
		MaxIdleConns   int    `mapstructure:"max_idle_conns"`
		MaxOpenConns   int    `mapstructure:"max_open_conns"`
	} `mapstructure:"database"`

	Redis struct {
		Cache     redisConfig `mapstructure:"cache"`
		Job       redisConfig `mapstructure:"job"`
		Snowflake redisConfig `mapstructure:"snowflake"`
		Pubsub    redisConfig `mapstructure:"pubsub"`
		Locker    redisConfig `mapstructure:"locker"`
	} `mapstructure:"redis"`

	Ethereum struct {
		ChainID                 int64  `mapstructure:"chain_id"`
		RPCURL                  string `mapstructure:"rpc_url"`
		DefaultStartBlockNumber uint64 `mapstructure:"default_start_block_number"`
		WaitBlockNumber         uint64 `mapstructure:"wait_block_number"`
		MultiCallAddress        string `mapstructure:"multi_call_address"`
	} `mapstructure:"ethereum"`

	Etherscan struct {
		APIURL string `mapstructure:"api_url"`
		APIKey string `mapstructure:"api_key"`
	} `mapstructure:"etherscan"`

	Contract struct {
		Exchange Contract `mapstructure:"exchange"`
	} `mapstructure:"contract"`

	API apiConfig `mapstructure:"api"`
}

var (
	Global *Config
)

func LoadLocal(p string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(p)
	if err := v.MergeInConfig(); err != nil {
		return nil, err
	}

	v.SetEnvPrefix("EXCHANGE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	Global = &cfg

	return &cfg, nil
}

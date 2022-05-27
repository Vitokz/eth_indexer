package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	setDefaults(viper.GetViper())
}

const (
	configPath = "./config"
	configType = "env"
	configName = "config.env"
)

type ConfigI interface {
	GetTendermintRPC() string
	GetTendermintWebsocket() string
	GetEthereumJsonRPC() string
	GetEthereumWebsocket() string
	GetGrpcAddress() string
	GetCosmosRestAddress() string
	GetPgDbDsn() string
}

//TODO: add two config structures 1)env 2)yaml

// Config is configuration struct
type Config struct {
	// addresses
	TendermintRPC       string `mapstructure:"TENDERMINT_RPC"`
	TendermintWebsocket string `mapstructure:"TENDERMINT_WEBSOCKET"`
	EthereumJsonRPC     string `mapstructure:"ETHEREUM_JSON_RPC"`
	EthereumWebsocket   string `mapstructure:"ETHEREUM_WEBSOCKET"`
	GrpcAddress         string `mapstructure:"INDEXER_GRPC_ADDRESS"`
	CosmosRestAddress   string `mapstructre:"COSMOS_REST_ADDRESS"`
	PgDbDsn             string `mapstructure:"PG_DB_DSN"`
}

// ParseConfigInConfigFile read configs from other local config file
func ParseConfigInConfigFile() ConfigI {
	var cfg = new(Config)

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return cfg
}

// ParseConfigInOsEnv read configs from local env variables
func ParseConfigInOsEnv() ConfigI {
	var cfg = new(Config)

	viper.AutomaticEnv()
	// TODO: set env prefix

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("TENDERMINT_RPC", "http://127.0.0.1:26657")
	v.SetDefault("TENDERMINT_WEBSOCKET", "http://127.0.0.1:26657/websocket")
	v.SetDefault("ETHEREUM_JSON_RPC", "http://127.0.0.1:8545")
	v.SetDefault("ETHEREUM_WEBSOCKET", "http://127.0.0.1:8545")
	v.SetDefault("INDEXER_GRPC_ADDRESS", "127.0.0.1:9000")
	v.SetDefault("COSMOS_REST_ADDRESS", "127.0.0.1:1317")
	v.SetDefault("PG_DB_DSN", "host=127.0.0.1 port=5434 user=postgres password=postgres dbname=indexer sslmode=disable")
}

func (c *Config) GetTendermintRPC() string       { return c.TendermintRPC }
func (c *Config) GetTendermintWebsocket() string { return c.TendermintWebsocket }
func (c *Config) GetEthereumJsonRPC() string     { return c.EthereumJsonRPC }
func (c *Config) GetEthereumWebsocket() string   { return c.EthereumWebsocket }
func (c *Config) GetGrpcAddress() string         { return c.GrpcAddress }
func (c *Config) GetCosmosRestAddress() string   { return c.CosmosRestAddress }
func (c *Config) GetPgDbDsn() string             { return c.PgDbDsn }

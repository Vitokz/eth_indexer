package multiclient

import (
	"github.com/Vitokz/eth_indexer/config"
	"github.com/Vitokz/eth_indexer/internal/clients/grpc"
	"github.com/Vitokz/eth_indexer/internal/clients/jsonrpc"
	"github.com/Vitokz/eth_indexer/internal/clients/rest"
	"github.com/cosmos/cosmos-sdk/simapp/params"
)

type MultiClient struct {
	Grpc    grpc.Clients
	Rest    rest.Clients
	Jsonrpc jsonrpc.Clients
}

func NewMultiClient(cfg config.ConfigI, cdc *params.EncodingConfig) (MultiClient, error) {
	var mc MultiClient

	//rest clients
	cosmosRestClient := rest.NewCosmosClient(cfg.GetCosmosRestAddress())
	//grpc clients
	cosmosGrpcClient, err := grpc.NewClient(cfg.GetGrpcAddress())
	if err != nil {
		return mc, err
	}

	//TODO: tendermint websocket
	//TODO: eth websocket

	return MultiClient{
		Grpc: cosmosGrpcClient,
		Rest: rest.NewRestClients(cosmosRestClient),
	}, nil
}

//type WebSocketClients struct {
//	clients map[string]
//}

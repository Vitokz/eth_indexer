package rest

type CosmosClient interface {
}

type cosmosClient client

func NewCosmosClient(addr string) CosmosClient {
	clt := NewClient(addr)
	return &clt
}

//TODO: Add requestsz

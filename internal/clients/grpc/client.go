package grpc

import (
	sdkAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkDistributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"google.golang.org/grpc"
)

//TODO: Make all query clients interfaces with create new message and push request + unmarshal response

type Clients interface {
	StopAllCons()

	AuthClient() sdkAuthTypes.QueryClient
	DistributionClient() sdkDistributionTypes.QueryClient
}

type clients struct {
	authClient         sdkAuthTypes.QueryClient
	distributionClient sdkDistributionTypes.QueryClient

	cons []*grpc.ClientConn
}

func NewClient(addr string) (*clients, error) {
	var (
		err error
	)

	grpc.WithInsecure()

	aClient, aConn, err := newAuthClient(addr)
	if err != nil {
		return nil, err
	}

	dstClient, dstConn, err := newDistributionClient(addr)
	if err != nil {
		return nil, err
	}

	clt := clients{
		authClient:         aClient,
		distributionClient: dstClient,
	}

	clt.cons = append(clt.cons,
		aConn,
		dstConn,
	)

	return &clt, nil
}

func newAuthClient(addr string) (sdkAuthTypes.QueryClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := sdkAuthTypes.NewQueryClient(conn)

	return client, nil, nil
}

func newDistributionClient(addr string) (sdkDistributionTypes.QueryClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := sdkDistributionTypes.NewQueryClient(conn)

	return client, nil, nil
}

func (c *clients) StopAllCons() {
	for _, v := range c.cons {
		v.Close()
	}
}

func (c *clients) AuthClient() sdkAuthTypes.QueryClient                 { return c.authClient }
func (c *clients) DistributionClient() sdkDistributionTypes.QueryClient { return c.distributionClient }

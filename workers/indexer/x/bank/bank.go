package bank

import (
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
)

var (
	msgSend      = sdkTypes.MsgTypeURL(&bankTypes.MsgSend{})
	msgMultiSend = sdkTypes.MsgTypeURL(&bankTypes.MsgMultiSend{})
)

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (s *Bank) GetHandler() types.HandlerI {
	return func(msg sdkTypes.Msg) error {
		switch msg.(type) {
		case *bankTypes.MsgSend:
			return s.msgSend()
		case *bankTypes.MsgMultiSend:
			return s.msgMultiSend()
		default:
			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
		}
	}
}

func (s *Bank) msgSend() error {

	return nil
}

func (s *Bank) msgMultiSend() error {

	return nil
}

func (s *Bank) Msgs() []string {
	return []string{
		msgSend,
		msgMultiSend,
	}
}

func (s *Bank) GetName() string {
	return bankTypes.ModuleName
}

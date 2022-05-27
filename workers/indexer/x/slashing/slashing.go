package slashing

import (
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/pkg/errors"
)

var (
	msgUnjail = sdkTypes.MsgTypeURL(&slashingTypes.MsgUnjail{})
)

type Slashing struct {
}

func NewSlashing() *Slashing {
	return &Slashing{}
}

func (s *Slashing) GetHandler() types.HandlerI {
	return func(msg sdkTypes.Msg) error {
		switch msg.(type) {
		case *slashingTypes.MsgUnjail:
			return s.msgUnjail()
		default:
			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
		}
	}
}

func (s *Slashing) msgUnjail() error {

	return nil
}
func (s *Slashing) Msgs() []string {
	return []string{
		msgUnjail,
	}
}

func (s *Slashing) GetName() string {
	return slashingTypes.ModuleName
}

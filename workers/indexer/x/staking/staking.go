package staking

import (
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
)

var (
	msgDelegate        = sdkTypes.MsgTypeURL(&stakingTypes.MsgDelegate{})
	msgEditValidator   = sdkTypes.MsgTypeURL(&stakingTypes.MsgEditValidator{})
	msgBeginRedelegate = sdkTypes.MsgTypeURL(&stakingTypes.MsgBeginRedelegate{})
	msgUndelegate      = sdkTypes.MsgTypeURL(&stakingTypes.MsgUndelegate{})
	msgCreateValidator = sdkTypes.MsgTypeURL(&stakingTypes.MsgCreateValidator{})
)

type Staking struct {
}

func NewStaking() types.Module {
	return &Staking{}
}

func (s *Staking) GetHandler() types.HandlerI {
	return func(msg sdkTypes.Msg) error {
		switch msg.(type) {
		case *stakingTypes.MsgCreateValidator:
			return s.msgCreateValidator()
		case *stakingTypes.MsgEditValidator:
			return s.msgEditValidator()
		case *stakingTypes.MsgDelegate:
			return s.msgDelegate()
		case *stakingTypes.MsgBeginRedelegate:
			return s.msgRedelegate()
		case *stakingTypes.MsgUndelegate:
			return s.msgUndelegate()
		default:
			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
		}
	}
}

func (s *Staking) msgDelegate() error {

	return nil
}

func (s *Staking) msgUndelegate() error {

	return nil
}

func (s *Staking) msgEditValidator() error {

	return nil
}

func (s *Staking) msgCreateValidator() error {

	return nil
}

func (s *Staking) msgRedelegate() error {

	return nil
}

func (s *Staking) Msgs() []string {
	return []string{
		msgDelegate,
		msgEditValidator,
		msgBeginRedelegate,
		msgUndelegate,
		msgCreateValidator,
	}
}

func (s *Staking) GetName() string {
	return stakingTypes.ModuleName
}

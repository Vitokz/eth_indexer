package gov

import (
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/pkg/errors"
)

var (
	msgVote           = sdkTypes.MsgTypeURL(&govTypes.MsgVote{})
	msgDeposit        = sdkTypes.MsgTypeURL(&govTypes.MsgDeposit{})
	msgVoteWeighted   = sdkTypes.MsgTypeURL(&govTypes.MsgVoteWeighted{})
	msgSubmitProposal = sdkTypes.MsgTypeURL(&govTypes.MsgSubmitProposal{})
)

type Gov struct {
}

func NewGov() *Gov {
	return &Gov{}
}

func (s *Gov) GetHandler() types.HandlerI {
	return func(msg sdkTypes.Msg) error {
		switch msg.(type) {
		case *govTypes.MsgDeposit:
			return s.msgDeopist()
		case *govTypes.MsgSubmitProposal:
			return s.msgSubmitProposal()
		case *govTypes.MsgVote:
			return s.msgVote()
		case *govTypes.MsgVoteWeighted:
			return s.msgVoteWeighted()
		default:
			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
		}
	}
}

func (s *Gov) msgVote() error {

	return nil
}

func (s *Gov) msgDeopist() error {

	return nil
}

func (s *Gov) msgVoteWeighted() error {

	return nil
}

func (s *Gov) msgSubmitProposal() error {

	return nil
}

func (s *Gov) Msgs() []string {
	return []string{
		msgVote,
		msgDeposit,
		msgSubmitProposal,
		msgVoteWeighted,
	}
}

func (s *Gov) GetName() string {
	return govTypes.ModuleName
}

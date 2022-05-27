package types

import (
	"github.com/Vitokz/eth_indexer/workers/indexer/x/bank"
	"github.com/Vitokz/eth_indexer/workers/indexer/x/gov"
	"github.com/Vitokz/eth_indexer/workers/indexer/x/slashing"
	"github.com/Vitokz/eth_indexer/workers/indexer/x/staking"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Module interface {
	GetHandler() HandlerI
	Msgs() []string
	GetName() string
}

type HandlerI func(msg sdkTypes.Msg) error

type ModuleManager interface {
	GetMsgsTypesAndThemHandlers() MsgsTypes
}

type MsgsTypes interface {
	GetMsgHandler(msgType string) HandlerI
}

type moduleManager struct {
	Modules []Module

	modulesMsgs map[string][]string

	handlers map[string]HandlerI
}

func NewModuleManager() ModuleManager {
	var mm moduleManager

	mm.Modules = []Module{
		staking.NewStaking(),
		slashing.NewSlashing(),
		gov.NewGov(),
		bank.NewBank(),
	}
	mm.modulesMsgs = make(map[string][]string)
	mm.handlers = make(map[string]HandlerI)

	mm.setMsgsAndThemHandlers()
	return &mm
}

func (m *moduleManager) setMsgsAndThemHandlers() {
	for _, v := range m.Modules {
		m.handlers[v.GetName()] = v.GetHandler()
		m.modulesMsgs[v.GetName()] = v.Msgs()
	}
}

func (m *moduleManager) GetMsgsTypesAndThemHandlers() MsgsTypes {
	result := make(msgsTypes)
	for module, msgs := range m.modulesMsgs {
		for _, msg := range msgs {
			result[msg] = m.handlers[module]
		}
	}

	return &result
}

type msgsTypes map[string]HandlerI

func (t msgsTypes) GetMsgHandler(msgType string) HandlerI { return t[msgType] }

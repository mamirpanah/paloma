package eventbus

import (
	"context"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/v2/util/liblog"
)

var (
	skywayBatchBuilt  = newEvent[SkywayBatchBuiltEvent]()
	evmActivatedChain = newEvent[EVMActivatedChainEvent]()
)

type (
	EventHandler[E any] func(context.Context, E) error
	Event[E any]        struct {
		subscribers map[string]EventHandler[E]
	}
)

func newEvent[E any]() Event[E] {
	return Event[E]{
		subscribers: make(map[string]EventHandler[E]),
	}
}

func (e Event[E]) Publish(ctx context.Context, event E) {
	keys := make([]string, 0, len(e.subscribers))
	for k := range e.subscribers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if e.subscribers[k] != nil {
			logger := liblog.FromSDKLogger(sdk.UnwrapSDKContext(ctx).Logger()).
				WithComponent("eventbus").
				WithFields("event", event).
				WithFields("subscriber", k)
			logger.Debug("Handling event")
			if err := e.subscribers[k](ctx, event); err != nil {
				logger.WithError(err).Error("Failed to handle event")
			}
		}
	}
}

func (e Event[E]) Subscribe(id string, fn EventHandler[E]) {
	e.subscribers[id] = fn
}

func (e Event[E]) Unsubscribe(id string) {
	e.subscribers[id] = nil
}

type SkywayBatchBuiltEvent struct {
	ChainReferenceID string
}

func SkywayBatchBuilt() *Event[SkywayBatchBuiltEvent] {
	return &skywayBatchBuilt
}

type EVMActivatedChainEvent struct {
	ChainReferenceID      string
	SmartContractUniqueID []byte
}

func EVMActivatedChain() *Event[EVMActivatedChainEvent] {
	return &evmActivatedChain
}

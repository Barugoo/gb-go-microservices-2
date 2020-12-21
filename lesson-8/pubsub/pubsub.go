package pubsub

import (
	"time"

	"github.com/mediocregopher/radix/v3"
)

const timeout = time.Second

type (
	PubSub interface {
		Publish(string) error
		Subscribe(chan<- radix.PubSubMessage) error
		Unsubscribe(chan<- radix.PubSubMessage) error
	}

	RedisPubsub struct {
		name   string
		pool   *radix.Pool
		pubsub radix.PubSubConn
	}
)

var (
	connFunc = func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(timeout),
		)
	}
)

func NewPubSub(name, address string) PubSub {
	pool, err := radix.NewPool("tcp", address, 1, radix.PoolConnFunc(connFunc))
	if err != nil {
		panic(err)
	}
	pubsub, err := radix.PersistentPubSubWithOpts("tcp", address)
	if err != nil {
		panic(err)
	}
	return &RedisPubsub{
		name:   name,
		pool:   pool,
		pubsub: pubsub,
	}
}

func (ps *RedisPubsub) Publish(v string) error {
	return ps.pool.Do(radix.Cmd(nil, "PUBLISH", ps.name, v))
}

func (ps *RedisPubsub) Subscribe(ch chan<- radix.PubSubMessage) error {
	return ps.pubsub.Subscribe(ch, ps.name)
}

func (ps *RedisPubsub) Unsubscribe(ch chan<- radix.PubSubMessage) error {
	return ps.pubsub.Unsubscribe(ch, ps.name)
}

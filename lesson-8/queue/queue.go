package queue

import (
	"time"

	"github.com/mediocregopher/radix/v3"
)

const timeout = time.Second

type (
	Queue interface {
		Pop() (string, error)
		Push(string) error
	}

	RedisQueue struct {
		name string
		pool *radix.Pool
	}
)

var (
	connFunc = func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(timeout),
		)
	}
)

func NewQueue(name, address string) Queue {
	pool, err := radix.NewPool("tcp", address, 1, radix.PoolConnFunc(connFunc))
	if err != nil {
		panic(err)
	}
	return &RedisQueue{
		name: name,
		pool: pool,
	}
}

func (q *RedisQueue) Pop() (string, error) {
	var result string
	err := q.pool.Do(radix.Cmd(&result, "LPOP", q.name))
	return result, err
}

func (q *RedisQueue) Push(v string) error {
	return q.pool.Do(radix.Cmd(nil, "RPUSH", q.name, v))
}

package pubsub

import (
	"strings"
	"sync"
	"time"

	"github.com/mediocregopher/radix/v3"
	. "gopkg.in/check.v1"

	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type PubSubSuite struct{}

var _ = Suite(&PubSubSuite{})

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func (s *PubSubSuite) TestPubSub(c *C) {
	var (
		ch1 = make(chan radix.PubSubMessage)
		ch2 = make(chan radix.PubSubMessage)

		err error
		wg  = &sync.WaitGroup{}
	)

	p := NewPubSub("topic", "127.0.0.1:6379")
	s1 := NewPubSub("topic", "127.0.0.1:6379")
	s2 := NewPubSub("topic", "127.0.0.1:6379")

	err = s1.Subscribe(ch1)
	c.Assert(err, IsNil)
	err = s2.Subscribe(ch2)
	c.Assert(err, IsNil)

	var (
		sb1 = strings.Builder{}
		sb2 = strings.Builder{}
	)
	wg.Add(2)
	go func() {
		for v := range ch1 {
			sb1.Write(v.Message)
		}
		wg.Done()
	}()
	go func() {
		for v := range ch2 {
			sb2.Write(v.Message)
		}
		wg.Done()
	}()

	for _, v := range alphabet {
		err := p.Publish(string(v))
		c.Assert(err, IsNil)
	}

	time.Sleep(time.Second)
	err = s1.Unsubscribe(ch1)
	c.Assert(err, IsNil)
	close(ch1)
	err = s2.Unsubscribe(ch2)
	c.Assert(err, IsNil)
	close(ch2)

	wg.Wait()
	c.Assert(sb1.String(), Equals, alphabet)
	c.Assert(sb2.String(), Equals, alphabet)
}

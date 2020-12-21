package queue

import (
	. "gopkg.in/check.v1"

	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type QueueSuite struct{}

var _ = Suite(&QueueSuite{})

func (s *QueueSuite) TestQueue(c *C) {
	ss := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	q := NewQueue("test_queeu", "127.0.0.1:6379")
	for _, s := range ss {
		err := q.Push(s)
		c.Assert(err, IsNil)
	}

	for i := 0; i < len(ss); i += 1 {
		s, err := q.Pop()
		c.Assert(err, IsNil)
		c.Assert(s, Equals, ss[i])
	}
}

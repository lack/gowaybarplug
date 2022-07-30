package gowaybarplug

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ChBuf struct {
	Ch   chan bool
	Buff bytes.Buffer
}

func NewChBuf() *ChBuf {
	return &ChBuf{
		Ch: make(chan bool, 100),
	}
}

func (c *ChBuf) Write(data []byte) (int, error) {
	n, err := c.Buff.Write(data)
	c.Ch <- err == nil
	return n, err
}

func (c *ChBuf) WaitForBuffer(deadline time.Duration) {
	select {
	case <-c.Ch:
	case <-time.After(deadline):
	}
}

func TestNewUpdater(t *testing.T) {
	buff := NewChBuf()
	u := NewUpdater().OutputTo(buff)
	u.Ch <- &Status{Text: "test"}
	buff.WaitForBuffer(2 * time.Second)
	result, err := buff.Buff.ReadString('\n')
	assert.NoError(t, err)
	assert.Equal(t, result, "{\"text\":\"test\"}\n")
}

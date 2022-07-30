package gowaybarplug

import (
	"fmt"
	"io"
	"os"
)

// Updater runs a goroutine that accepts Status updates in a channel and ptints them to stdout.
type Updater struct {
	// Status is the main Status reporting channel. Every Status submitted here will be sent to stdout.
	Status chan *Status

	last   string
	writer io.Writer
}

// Create a new Updater and start the receiver thread
func NewUpdater() *Updater {
	u := Updater{
		Status: make(chan *Status, 10),
		writer: os.Stdout,
	}
	go u.run()
	return &u
}

func (u *Updater) run() {
	for s := range u.Status {
		next := s.String()
		if next != u.last {
			u.last = next
			fmt.Fprintln(u.writer, next)
		}
	}
}

func (u *Updater) OutputTo(writer io.Writer) *Updater {
	u.writer = writer
	return u
}

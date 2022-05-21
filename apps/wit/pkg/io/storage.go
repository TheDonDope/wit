package io

import (
	"io"
	"wit/apps/wit/pkg/tracker"
)

// Persistor provides functionality to write to I/O and persist stashes.
type Persistor interface {
	// Commit persists the given stashes to given io.Writer.
	Commit(w io.Writer, s ...*tracker.Stash)
}

// Selector provides functionality read from I/O and return a Stash
type Selector interface {
	// Checkout retrieves the stash from the given io.Reader
	Checkout(r io.Reader) *tracker.Stash
}

// Repository holds the information about all the stashes.
type Repository struct {
	// All stashes as mapping of their name to their instance
	r       io.Reader
	w       io.Writer
	Stashes map[string]*tracker.Stash
}

// NewRepository ...
func NewRepository(r io.Reader, w io.Writer) *Repository {
	return &Repository{
		r:       r,
		w:       w,
		Stashes: make(map[string]*tracker.Stash),
	}
}

package io

import (
	"io"
	"wit/apps/wit/pkg/tracker"
)

// Persistor provides a way to save the state of a repository to disk.
type Persistor interface {
	// Persists the given stash to disk.
	Commit(w io.Writer, s ...*tracker.Stash)
}

// Selector retrieves the stash with the given name from disk.
type Selector interface {
	// Retrieve the stash with the given name from disk.
	Pull(r io.Reader) *tracker.Stash
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

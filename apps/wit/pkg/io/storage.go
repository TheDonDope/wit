package io

import (
	"io"
	"wit/apps/wit/pkg/tracker"

	"github.com/spf13/viper"
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

// Initializer provides functionality to create the folder
// structure for local storage.
type Initializer interface {
	// Init creates the folder structore to store the files
	// of the given Repository. The repositories configuration
	// (e.g. $WIT_DIR) will be honored.
	Init(r *Repository)
}

// Puller provides functionality to update the in-memory
// state of the repository with the persistant state
// from I/O.
type Puller interface {
	// Pull gets the latest persistant state from I/O
	// and updates the in-memory state accordingly.
	Pull(r *Repository)
}

// Pusher provides functionality to persist the current
// in-memory state of the repository to I/O.
type Pusher interface {
	// Push persists the current repository state to I/O
	Push(r *Repository)
}

// Repository holds the information about all the stashes.
type Repository struct {
	// All stashes as mapping of their name to their instance
	Persistor Persistor
	Selector  Selector
	WitDir    string
	Stashes   map[string]*tracker.Stash
}

// RepoOption ...
type RepoOption func(r *Repository) RepoOption

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (r *Repository) Option(opts ...RepoOption) (previous RepoOption) {
	for _, opt := range opts {
		previous = opt(r)
	}
	return previous
}

// NewRepository ...
func NewRepository(p Persistor, s Selector) *Repository {
	return &Repository{
		Persistor: p,
		Selector:  s,
		WitDir:    viper.GetString("WitDir"),
		Stashes:   make(map[string]*tracker.Stash),
	}
}

// WitDir ...
func WitDir(p string) RepoOption {
	return func(r *Repository) RepoOption {
		previous := r.WitDir
		r.WitDir = p
		return WitDir(previous)
	}
}

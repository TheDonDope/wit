package io

import (
	"wit/apps/wit/pkg/tracker"
)

// Persistor provides a way to save the state of a repository to disk.
type Persistor interface {
	// Persists the given stash to disk.
	Commit(stash *tracker.Stash)
}

// Selector retrieves the stash with the given name from disk.
type Selector interface {
	// Retrieve the stash with the given name from disk.
	Pull(name string) *tracker.Stash
}

// Repository holds the information about all the stashes.
type Repository struct {
	// All stashes as mapping of their name to their instance
	Stashes map[string]*tracker.Stash
}

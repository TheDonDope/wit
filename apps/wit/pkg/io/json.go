package io

import (
	"encoding/json"
	"fmt"
	"wit/apps/wit/pkg/tracker"
)

// NewJSONPersistor ...
func NewJSONPersistor(d Destination) Persistor {
	return d
}

// NewJSONSelector ...
func NewJSONSelector(s Source) Selector {
	return s
}

// Commit ...
func (d Destination) Commit(stash *tracker.Stash) {
	fmt.Printf("Persisting stash: %v at destination: %v", stash, d)
	b, err := json.Marshal(stash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bytes: %v", b)
}

// Pull ...
func (s Source) Pull(name string) *tracker.Stash {
	fmt.Printf("Pulling stash: %v from source: %v", name, s)
	return &tracker.Stash{}
}

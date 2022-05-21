package io

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"wit/apps/wit/pkg/tracker"
)

// JSONSelector ...
type JSONSelector struct{}

// JSONPersistor ...
type JSONPersistor struct{}

// Checkout reads from the given io.Reader assuming JSON bytes
// and returning an unmarshaled Stash instance.
func (j *JSONSelector) Checkout(r io.Reader) *tracker.Stash {
	fmt.Println(fmt.Printf("JSONSelector.Pull(r: %v)", r))

	var stash *tracker.Stash

	b, err := ioutil.ReadAll(r)
	fmt.Println(fmt.Printf("  ioutil.ReadAll(): %v", b))
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &stash); err != nil {
		panic(err)
	}

	return stash
}

// Commit writes to the given io.Writer
func (j *JSONPersistor) Commit(w io.Writer, s ...*tracker.Stash) {
	fmt.Println(fmt.Printf("JSONPersistor.Commit(w: %v, s: %v)", w, s))
	for _, v := range s {
		b, err := json.Marshal(v)
		fmt.Println(fmt.Printf("  json.Marshal(): %v", b))
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(b); err != nil {
			panic(err)
		}
	}
}

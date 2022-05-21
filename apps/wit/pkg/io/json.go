package io

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"wit/apps/wit/pkg/tracker"
)

// JSONSelector acts as a type holder to implement the Selector interface.
type JSONSelector struct{}

// JSONPersistor acts as a type holder to implement the Persitor interface.
type JSONPersistor struct{}

// JSONRepository acts as a type holder to implement the Puller and Pusher interfaces.
type JSONRepository struct{}

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

// Pull checks for all existing (in-memory) stashes their corresponding
// latest persistant state and updates accordingly.
func (j *JSONRepository) Pull(r *Repository) {
	fmt.Println(fmt.Printf("JSONRepository.Pull(r: %v)", r))

	for _, v := range r.Stashes {
		path := r.HomePath + "/" + v.Strain + ".json"
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		stash := r.Selector.Checkout(f)
		r.Stashes[v.Strain] = stash
	}
}

// Push persists all existing (in-memory) stashes to their corresponding
// json files.
func (j *JSONRepository) Push(r *Repository) {
	fmt.Println(fmt.Printf("JSONRepository.Push(r: %v)", r))
	for _, v := range r.Stashes {
		path := r.HomePath + "/" + v.Strain + ".json"
		if _, err := os.Stat(path); err == nil {
			os.Remove(path)
		}
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r.Persistor.Commit(f)
	}
}

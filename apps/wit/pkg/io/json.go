package io

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"wit/apps/wit/pkg/tracker"
)

type jsonSelector struct {
	r io.Reader
}

type jsonPersistor struct {
	w io.Writer
}

// NewJSONSelector ...
func NewJSONSelector(r io.Reader) Selector {
	return &jsonSelector{r}
}

// NewJSONPersistor ...
func NewJSONPersistor(w io.Writer) Persistor {
	return &jsonPersistor{w}
}

// Pull ...
func (j *jsonSelector) Pull(r io.Reader) *tracker.Stash {
	fmt.Println(fmt.Printf("jsonSelector.Pull(r: %v)", r))

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

// Commit ...
func (j *jsonPersistor) Commit(w io.Writer, s ...*tracker.Stash) {
	fmt.Println(fmt.Printf("jsonPersistor.Commit(w: %v, s: %v)", w, s))
	for _, v := range s {
		path := v.Strain + ".json"
		b, err := json.Marshal(v)
		fmt.Println(fmt.Printf("  json.Marshal(): %v", b))
		if err != nil {
			panic(err)
		}

		//
		if err := ioutil.WriteFile(path, b, 0644); err != nil {
			panic(err)
		}
	}
}

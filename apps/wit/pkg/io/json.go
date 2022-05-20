package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wit/apps/wit/pkg/tracker"
)

// Commit ...
func (d Destination) Commit(stash *tracker.Stash) {
	fmt.Println(fmt.Printf("Destination.Commit(d: %v, stash: %v)", d, stash))

	b, err := json.Marshal(stash)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("  b: %v", b))

	if _, err := os.Stat(string(d)); err == nil {
		os.Remove(string(d))
	}

	f, err := os.OpenFile(string(d), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(fmt.Printf("  f: %v", f))

	f.Write(b)
}

// Pull ...
func (s Source) Pull(name string) *tracker.Stash {
	fmt.Println(fmt.Printf("Source.Pull(s: %v, name: %v)", s, name))

	f, err := os.Open(string(s))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(fmt.Printf("  f: %v", f))

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("  b: %v", b))

	var stashes []*tracker.Stash
	json.Unmarshal(b, &stashes)
	fmt.Println(fmt.Printf("  stashes:  %v", stashes))

	for _, s := range stashes {
		fmt.Println(fmt.Printf("  s: %v", s))
		if s.Strain == name {
			fmt.Println("found")
			return s
		}
	}

	fmt.Println("not found")
	return &tracker.Stash{}
}

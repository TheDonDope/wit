package io

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"testing"
	"wit/apps/wit/pkg/tracker"
)

func TestJSONSelector_Checkout(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		j    *JSONSelector
		args args
		want *tracker.Stash
	}{
		{"In-Memory JSON String", &JSONSelector{}, args{r: strings.NewReader(defaultStashJSON())}, defaultStash()},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONSelector{}
			if got := j.Checkout(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONSelector.Checkout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONPersistor_Commit(t *testing.T) {
	stashes := []*tracker.Stash{
		{Strain: "hindu-kush", Amount: 9.0},
		{Strain: "gorilla-glue-#4", Amount: 12.0},
	}
	stashesJSON := stashJSON(stashes[0]) + stashJSON(stashes[1])

	type args struct {
		s []*tracker.Stash
	}
	tests := []struct {
		name  string
		j     *JSONPersistor
		args  args
		wantW string
	}{
		{"Two stashes", &JSONPersistor{}, args{s: stashes}, string(stashesJSON)},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONPersistor{}
			w := &bytes.Buffer{}
			j.Commit(w, tt.args.s...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("JSONPersistor.Commit() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestJSONRepository_Pull(t *testing.T) {
	type args struct {
		r *Repository
	}
	tests := []struct {
		name string
		j    *JSONRepository
		args args
	}{
		{"Zero stashes", &JSONRepository{}, args{r: NewRepository(&JSONPersistor{}, &JSONSelector{})}},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRepository{}
			j.Pull(tt.args.r)
		})
	}
}

func TestJSONRepository_Push(t *testing.T) {
	type args struct {
		r *Repository
	}
	tests := []struct {
		name string
		j    *JSONRepository
		args args
	}{
		{"Zero stashes", &JSONRepository{}, args{r: NewRepository(&JSONPersistor{}, &JSONSelector{})}},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRepository{}
			j.Push(tt.args.r)
		})
	}
}

func stash(strain string, amount float64) *tracker.Stash {
	return &tracker.Stash{Strain: strain, Amount: amount}
}

func defaultStash() *tracker.Stash {
	return stash("hindu-kush", 8.0)
}

func stashJSON(s *tracker.Stash) string {
	json, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(json)
}

func defaultStashJSON() string {
	json, err := json.Marshal(defaultStash())
	if err != nil {
		panic(err)
	}
	return string(json)
}

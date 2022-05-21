package io

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
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
		{
			"From JSON on disk",
			&JSONSelector{},
			args{r: jsonFile("../../testdata/hindu-kush.json")},
			stash(),
		},
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

/* func TestJSONPersistor_Commit(t *testing.T) {
	stashes := []*tracker.Stash{
		{Strain: "Hindu Kush", Amount: 9.0},
		{Strain: "Gorilla Glue #4", Amount: 12.0},
	}
	stashesJSON := jsonString(stashes[0]) + jsonString(stashes[1])

	type args struct {
		s []*tracker.Stash
	}
	tests := []struct {
		name  string
		j     *JSONPersistor
		args  args
		wantW string
	}{
		{
			"Two stashes",
			&JSONPersistor{},
			args{s: stashes},
			string(stashesJSON),
		},
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
} */

func TestJSONRepository_Pull(t *testing.T) {
	type args struct {
		r *Repository
	}
	tests := []struct {
		name string
		j    *JSONRepository
		args args
	}{
		{
			"Zero stashes",
			&JSONRepository{},
			args{
				r: jsonRepository(),
			},
		},
		{
			"One stash",
			&JSONRepository{},
			args{
				r: jsonRepository(stash()),
			},
		},
		{
			"Two stashes",
			&JSONRepository{},
			args{
				r: jsonRepository(stashes()...),
			},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRepository{}
			j.Pull(tt.args.r)
		})
	}
}

/* func TestJSONRepository_Push(t *testing.T) {
	type args struct {
		r *Repository
	}
	tests := []struct {
		name string
		j    *JSONRepository
		args args
	}{
		{
			"Zero stashes",
			&JSONRepository{},
			args{
				r: jsonRepository(),
			},
		},
		{
			"One stash",
			&JSONRepository{},
			args{
				r: jsonRepository(stash()),
			},
		},
		{
			"Two stashes",
			&JSONRepository{},
			args{
				r: jsonRepository(stashes()...),
			},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRepository{}
			j.Push(tt.args.r)
		})
	}
}
*/

func jsonString(s *tracker.Stash) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func jsonFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func jsonRepository(s ...*tracker.Stash) *Repository {
	r := NewRepository(&JSONPersistor{}, &JSONSelector{})
	for _, v := range s {
		r.Stashes[v.Strain] = v
	}
	r.Option(HomePath("../../testdata"))
	return r
}

func stash() *tracker.Stash {
	return &tracker.Stash{Strain: "hindu-kush", Amount: 8.0}
}

func stashes() []*tracker.Stash {
	stashes := []*tracker.Stash{
		{Strain: "hindu-kush", Amount: 8.0},
		{Strain: "gorilla-glue-#4", Amount: 10.0},
	}
	return stashes
}

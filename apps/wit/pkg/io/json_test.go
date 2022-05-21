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
	hc := &tracker.Stash{Strain: "hindu-kush", Amount: 8.0}
	hcJSON, err := json.Marshal(hc)
	if err != nil {
		panic(err)
	}

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		j    *JSONSelector
		args args
		want *tracker.Stash
	}{
		{"In-Memory JSON String", &JSONSelector{}, args{r: strings.NewReader(string(hcJSON))}, hc},
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
	type args struct {
		s []*tracker.Stash
	}
	tests := []struct {
		name  string
		j     *JSONPersistor
		args  args
		wantW string
	}{
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRepository{}
			j.Push(tt.args.r)
		})
	}
}

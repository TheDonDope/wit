package io

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"wit/apps/wit/pkg/tracker"
)

func TestNewJSONSelector(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Selector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONSelector(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJSONPersistor(t *testing.T) {
	tests := []struct {
		name  string
		want  Persistor
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := NewJSONPersistor(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONPersistor() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewJSONPersistor() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_jsonSelector_Pull(t *testing.T) {
	type fields struct {
		r io.Reader
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tracker.Stash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonSelector{
				r: tt.fields.r,
			}
			if got := j.Pull(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonSelector.Pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonPersistor_Commit(t *testing.T) {
	type fields struct {
		w io.Writer
	}
	type args struct {
		s []*tracker.Stash
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantW  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonPersistor{
				w: tt.fields.w,
			}
			w := &bytes.Buffer{}
			j.Commit(w, tt.args.s...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("jsonPersistor.Commit() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

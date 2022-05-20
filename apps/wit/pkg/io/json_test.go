package io

import (
	"reflect"
	"testing"
	"wit/apps/wit/pkg/tracker"
)

func TestDestination_Commit(t *testing.T) {
	type args struct {
		stash *tracker.Stash
	}
	tests := []struct {
		name string
		d    Destination
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Commit(tt.args.stash)
		})
	}
}

func TestSource_Pull(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		s    Source
		args args
		want *tracker.Stash
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pull(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source.Pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

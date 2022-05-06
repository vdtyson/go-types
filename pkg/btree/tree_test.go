package btree

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name string
		want *Tree[int]
	}{
		{
			name: "Should return expected",
			want: &Tree[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := New[int](); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

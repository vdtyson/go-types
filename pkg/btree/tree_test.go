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

func TestFromSlice(t *testing.T) {
	type args struct {
		items []any
	}
	tests := []struct {
		name string
		args args
		want *Tree[any]
	}{
		{
			name: "When items is empty, should return tree with nil root",
			args: args{items: nil},
			want: &Tree[any]{},
		},
		{
			name: "When items are non empty, should return expected tree",
			args: args{items: []any{1, 2, 3, 4, 5, 6}},
			want: func() *Tree[any] {
				root := &Node[any]{
					Value: 1,
					Left: &Node[any]{
						Value: 2,
						Left:  &Node[any]{Value: 4},
						Right: &Node[any]{Value: 5},
					},
					Right: &Node[any]{
						Value: 3,
						Left:  &Node[any]{Value: 6},
					},
				}
				return &Tree[any]{Root: root}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := FromSlice(tt.args.items); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromSlice() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

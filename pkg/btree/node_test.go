package btree

import (
	"reflect"
	"testing"
)

func Test_NewNode(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name string
		args args
		want *Node[any]
	}{
		{
			name: "Should return expected",
			args: args{val: 1},
			want: &Node[any]{Value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewNode(tt.args.val); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewNode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

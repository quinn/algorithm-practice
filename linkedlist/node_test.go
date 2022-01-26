package linkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		rootVal string
	}
	tests := []struct {
		name string
		args args
		want *Node[string]
	}{
		{
			name: "test New with string",
			args: args{
				rootVal: "test1",
			},
			want: &Node[string]{
				value: "test1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rootVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

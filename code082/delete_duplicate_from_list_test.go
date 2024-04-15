package code082

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 3, 4, 4, 5}),
			},
			want: createLinkedList([]int{1, 2, 5}),
		},
		{
			name: "test2",
			args: args{
				head: createLinkedList([]int{1, 1, 1, 2, 3}),
			},
			want: createLinkedList([]int{2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

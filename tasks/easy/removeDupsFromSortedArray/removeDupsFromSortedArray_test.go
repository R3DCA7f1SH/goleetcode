package tasks

import "testing"

func Test_removeDupsFromSortedArray(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestCase#1",
			args: args{
				s: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "TestCase#2",
			args: args{
				s: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDupsFromSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

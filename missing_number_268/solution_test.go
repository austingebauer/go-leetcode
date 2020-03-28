package missing_number_268

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "missing number",
			args: args{
				nums: []int{
					3, 0, 1,
				},
			},
			want: 2,
		},
		{
			name: "missing number",
			args: args{
				nums: []int{
					9, 6, 4, 2, 3, 5, 7, 0, 1,
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, missingNumber(tt.args.nums))
		})
	}
}

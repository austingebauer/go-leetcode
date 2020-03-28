package distribute_candies_575

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "distribute candies",
			args: args{
				candies: []int{
					1, 1, 2, 2, 3, 3,
				},
			},
			want: 3,
		},
		{
			name: "distribute candies",
			args: args{
				candies: []int{
					1, 1, 2, 3,
				},
			},
			want: 2,
		},
		{
			name: "distribute candies",
			args: args{
				candies: []int{
					1, 1, 1, 1, 2, 2, 2, 3, 3, 3,
				},
			},
			want: 3,
		},
		{
			name: "distribute candies",
			args: args{
				candies: []int{
					1000, 1000, 2, 1, 2, 5, 3, 1,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, distributeCandies(tt.args.candies))
		})
	}
}

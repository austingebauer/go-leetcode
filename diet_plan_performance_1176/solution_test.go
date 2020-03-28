package diet_plan_performance_1176

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dietPlanPerformance(t *testing.T) {
	type args struct {
		calories []int
		k        int
		lower    int
		upper    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				calories: []int{
					1, 2, 3, 4, 5,
				},
				k:     1,
				lower: 3,
				upper: 3,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				calories: []int{
					3, 2,
				},
				k:     2,
				lower: 0,
				upper: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				calories: []int{
					6, 5, 0, 0,
				},
				k:     2,
				lower: 1,
				upper: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want,
				dietPlanPerformance(tt.args.calories, tt.args.k, tt.args.lower, tt.args.upper))
		})
	}
}

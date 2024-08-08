package find_the_highest_altitude_1732

import "testing"

func Test_largestAltitude(t *testing.T) {
	type args struct {
		gain []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{gain: []int{-5, 1, 5, 0, -7}},
			want: 1,
		},
		{
			name: "2",
			args: args{gain: []int{-4, -3, -2, -1, 4, 3, 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.args.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

package reorder_data_in_log_files_937

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reorder log files",
			args: args{
				logs: []string{
					"dig1 8 1 5 1",
					"let1 art can",
					"dig2 3 6",
					"let2 own kit dig",
					"let3 art zero",
				},
			},
			want: []string{
				"let1 art can",
				"let3 art zero",
				"let2 own kit dig",
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
		},
		{
			name: "reorder log files",
			args: args{
				logs: []string{
					"dig1 8 1 5 1",
					"let1 art can",
					"dig2 3 6",
					"let2 art can",
					"let3 art can",
				},
			},
			want: []string{
				"let1 art can",
				"let2 art can",
				"let3 art can",
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
		},
		{
			name: "reorder log files",
			args: args{
				logs: []string{},
			},
			want: []string{},
		},
		{
			name: "reorder log files",
			args: args{
				logs: []string{
					"let1 art can",
				},
			},
			want: []string{
				"let1 art can",
			},
		},
		{
			name: "reorder log files",
			args: args{
				logs: []string{
					"dig1 8 1 5 1",
				},
			},
			want: []string{
				"dig1 8 1 5 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reorderLogFiles(tt.args.logs))
		})
	}
}

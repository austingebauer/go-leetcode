package excel_sheet_column_number_171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "excel sheet column number",
			args: args{
				s: "A",
			},
			want: 1,
		},
		{
			name: "excel sheet column number",
			args: args{
				s: "AB",
			},
			want: 28,
		},
		{
			name: "excel sheet column number",
			args: args{
				s: "ZY",
			},
			want: 701,
		},
		{
			name: "excel sheet column number",
			args: args{
				s: "YB",
			},
			want: 652,
		},
		{
			name: "excel sheet column number",
			args: args{
				s: "AAA",
			},
			want: 703,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, titleToNumber(tt.args.s))
		})
	}
}

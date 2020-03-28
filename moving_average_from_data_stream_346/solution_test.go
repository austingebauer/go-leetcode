package moving_average_from_data_stream_346

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovingAverageFromDataStream(t *testing.T) {
	m := Constructor(3)
	assert.Equal(t, 1.0, m.Next(1))
	assert.Equal(t, (1+10)/2.0, m.Next(10))
	assert.Equal(t, (1+10+3)/3.0, m.Next(3))
	assert.Equal(t, (10+3+5)/3.0, m.Next(5))
}

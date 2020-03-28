package find_median_from_data_stream_295

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianFinder_0(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	assert.Equal(t, 1.0, mf.FindMedian())
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian())
	mf.AddNum(3)
	assert.Equal(t, 2.0, mf.FindMedian())
}

func TestMedianFinder_1(t *testing.T) {
	mf := Constructor()
	mf.AddNum(-1)
	assert.Equal(t, float64(-1), mf.FindMedian())
	mf.AddNum(-2)
	assert.Equal(t, -1.5, mf.FindMedian())
	mf.AddNum(-3)
	assert.Equal(t, -2.0, mf.FindMedian())
	mf.AddNum(-4)
	assert.Equal(t, -2.5, mf.FindMedian())
	mf.AddNum(-5)
	assert.Equal(t, -3.0, mf.FindMedian())
}

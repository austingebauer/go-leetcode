package moving_average_from_data_stream_346

type MovingAverage struct {
	q    []int
	size int
	sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		q:    make([]int, 0),
		size: size,
		sum:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.q) == this.size {
		dq := this.q[0]
		this.sum -= dq
		this.q = this.q[1:]
		this.q = append(this.q, val)
		this.sum += val
	} else {
		this.sum += val
		this.q = append(this.q, val)
	}

	return float64(this.sum) / float64(len(this.q))
}

package utils

type (
	avgSrc struct {
		Val    float64
		Weight float64
	}

	AvgBuff struct {
		buff []avgSrc
		size uint
		pos  uint
	}
)

func NewAvgBuff(size int) *AvgBuff {
	return &AvgBuff{
		buff: make([]avgSrc, size),
		size: uint(size),
	}
}

func (b *AvgBuff) Push(val, weight float64) {
	i := b.pos % b.size
	b.buff[i] = avgSrc{val, weight}
	b.pos++
}

func (b *AvgBuff) Avg() float64 {
	count := b.pos
	if count > b.size {
		count = b.size
	}

	var (
		val    float64
		weight float64
	)
	for i := uint(0); i < count; i++ {
		val += b.buff[i].Val
		weight += b.buff[i].Weight
	}

	return val / weight
}

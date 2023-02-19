package hyperloglog

import "math"

type HyperLogLog struct {
	p     uint8   //precision
	reg   []uint8 // register :-> m =2^p
	m     uint32  // buckets
	alpha float64 // correction alpha
}

type Hash32 interface {
	Sum32() uint32
}

func NewHyperLogLog(p uint8) *HyperLogLog {
	m := 1 << p
	return &HyperLogLog{
		p:     p,
		reg:   make([]uint8, m),
		m:     uint32(m),
		alpha: getAlpha(p),
	}
}

func getAlpha(p uint8) float64 {
	// calculates the correction factor
	switch p {
	case 4:
		return 0.673
	case 5:
		return 0.697
	case 6:
		return 0.709
	default:
		return 0.7213 / (1.0 + 1.079/math.Pow(2, float64(p)))
	}
}

func (h *HyperLogLog) Add(item Hash32) {
	// Get the 32-bit hash value
	x := item.Sum32()

	// get the most significant bits from the has value
	p := h.p
	// index
	index := x >> (32 - p)

	// hash the remaining bits
	w := x << p
	zeroBits := leadingZeros(w) + 1

	//if we have found a new higher number of zeros for that
	// register we update it
	if zeroBits > h.reg[index] {
		h.reg[index] = zeroBits
	}

}

func (h *HyperLogLog) Count() float64 {
	var sum float64
	for i := 0; i < (1 << h.p); i++ {
		sum += 1.0 / math.Pow(2, float64(h.reg[i]))
	}
	// estimate := h.alpha * float64(1<<uint(h.p)) * float64(1<<uint(h.p)) / sum
	estimate := h.alpha * math.Pow(2, float64(h.p)) * math.Pow(2, float64(h.p)) / sum

	if estimate <= float64((5.0/2.0)*math.Pow(2, float64(h.p))) {
		v := 0
		for i := 0; i < (1 << h.p); i++ {
			if h.reg[i] == 0 {
				v++
			}
		}
		if v != 0 {
			return math.Pow(2, float64(h.p)) * math.Log(math.Pow(2, float64(h.p))/float64(v))
		}
	} else if estimate > float64(1.0/30.0)*float64(1<<32) {
		return -float64(1<<32) * math.Log(1.0-estimate/float64(1<<32))
	}
	return estimate
}

func leadingZeros(x uint32) uint8 {
	// calculates the number of leading zeros
	var zeros uint8

	// we are looping from right to left
	for i := 31; i >= 0; i-- {
		if (x & (1 << uint(i))) == 0 {
			zeros++
		} else {
			break
		}
	}
	return zeros
}

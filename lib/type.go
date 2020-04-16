package lib

func CopyZZZ(values []ZZZ) []ZZZ {
	dst := make([]ZZZ, len(values))
	copy(dst, values)
	return dst
}

func ReverseZZZ(values []ZZZ) []ZZZ {
	newValues := CopyZZZ(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

type Func2ZZZ func(v1, v2 ZZZ) ZZZ

type CounterZZZ struct {
	M     map[ZZZ]int64
	Func2 Func2ZZZ
}

func NewCounterZZZ() *CounterZZZ {
	return &CounterZZZ{M: map[ZZZ]int64{}}
}

func (c CounterZZZ) CountBy2ZZZ(i, j ZZZ) {
	c.M[c.Func2(i, j)]++
}

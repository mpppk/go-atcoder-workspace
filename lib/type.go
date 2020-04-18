package lib

// CopyZZZ は与えられたsliceをコピーして返します。
func CopyZZZ(values []ZZZ) []ZZZ {
	dst := make([]ZZZ, len(values))
	copy(dst, values)
	return dst
}

// ReverseZZZ は与えられたsliceの順序を逆転させたsliceを返します。(非破壊)
func ReverseZZZ(values []ZZZ) []ZZZ {
	newValues := CopyZZZ(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		newValues[i], newValues[j] = values[j], values[i]
	}
	return newValues
}

type FuncZZZ func(values ...ZZZ) ZZZ

// ZZZCounter は値ごとの出現頻度をカウントするための構造体です。
type ZZZCounter struct {
	M    map[ZZZ]int64
	Func FuncZZZ
}

func NewCounterZZZ() *ZZZCounter {
	return &ZZZCounter{M: map[ZZZ]int64{}}
}

// Increment は与えられたキーのカウントを増やします。
func (c ZZZCounter) Increment(value ZZZ) {
	c.M[value]++
}

// IncrementBy 与えられた値から、事前に設定された関数によって計算されたキーのカウントを増やします。
func (c ZZZCounter) IncrementBy(values ...ZZZ) {
	c.M[c.Func(values...)]++
}

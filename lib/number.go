package lib

import "errors"

// MAXAAA は与えられた値のうち、最大であるものを返します。
// 値が一つも与えられなかった場合、エラーを返します。
func MaxAAA(values ...AAA) (max AAA, err error) {
	if len(values) == 0 {
		return 0, errors.New("empty slice is given")
	}

	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}

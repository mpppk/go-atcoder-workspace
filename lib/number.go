package lib

import "errors"

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

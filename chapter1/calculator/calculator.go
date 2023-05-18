package calculator

import "errors"

func Addition(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, errors.New("invalid number")
	}

	res := x + y
	return res, nil
}

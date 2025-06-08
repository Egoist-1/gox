package slice

import "gox/internal/err"

func Add[T any](src []T, ele T, idx int) ([]T, error) {
	if idx < 0 || idx > len(src) {
		return nil, err.NewIndexErr(idx)
	}
	if idx == 0 {
		res := []T{ele}
		return append(res, src...), nil
	}
	if len(src) == idx {
		return append(src, ele), nil
	}
	res := make([]T, 0, 0)
	res = append(res, src[0:idx]...)
	res = append(res, ele)
	res = append(res, src[idx:]...)
	return res, nil
}

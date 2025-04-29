package slicex

// 切片映射
func Map[src any, dst any](slice []src, fn func(s src) dst) []dst {
	res := make([]dst, len(slice))
	for _, val := range slice {
		r := fn(val)
		res = append(res, r)
	}
	return res
}

// 切片转map
func ToMap[T any, key comparable](src []T, fn func(T) (k key)) map[key]T {
	res := make(map[key]T)
	for _, v := range src {
		k1 := fn(v)
		res[k1] = v
	}
	return res
}

// 切片转mapV2
func ToMapV2[T any, key comparable, val any](src []T, fn func(T) (k key, v val)) map[key]val {
	res := make(map[key]val)
	for _, v := range src {
		k1, v1 := fn(v)
		res[k1] = v1
	}
	return res
}

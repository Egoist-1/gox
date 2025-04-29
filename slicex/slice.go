package slicex

// Map 切片映射
func Map[src any, dst any](slice []src, fn func(s src) dst) []dst {
	res := make([]dst, 0, len(slice))
	for _, val := range slice {
		r := fn(val)
		res = append(res, r)
	}
	return res
}

// ToMap 切片转Map, T 切片类型 , map的val 为切片类型,
func ToMap[T any, key comparable](src []T, fn func(T) (k key)) map[key]T {
	res := make(map[key]T)
	for _, v := range src {
		k1 := fn(v)
		res[k1] = v
	}
	return res
}

// ToMapV2 切片转Map, T 切片类型 , map的val可以指定类型
func ToMapV2[T any, key comparable, val any](src []T, fn func(T) (k key, v val)) map[key]val {
	res := make(map[key]val)
	for _, v := range src {
		k1, v1 := fn(v)
		res[k1] = v1
	}
	return res
}

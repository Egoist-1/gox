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
//func ToMap[slice []T, k,v T, T any](src slice, fn func(T) (key, val string)) m {
//	res := make(map[k]v)
//	for _, val := range src {
//		key, v := fn(val)
//		res[key] = v
//	}
//	return res
//}

package mapx

// ToSlice map转切片 ,key,val 为 map的 k,v类型 ,T 为切片类型
func ToSlice[Map ~map[key]val, T any, key comparable, val any](m Map, fn func(key, val) T) []T {
	res := make([]T, 0, len(m))
	for k, v := range m {
		res = append(res, fn(k, v))
	}
	return res
}
func Contain() {

}

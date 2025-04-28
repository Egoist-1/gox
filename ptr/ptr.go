package ptr

// 转为指针
func ToPtr[T any](t T) *T {
	return &t
}

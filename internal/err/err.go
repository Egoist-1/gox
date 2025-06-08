package err

import "fmt"

func NewIndexErr(idx int) error {
	return fmt.Errorf("下标越界index:%d\n", idx)
}
func NewCapFull() error {
	return fmt.Errorf("容量已满\n")
}
func NewNullData() error {
	return fmt.Errorf("空结构")
}

package queue

import (
	"container/list"
	"gox/internal/err"
)

// 如果 len = 0 为无界队列
func NewNormalQueue[T any](len int) *NormalQueue[T] {
	return &NormalQueue[T]{
		data: make([]T, 0, len),
	}
}

type NormalQueue[T any] struct {
	unused     int
	len        int
	frontIndex int
	data       []T
}

func T() {
	list.New()
}

func (p *NormalQueue[T]) Peek() (T, error) {
	if p.len == 0 {
		var res T
		return res, err.NewIndexErr(0)
	}
	return p.data[p.frontIndex], nil
}

func (p *NormalQueue[T]) Unused() int {
	p.unused = cap(p.data) - len(p.data)
	return p.unused
}

func (p *NormalQueue[T]) Len() int {
	return len(p.data)
}
func (p *NormalQueue[T]) Cap() int {
	return cap(p.data)
}
func (p *NormalQueue[T]) Enqueue(ele T) error {
	if p.len == p.Cap() {
		return err.NewCapFull()
	}
	//(frontIdx + len -1 ) % cap
	p.len += 1
	idx := (p.frontIndex + p.len - 1) % p.Cap()
	p.data[idx] = ele
	return nil
}

func (p *NormalQueue[T]) Dequeue() (T, error) {
	var res T
	if p.len == 0 {
		return res, err.NewNullData()
	}
	res = p.data[p.frontIndex]
	if p.frontIndex+1 >= p.Cap() {
		p.frontIndex = 0
	} else {
		p.frontIndex += 1
	}
	return res, nil
}

package queue

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Peek() (T, error)
	Cap() int
	Len() int
}

package limit

type SlidWindows struct {
	WindowSize int64
	bucket     chan int64
}

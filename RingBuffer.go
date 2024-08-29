package main

type ringBuffer[T any] struct{
    buf []T
	head int
	tail int
	len int
	cap int
}

func newRingBuffer[T any](cap int) *ringBuffer[T]{
	return &ringBuffer[T]{
		buf: make([]T,cap),
		cap: cap,
	}
	
}

func (r *ringBuffer[T]) enqueue(data T){
	if r.len == r.cap{
		r.head = (r.head + 1) % r.cap
	}else{
		r.len++
	}
	r.buf[r.tail] = data
	r.tail = (r.tail + 1) % r.cap
}

func (r *ringBuffer[T]) dequeue()T{
	if r.len == 0{
		panic("ringBuffer is empty")
	}
	b := r.buf[r.head] 
	r.head = (r.head + 1) % r.cap
	r.len--
	return b
}




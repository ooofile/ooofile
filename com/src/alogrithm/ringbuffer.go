package alogrithm

import (
	"errors"
)

type RingBufferElem interface{}

type RingBuffer struct {
	data    []RingBufferElem
	in, out int
	size    int
}

func NewRingBuffer(n int) *RingBuffer {
	return &RingBuffer{make([]RingBufferElem, n), 0, 0, 0}
}

func (rb *RingBuffer) next(cursor int) int {
	return cursor & (cap(rb.data) - 1)
}

func (rb *RingBuffer) Put(e RingBufferElem) error {
	if rb.size >= cap(rb.data) {
		return errors.New("ring buffer is full")
	}
	rb.size++
	rb.data[rb.in] = e
	rb.in = rb.next(rb.in)
	return nil
}

func (rb *RingBuffer) Get() RingBufferElem {
	if rb.size <= 0 {
		return nil
	}
	rb.size--
	re := rb.data[rb.out]
	rb.out = rb.next(rb.out)
	return re
}

func (rb *RingBuffer) IsFull() bool {
	if rb.data == nil {
		return true
	}
	if rb.size >= cap(rb.data) && rb.data != nil {
		return true
	}
	return false
}
func (rb *RingBuffer) Len() int {
	return rb.size
}

package alogrithm

import (
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(10)
	for i := 0; i < 10; i++ {
		rb.Put(i)
	}
	if rb.Len() != 10 {
		t.Fatal("ring buffer length error")
	}
	err := rb.Put(11)
	if err == nil {
		t.Fatal("ring buffer is full ")
	}
	for i := 0; i < 10; i++ {
		if i != rb.Get() {
			t.Fatal("ring buffer get error")
		}
	}
	if rb.Len() > 0 {
		t.Fatal("ring buffer length error")
	}
	i := rb.Get()
	if i != nil {
		t.Fatal("ring buffer is empty")
	}
}

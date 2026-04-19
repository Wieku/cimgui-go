package vectors

import "C"
import (
	"runtime"
	"unsafe"
)

// Vector is a representation of ImVector from Dear ImGui.
type Vector[T any] struct {
	buf    []T
	pinner *runtime.Pinner
}

func NewVector[T any](size, capacity int) Vector[T] {
	buf := make([]T, size, capacity)

	return Vector[T]{buf: buf, pinner: &runtime.Pinner{}}
}

func NewVectorFromC[T, V any, CINT ~int32](size, capacity CINT, data *V, ctor func(data V) T) Vector[T] {
	buf := make([]T, size, capacity)

	for i, d := range unsafe.Slice(data, size) {
		buf[i] = ctor(d)
	}

	return Vector[T]{buf: buf, pinner: &runtime.Pinner{}}
}

func (v *Vector[T]) Unpin() {
	v.pinner.Unpin()
}

// Slice converts a Vector to a slice []T.
func (v Vector[T]) Slice() []T {
	return v.buf
}

func (v Vector[T]) Push(val T) {
	v.buf = append(v.buf, val)
}

func (v Vector[T]) Size() int {
	return len(v.buf)
}

func (v Vector[T]) Capacity() int {
	return cap(v.buf)
}

func (v Vector[T]) CData(handler func(T) (any, func())) any {
	buf2 := make([]any, len(v.buf), cap(v.buf))

	for i := range v.buf {
		han, end := handler(v.buf[len(v.buf)-i-1])

		buf2[i] = han

		end()
	}

	return &buf2[0]
}

func ImArray[V, T any](v Vector[T], handler func(T) (V, func())) *V {
	nBuf := make([]V, len(v.buf), cap(v.buf))

	for i, d := range v.buf {
		han, end := handler(d)

		nBuf[i] = han

		end()
	}

	v.pinner.Pin(&nBuf[0])

	return &nBuf[0]
}

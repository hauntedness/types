package types

import (
	"unsafe"
)

type Arena struct {
	data []int
	len  int // current length of data
}

func NewArena() *Arena {
	return &Arena{
		data: make([]int, 64),
		len:  0,
	}
}

func (a *Arena) Free() {
	a.len = 0
	a.data = nil
}

func (a *Arena) Reset() {
	a.len = 0
	a.data = a.data[:0]
}

// Alloc alloc n bytes space.
//
//	# n always be promoted to multiplies of 8(4 in 32bit os)
func (a *Arena) Alloc(n int) (offset int) {
	words := (n + int(unsafe.Sizeof((*int)(nil))-1)) / 8
	ret := a.len
	if ret+words > len(a.data) {
		a.data = append(a.data, make([]int, len(a.data))...)
	}
	a.len = ret + words
	return ret
}

// weak reference of T on arena
//
//	# WeakRef[T] get rid of garbage collector but then could not benefit from go type system!
//	use this carefully
type WeakRef[T any] struct {
	arena  *Arena
	offset int
}

// NewWeakRef make weak reference of T on [Arena].
func NewWeakRef[T any](arena *Arena) WeakRef[T] {
	offset := arena.Alloc(int(unsafe.Sizeof((*T)(nil))))
	return WeakRef[T]{arena: arena, offset: offset}
}

func (wr WeakRef[T]) Size() int {
	return int(unsafe.Sizeof((*T)(nil)))
}

func (wr WeakRef[T]) MustSet(i int, v int) {
	wr.arena.data[wr.offset+i] = v
}

func (wr WeakRef[T]) MustGet(i int) int {
	return wr.arena.data[wr.offset+i]
}

func (wr WeakRef[T]) Offset() int {
	return wr.offset
}

func (wr WeakRef[T]) MustSetSlice(i int, slice []int) {
	copy(wr.arena.data[wr.offset+i:], slice)
}

func (wr WeakRef[T]) MustGetSlice(i int) []int {
	return wr.arena.data[wr.offset+i : wr.offset+wr.Size()]
}

func (wr WeakRef[T]) Pointer() *T {
	return (*T)(unsafe.Pointer(uintptr(unsafe.Pointer(wr.arena)) + uintptr(wr.offset)*unsafe.Sizeof((*int)(nil))))
}

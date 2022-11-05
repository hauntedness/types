package types

import (
	"unsafe"
)

// weak reference of T on arena
//
//	WeakRef[T] get rid of garbage collector but then could not benefit from go type system!
//	use this carefully
type WeakRef[T any] struct {
	arena    *Arena
	offset   int
	pointers *[]unsafe.Pointer
}

// NewWeakRef make weak reference of T on [Arena].
func NewWeakRef[T any](arena *Arena) WeakRef[T] {
	offset := arena.Alloc(int(unsafe.Sizeof(*(*T)(nil))))
	p := make([]unsafe.Pointer, 0)
	return WeakRef[T]{arena: arena, offset: offset, pointers: &p}
}

func (wr WeakRef[T]) Size() int {
	return int(unsafe.Sizeof(*(*T)(nil)))
}

func (wr WeakRef[T]) Offset() int {
	return wr.offset
}

func (wr WeakRef[T]) MustSetInt(i int, v int) {
	wr.arena.data[wr.offset+i] = v
}

func (wr WeakRef[T]) MustGetInt(i int) int {
	return wr.arena.data[wr.offset+i]
}

// use pointer receiver is not recommend as the
func (wr WeakRef[T]) MustSetString(i int, v string) {
	wr.MustSetPointer(i, unsafe.Pointer(&v))
}

func (wr WeakRef[T]) MustGetString(i int) string {
	return *(*string)(wr.MustGetPointer(i))
}

func (wr WeakRef[T]) MustSetPointer(i int, ptr unsafe.Pointer) {
	wr.arena.data[wr.offset+i] = len(*wr.pointers)
	*wr.pointers = append(*wr.pointers, ptr)
}

func (wr WeakRef[T]) MustGetPointer(i int) unsafe.Pointer {
	return (*wr.pointers)[wr.arena.data[wr.offset+i]]
}

func (wr WeakRef[T]) MustSetBytes(i int, v []byte) {
	wr.MustSetPointer(i, unsafe.Pointer(&v))
}

func (wr WeakRef[T]) MustGetSlice(i int) []byte {
	return *(*[]byte)(wr.MustGetPointer(i))
}

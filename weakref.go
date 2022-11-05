package types

import (
	"unsafe"
)

// weak reference of T on arena
//
//	WeakRef[T] get rid of garbage collector but then could not benefit from go type system!
//	use this carefully
type WeakRef struct {
	arena    *Arena
	offset   int
	pointers *[]unsafe.Pointer
}

// NewWeakRef make weak reference of T on [Arena].
func NewWeakRef(arena *Arena, size int) WeakRef {
	offset := arena.Alloc(size)
	p := make([]unsafe.Pointer, 0)
	return WeakRef{arena: arena, offset: offset, pointers: &p}
}

func (wr WeakRef) Offset() int {
	return wr.offset
}

func (wr WeakRef) MustSetInt(i int, v int) {
	wr.arena.data[wr.offset+i] = v
}

func (wr WeakRef) MustGetInt(i int) int {
	return wr.arena.data[wr.offset+i]
}

func (wr WeakRef) MustSetBytes(i int, v []byte) {
	wr.MustSetPointer(i, unsafe.Pointer(&v))
}

func (wr WeakRef) MustGetBytes(i int) []byte {
	return *(*[]byte)(wr.MustGetPointer(i))
}

// set pointer string or []byte is not recommend
// as the native way is better choice
func (wr WeakRef) MustSetString(i int, v string) {
	wr.MustSetPointer(i, unsafe.Pointer(&v))
}

func (wr WeakRef) MustGetString(i int) string {
	return *(*string)(wr.MustGetPointer(i))
}

func (wr WeakRef) MustSetPointer(i int, ptr unsafe.Pointer) {
	wr.arena.data[wr.offset+i] = len(*wr.pointers)
	*wr.pointers = append(*wr.pointers, ptr)
}

func (wr WeakRef) MustGetPointer(i int) unsafe.Pointer {
	return (*wr.pointers)[wr.arena.data[wr.offset+i]]
}

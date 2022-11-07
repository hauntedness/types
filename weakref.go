package types

// weak reference of T on arena
//
//	WeakRef[T] get rid of garbage collector but then could not benefit from go type system!
//	use this carefully
type WeakRef struct {
	arena  *Arena
	offset int
}

// NewWeakRef make weak reference of T on [Arena].
func NewWeakRef(arena *Arena, size int) WeakRef {
	offset := arena.Alloc(size)
	return WeakRef{arena: arena, offset: offset}
}

func (wr WeakRef) Offset() int {
	return wr.offset
}

func (wr WeakRef) SetInt(i int, v int) {
	wr.arena.data[wr.offset+i] = v
}

func (wr WeakRef) GetInt(i int) int {
	return wr.arena.data[wr.offset+i]
}

func (wr WeakRef) SetBytes(i int, v []byte) {
	wr.arena.data[wr.offset+i] = len(wr.arena.bytes)
	wr.arena.data[wr.offset+i+1] = len(v)
	wr.arena.bytes = append(wr.arena.bytes, v...)
}

func (wr WeakRef) GetBytes(i int) []byte {
	start := wr.arena.data[wr.offset+i]
	length := wr.arena.data[wr.offset+i+1]
	ret := make([]byte, length)
	copy(ret, wr.arena.bytes[start:])
	return ret
}

func (wr WeakRef) SetString(i int, v string) {
	wr.SetBytes(i, []byte(v))
}

func (wr WeakRef) GetString(i int) string {
	return string(wr.GetBytes(i))
}

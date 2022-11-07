package types

// weak reference of T on arena
//
//	WeakRef[T] get rid of garbage collector but then could not benefit from go type system!
//	use this carefully
type WeakRef struct {
	a *Arena // arena
	o int    // offset
}

// NewWeakRef make weak reference of T on [Arena].
func NewWeakRef(arena *Arena, size int) WeakRef {
	offset := arena.Alloc(size)
	return WeakRef{a: arena, o: offset}
}

func (w WeakRef) Offset() int {
	return w.o
}

func (w WeakRef) SetInt(i int, v int) {
	w.a.data[w.o+i] = v
}

func (w WeakRef) GetInt(i int) int {
	return w.a.data[w.o+i]
}

func (w WeakRef) SetBytes(i int, v []byte) {
	w.a.data[w.o+i] = len(w.a.bytes)
	w.a.data[w.o+i+1] = len(v)
	w.a.bytes = append(w.a.bytes, v...)
}

func (w WeakRef) GetBytes(i int) []byte {
	start := w.a.data[w.o+i]
	length := w.a.data[w.o+i+1]
	ret := make([]byte, length)
	copy(ret, w.a.bytes[start:])
	return ret
}

func (w WeakRef) SetString(i int, v string) {
	w.SetBytes(i, []byte(v))
}

func (w WeakRef) GetString(i int) string {
	return string(w.GetBytes(i))
}

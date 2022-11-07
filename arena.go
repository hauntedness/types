package types

import "unsafe"

// Arena Arena represent a continuous region of memory
//
//	for easy use, all value stored in int type
//	if need, you can put multiple tiny objects size in one integer
type Arena struct {
	data  []int  // primitive data
	len   int    // current length of data
	bytes []byte // fat pointer storage data like []byte or string
}

// NewArena create Arena of 512 bytes
func NewArena() *Arena {
	return &Arena{
		data:  make([]int, 64),
		len:   0,
		bytes: make([]byte, 512),
	}
}

func NewArenaWithSize(size int, byteSize int) *Arena {
	return &Arena{
		data:  make([]int, size),
		len:   0,
		bytes: make([]byte, byteSize),
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
//	n always be promoted to multiplies of 8(4 in 32bit os)
func (a *Arena) Alloc(n int) (offset int) {
	words := (n + int(unsafe.Sizeof((*int)(nil))-1)) / 8
	ret := a.len
	for ret+words > len(a.data) {
		a.data = append(a.data, make([]int, len(a.data))...)
	}
	a.len = ret + words
	return ret
}

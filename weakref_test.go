package types

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestWeakRef_Offset(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	wr2 := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	fmt.Println(wr2.Offset())
	// text := "hello world"
	wr.MustSetInt(0, 1)
	fmt.Println(arena.data)
}

type SomeStruct struct {
	Ptr *int
	Num int
}

type Layout struct {
	SomeText string
	SomeInt  int
	SomeBool bool
}

func TestWeakRef_Pointer(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*SomeStruct)(nil))))
	var i int = 100
	wr.MustSetPointer(0, unsafe.Pointer(&i))
	ret := wr.MustGetPointer(0)
	fmt.Println(*(*int)(ret))

	wr2 := NewWeakRef(arena, int(unsafe.Sizeof(*(*SomeStruct)(nil))))
	wr2.MustSetPointer(0, unsafe.Pointer(&i))
	ret2 := wr2.MustGetPointer(0)
	fmt.Println(*(*int)(ret2))
}

func TestWeakRef_String(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	wr.MustSetString(0, "123")
	ret := wr.MustGetString(0)
	if ret != "123" {
		t.Error("want 123")
	}
}

func TestWeakRef_Bytes(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	bytes := []byte("123")
	wr.MustSetBytes(0, bytes)
	ret := wr.MustGetBytes(0)
	if !reflect.DeepEqual(bytes, ret) {
		t.Error("want 123")
	}
}

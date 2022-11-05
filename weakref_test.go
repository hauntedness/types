package types

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestWeakRef_Offset(t *testing.T) {
	var layout Layout
	arena := NewArena()
	wr := NewWeakRef[Layout](arena)
	if int(unsafe.Sizeof(layout)) != wr.Size() {
		t.Error("weak reference has wrong size")
	}
	wr2 := NewWeakRef[Layout](arena)
	fmt.Println(wr2.Offset())
	// text := "hello world"
	wr.MustSetInt(0, 1)
	fmt.Println(arena.data)
}

func TestWeakRef_Pointer(t *testing.T) {
	type SomeStruct struct {
		Ptr *int
		Num int
	}
	arena := NewArena()
	wr := NewWeakRef[SomeStruct](arena)
	var i int = 100
	wr.MustSetPointer(0, unsafe.Pointer(&i))
	ret := wr.MustGetPointer(0)
	fmt.Println(*(*int)(ret))

	wr2 := NewWeakRef[SomeStruct](arena)
	wr2.MustSetPointer(0, unsafe.Pointer(&i))
	ret2 := wr2.MustGetPointer(0)
	fmt.Println(*(*int)(ret2))
}

func TestWeakRef_String(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef[Layout](arena)
	wr.MustSetString(0, "123")
	ret := wr.MustGetString(0)
	if ret != "123" {
		t.Error("want 123")
	}
}

type Layout struct {
	SomeText string
	SomeInt  int
	SomeBool bool
}

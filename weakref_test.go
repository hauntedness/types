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
	wr.SetInt(0, 1)
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

func TestWeakRef_String(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	wr.SetString(0, "123")
	ret := wr.GetString(0)
	if ret != "123" {
		t.Error("want 123")
	}
}

func TestWeakRef_Bytes(t *testing.T) {
	arena := NewArena()
	wr := NewWeakRef(arena, int(unsafe.Sizeof(*(*Layout)(nil))))
	bytes := []byte("123")
	wr.SetBytes(0, bytes)
	ret := wr.GetBytes(0)
	if !reflect.DeepEqual(bytes, ret) {
		t.Error("want 123")
	}
}

package types

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestSize(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			in:  1, // 0001
			out: 1, // 1000
		},
		{
			in:  7, // 0111
			out: 1, // 1000
		},
		{
			in:  8, //
			out: 1,
		},
		{
			in:  15,
			out: 2,
		},
		{
			in:  17,
			out: 3,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			ret := (tt.in + int(unsafe.Sizeof((*int)(nil))-1)) / 8
			if ret != tt.out {
				t.Errorf("in:%d,out:%d,expect:%d", tt.in, ret, tt.out)
			}
		})
	}
}

func TestXxx(t *testing.T) {
	a := 1 << 2
	if a == 8 {
		t.Error(a)
	}
}

func TestSizeOf(t *testing.T) {
	u := unsafe.Sizeof((*int)(nil))
	if u != 8 {
		t.Error("expect 8 bytes")
	}
}

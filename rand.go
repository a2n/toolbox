package deg_

import (
	"crypto/rand"
	"encoding/binary"
)

var (
	Int     func() int
	Int31   func() int32
	Int63   func() int64
	Uint32  func() uint32
	Uint64  func() uint64
	Float32 func() float32
	Float64 func() float64
)

func init() {
	Int = func() int {
		u := uint(Int63())
		return int(u << 1 >> 1)
	}

	Int31 = func() int32 {
		var n int32
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}

	Int63 = func() int64 {
		var n int64
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}

	Uint32 = func() uint32 {
		var n uint32
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}

	Uint64 = func() uint64 {
		var n uint64
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}

	Float32 = func() float32 {
		var n float32
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}

	Float64 = func() float64 {
		var n float64
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		return n
	}
}

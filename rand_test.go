package toolbox

import (
	"fmt"
	"sort"
	"testing"
)

func TestInt31(t *testing.T) {
	t.Log(Int31())
}

func BenchmarkIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int31()
	}
}

func TestRand(t *testing.T) {
	t.Logf("Int: %d", Int())
	t.Logf("Int31: %d", Int31())
	t.Logf("Int63: %d", Int63())
	t.Logf("Uint32: %d", Uint32())
	t.Logf("Uint64: %d", Uint64())
	t.Logf("Float32: %v", Float32())
	t.Logf("Float64: %v", Float64())
	t.Logf("Int()%%6: %d", Int()%6)
}

func TestInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", Int())
		}
		fmt.Println()
	}
}

func TestMod52(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 52*8; i++ {
		m[Int()%52]++
	}

	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	n := 0
	for k := range keys {
		fmt.Printf("%d: %d  ", k, m[k])
		n++
		if n >= 9 {
			fmt.Println()
			n = 0
		}
	}
	fmt.Println()
}

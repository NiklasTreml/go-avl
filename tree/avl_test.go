package tree

import (
	"math/rand"
	"testing"
)

var result *Tree

func benchmarkInsertAVL(i int, b *testing.B) {
	b.StopTimer()
	var res *Tree
	values := randomValues(i, i)
	b.SetBytes(0)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		t := New(0)

		for _, v := range values {
			res = InsertAVL(t, v)
		}
	}
	result = res
}

func randomValues(amount, max int) []int {
	result := []int{}
	for i := 0; i < amount; i++ {
		result = append(result, rand.Intn(max))
	}

	return result
}

func BenchmarkAVL10(b *testing.B)          { benchmarkInsertAVL(10, b) }
func BenchmarkAVL100(b *testing.B)         { benchmarkInsertAVL(100, b) }
func BenchmarkAVL1000(b *testing.B)        { benchmarkInsertAVL(1000, b) }
func BenchmarkAVL10_000(b *testing.B)      { benchmarkInsertAVL(10000, b) }
func BenchmarkAVL100_000(b *testing.B)     { benchmarkInsertAVL(100000, b) }
func BenchmarkAVL1_000_000(b *testing.B)   { benchmarkInsertAVL(1000000, b) }
func BenchmarkAVL10_000_000(b *testing.B)  { benchmarkInsertAVL(10000000, b) }
func BenchmarkAVL100_000_000(b *testing.B) { benchmarkInsertAVL(100000000, b) }

func benchmarkDeleteAVL(i int, b *testing.B) {
	b.StopTimer()
	var res *Tree
	values := randomValues(i, i)
	b.SetBytes(0)
	t := New(0)

	for _, v := range values {
		res = InsertAVL(t, v)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		value := values[rand.Intn(len(values))]
		res = DeleteAVL(t, value)
	}

	result = res
}

func BenchmarkDeleteAVL10(b *testing.B)          { benchmarkDeleteAVL(10, b) }
func BenchmarkDeleteAVL100(b *testing.B)         { benchmarkDeleteAVL(100, b) }
func BenchmarkDeleteAVL1000(b *testing.B)        { benchmarkDeleteAVL(1000, b) }
func BenchmarkDeleteAVL10_000(b *testing.B)      { benchmarkDeleteAVL(10000, b) }
func BenchmarkDeleteAVL100_000(b *testing.B)     { benchmarkDeleteAVL(100000, b) }
func BenchmarkDeleteAVL1_000_000(b *testing.B)   { benchmarkDeleteAVL(1000000, b) }
func BenchmarkDeleteAVL10_000_000(b *testing.B)  { benchmarkDeleteAVL(10000000, b) }
func BenchmarkDeleteAVL100_000_000(b *testing.B) { benchmarkDeleteAVL(100000000, b) }

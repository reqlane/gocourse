package advanced

import (
	"math/rand"
	"testing"
)

// === go test file.go
// === go test -bench=. file.go
// === go test -bench=. -benchmem file.go
// === go test -bench=. -memprofile mem.pprof file.go
//    === go tool pprof mem.pprof
// 		 === top
// 		 === list mallocgc
// 		 === o
// 		 === help
// 		 === quit

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := generateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("Expected slice size: %d; received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for b.Loop() {
		generateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := generateRandomSlice(1000)
	b.ResetTimer()
	for b.Loop() {
		sumSlice(slice)
	}
}

// func add(a, b int) int {
// 	return a + b
// }

// // ===== BENCHMARKING
// func BenchmarkAddSmallInput(b *testing.B) {
// 	for b.Loop() {
// 		add(2, 3)
// 	}
// }

// func BenchmarkAddMediumInput(b *testing.B) {
// 	for b.Loop() {
// 		add(200, 300)
// 	}
// }

// func BenchmarkAddLargeInput(b *testing.B) {
// 	for b.Loop() {
// 		add(2000, 3000)
// 	}
// }

// ===== TESTING
// func TestAddSubtests(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			result := add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("result = %d; want %d", result, test.expected)
// 			}
// 		})
// 	}
// }

// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}

// 	for _, test := range tests {
// 		result := add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
// 		}
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("add(2, 3) = %d; want %d", result, expected)
// 	}
// }

package main

import "testing"

func calculateTest(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calucate(input)
	}
	
}

func BenchmarkB1(b *testing.B) {
	calculateTest(100, b)
}

// func B2(b *testing.B) {
// 	calculateTest(2000, b)
// }
// func B3(b *testing.B) {
// 	calculateTest(3000, b)
// }

package main 

import (
	"testing"
)

func BenchmarkAsync(b *testing.B) {
	Async()
}

func BenchmarkConv(b *testing.B) {
	Conv()
}
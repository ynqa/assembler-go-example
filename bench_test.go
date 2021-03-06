package main

import (
	"flag"
	"testing"

	"github.com/ynqa/asm-go-example/asm"
	"github.com/ynqa/asm-go-example/cgo"
	"github.com/ynqa/asm-go-example/slice"
)

var dimension = flag.Int("dim", 1024, "dimension of vectors. "+
	"(the value must be multiple of 8 because of no considering the case of fraction)")

func TestSliceDot(t *testing.T) {
	length := 8
	x := make([]float32, length)
	y := make([]float32, length)

	for i := 0; i < length; i++ {
		x[i] = 2.0
		y[i] = 3.0
	}

	res := slice.Dot(x, y)
	var expected float32 = 48
	if expected != res {
		t.Errorf("AddAsm returns wrong answer %v:%v", expected, res)
	}
}

func TestDotCgo(t *testing.T) {
	length := 8

	x := cgo.Malloc32(length)
	y := cgo.Malloc32(length)

	defer func() {
		cgo.Free32(x)
		cgo.Free32(y)
	}()

	for i := 0; i < length; i++ {
		x[i] = 2.0
		y[i] = 3.0
	}

	res := cgo.Dot(length, x, y)
	var expected float32 = 48
	if expected != res {
		t.Errorf("AddCgo returns wrong answer %v:%v", expected, res)
	}
}

func TestDotAsm(t *testing.T) {
	length := 8
	x := make([]float32, length)
	y := make([]float32, length)

	for i := 0; i < length; i++ {
		x[i] = 2.0
		y[i] = 3.0
	}

	res := asm.Dot(x, y)
	var expected float32 = 48
	if expected != res {
		t.Errorf("AddAsm returns wrong answer %v:%v", expected, res)
	}
}
func BenchmarkDotSlice(b *testing.B) {
	flag.Parse()
	x := make([]float32, *dimension)
	y := make([]float32, *dimension)

	for i := 0; i < *dimension; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice.Dot(x, y)
	}
}

func BenchmarkDotCgo(b *testing.B) {
	flag.Parse()
	x := cgo.Malloc32(*dimension)
	y := cgo.Malloc32(*dimension)

	defer func() {
		cgo.Free32(x)
		cgo.Free32(y)
	}()

	for i := 0; i < *dimension; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cgo.Dot(*dimension, x, y)
	}
	b.StopTimer()
}

func BenchmarkDotAsm(b *testing.B) {
	flag.Parse()
	x := make([]float32, *dimension)
	y := make([]float32, *dimension)

	for i := 0; i < *dimension; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		asm.Dot(x, y)
	}
}

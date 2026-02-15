package petname_test

import (
	"fmt"
	"math/rand"

	"github.com/dustinkirkland/golang-petname"
)

// Example of using the traditional package-level functions (backward compatible)
func ExampleGenerate() {
	// Old API - still works exactly as before
	// Uses global rand (auto-seeded in Go 1.20+)
	name := petname.Generate(2, "-")
	fmt.Println(name)
}

// Example of using the new Generator API with deterministic seed
func ExampleNew_deterministic() {
	// New API - deterministic generation for testing
	gen := petname.New(rand.New(rand.NewSource(42)))

	// Generate predictable names
	fmt.Println(gen.Generate(2, "-"))
	fmt.Println(gen.Generate(2, "-"))
	fmt.Println(gen.Generate(2, "-"))
	// Output:
	// guiding-dodo
	// relieved-bass
	// mature-zebra
}

// Example of using the new Generator API for concurrent use
func ExampleNew_concurrent() {
	// Each goroutine can have its own generator for thread-safe concurrent use
	gen1 := petname.New(rand.New(rand.NewSource(1)))
	gen2 := petname.New(rand.New(rand.NewSource(2)))

	// These can be safely used concurrently
	name1 := gen1.Generate(2, "-")
	name2 := gen2.Generate(2, "-")

	fmt.Println(name1)
	fmt.Println(name2)
	// Output:
	// touched-shark
	// relaxing-muskox
}

// Example of generating individual words with deterministic behavior
func ExampleGenerator_Adjective() {
	gen := petname.New(rand.New(rand.NewSource(100)))

	fmt.Println(gen.Adjective())
	fmt.Println(gen.Adverb())
	fmt.Println(gen.Name())
	// Output:
	// skilled
	// apparently
	// prawn
}

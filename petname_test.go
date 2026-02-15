/*
  petname: test of library for generating human-readable, random names
           for objects (e.g. hostnames, containers, blobs)

  Copyright 2014 Dustin Kirkland <dustin.kirkland@gmail.com>

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

// Package petname is a library for generating human-readable, random
// names for objects (e.g. hostnames, containers, blobs).
package petname

import (
	"math/rand"
	"sync"
	"testing"
)

// Make sure the generated names exist
func TestPetName(t *testing.T) {
	for i:=0; i<10; i++ {
		name := Generate(i, "-")
		if name == "" {
			t.Fatalf("Did not generate a %d-word name, '%s'", i, name)
		}
	}
}

// Test backward compatibility - existing API still works
func TestBackwardCompatibility(t *testing.T) {
	tests := []struct {
		name  string
		fn    func() string
	}{
		{"Adverb", Adverb},
		{"Adjective", Adjective},
		{"Name", Name},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn()
			if result == "" {
				t.Errorf("%s() returned empty string", tt.name)
			}
		})
	}

	// Test Generate with various word counts
	for words := 1; words <= 5; words++ {
		t.Run("Generate", func(t *testing.T) {
			result := Generate(words, "-")
			if result == "" {
				t.Errorf("Generate(%d, '-') returned empty string", words)
			}
		})
	}
}

// Test new Generator API works
func TestGeneratorAPI(t *testing.T) {
	gen := New(rand.New(rand.NewSource(42)))

	tests := []struct {
		name  string
		fn    func() string
	}{
		{"Adverb", gen.Adverb},
		{"Adjective", gen.Adjective},
		{"Name", gen.Name},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn()
			if result == "" {
				t.Errorf("%s() returned empty string", tt.name)
			}
		})
	}

	// Test Generate with various word counts
	for words := 1; words <= 5; words++ {
		t.Run("Generate", func(t *testing.T) {
			result := gen.Generate(words, "-")
			if result == "" {
				t.Errorf("Generate(%d, '-') returned empty string", words)
			}
		})
	}
}

// Test deterministic behavior with seeded rand
func TestDeterministicBehavior(t *testing.T) {
	seed := int64(12345)

	// Generate with first generator
	gen1 := New(rand.New(rand.NewSource(seed)))
	results1 := []string{
		gen1.Adverb(),
		gen1.Adjective(),
		gen1.Name(),
		gen1.Generate(2, "-"),
		gen1.Generate(3, "_"),
	}

	// Generate with second generator using same seed
	gen2 := New(rand.New(rand.NewSource(seed)))
	results2 := []string{
		gen2.Adverb(),
		gen2.Adjective(),
		gen2.Name(),
		gen2.Generate(2, "-"),
		gen2.Generate(3, "_"),
	}

	// Results should be identical
	for i := range results1 {
		if results1[i] != results2[i] {
			t.Errorf("Deterministic mismatch at index %d: got %q, want %q", i, results2[i], results1[i])
		}
	}
}

// Test thread-safe concurrent use with separate generators
func TestConcurrentGenerators(t *testing.T) {
	const numGoroutines = 100
	const numGenerations = 10

	var wg sync.WaitGroup
	results := make([][]string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Each goroutine gets its own generator with unique seed
			gen := New(rand.New(rand.NewSource(int64(id))))
			results[id] = make([]string, numGenerations)
			for j := 0; j < numGenerations; j++ {
				results[id][j] = gen.Generate(2, "-")
			}
		}(i)
	}

	wg.Wait()

	// Verify all goroutines completed and generated names
	for i, res := range results {
		if len(res) != numGenerations {
			t.Errorf("Goroutine %d: expected %d results, got %d", i, numGenerations, len(res))
		}
		for j, name := range res {
			if name == "" {
				t.Errorf("Goroutine %d, generation %d: got empty name", i, j)
			}
		}
	}
}

// Test that nil rand uses global rand (backward compatibility)
func TestNilRandUsesGlobalRand(t *testing.T) {
	gen := New(nil)

	// Should not panic and should generate names
	result := gen.Generate(2, "-")
	if result == "" {
		t.Error("Generator with nil rand returned empty string")
	}

	adverb := gen.Adverb()
	if adverb == "" {
		t.Error("Generator with nil rand Adverb() returned empty string")
	}

	adjective := gen.Adjective()
	if adjective == "" {
		t.Error("Generator with nil rand Adjective() returned empty string")
	}

	name := gen.Name()
	if name == "" {
		t.Error("Generator with nil rand Name() returned empty string")
	}
}

// Test separator handling
func TestSeparators(t *testing.T) {
	gen := New(rand.New(rand.NewSource(42)))

	separators := []string{"-", "_", ".", "", " ", "::", "---"}
	for _, sep := range separators {
		result := gen.Generate(3, sep)
		if result == "" {
			t.Errorf("Generate(3, %q) returned empty string", sep)
		}
	}
}

// Benchmark old API
func BenchmarkGenerateOldAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(2, "-")
	}
}

// Benchmark new API with global rand (nil)
func BenchmarkGenerateNewAPIGlobalRand(b *testing.B) {
	gen := New(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Generate(2, "-")
	}
}

// Benchmark new API with dedicated rand
func BenchmarkGenerateNewAPIDedicatedRand(b *testing.B) {
	gen := New(rand.New(rand.NewSource(42)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Generate(2, "-")
	}
}

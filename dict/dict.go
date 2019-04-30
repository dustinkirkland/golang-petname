package dict

import (
	"math/rand"
	"strings"
)

// DIct implements petname.Dict by providing random adverbs, adjectives and names
// which are combined to compose a pet name.
type Dict struct {
	Adverbs    WordSource
	Adjectives WordSource
	Names      WordSource
}

// Adverb returns a random adverb from the dictionary.
func (d *Dict) Adverb() string {
	return d.Adverbs.Rand()
}

// Adjective returns a random adjective from the dictionary.
func (d *Dict) Adjective() string {
	return d.Adjectives.Rand()
}

// Name returns a random name from the dictionary.
func (d *Dict) Name() string {
	return d.Names.Rand()
}

// WordSource represents a collection of line-separated words with a count.
//
// Why not a slice? Well, petname has large dictionaries, some with over 50,000 words.
// Storing 50,000 strings in a slice would allocate 50,000 pointers to those strings,
// which adds up to 400 KB of wasted memory on 64-bit architectures. Additionally,
// allocating this many pointers may slow down the GC as it looks for unused pointers
// to dispose of.
type WordSource struct {
	Words string
	Count int
}

// Word returns the word at the given index from the collection,
// or an empty string if index is out of bounds.
func (s WordSource) Word(index int) string {
	pos := 0
	for i, r := range s.Words {
		if pos == index {
			// Return current word beginning at `i` until the next line ending at `end`.
			end := i + strings.IndexRune(s.Words[i:], '\n')
			return s.Words[i:end]
		}
		if r == '\n' {
			pos++
		}
	}
	return ""
}

// Rand returns a random word from the collection.
func (s WordSource) Rand() string {
	return s.Word(rand.Intn(s.Count))
}

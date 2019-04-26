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
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/zippoxer/golang-petname/dict"
	"github.com/zippoxer/golang-petname/dict/large"
	"github.com/zippoxer/golang-petname/dict/medium"
	"github.com/zippoxer/golang-petname/dict/small"
)

func TestDicts(t *testing.T) {
	dicts := []*dict.Dict{
		small.Dict,
		medium.Dict,
		large.Dict,
	}
	for di, d := range dicts {
		wordSources := []dict.WordSource{
			d.Adverbs,
			d.Adjectives,
			d.Names,
		}
		for wi, w := range wordSources {
			if err := testWordSource(w); err != nil {
				t.Fatalf("word source %d of dict %d: %v",
					wi, di, err)
			}
		}
	}
}

// Make sure the generated names exist
func testWordSource(w dict.WordSource) error {
	lines := strings.Count(w.Words, "\n")
	if w.Count != lines {
		return fmt.Errorf("has %d lines but says it has %d words", lines, w.Count)
	}

	// Test word access by iterating the source's lines and comparing them
	// to the result of `WordSource.Word(lineIndex)`.
	scanner := bufio.NewScanner(strings.NewReader(w.Words))
	expectingEOF := false
	for i := 0; scanner.Scan(); i++ {
		if expectingEOF {
			return fmt.Errorf("expected eof after blank line #%d, got %q instead", i, scanner.Text())
		}
		if scanner.Text() == "" {
			expectingEOF = true
		}
		if scanner.Text() != w.Word(i) {
			return fmt.Errorf("word at line %d should be %q, got %q instead", i, scanner.Text(), w.Word(i))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Make sure the generated names exist.
func TestPetName(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := Generate(small.Dict, i, "-")
		if name == "" {
			t.Fatalf("Did not generate a %d-word name, '%s'", i, name)
		}
	}
}

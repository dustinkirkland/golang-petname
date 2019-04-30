/*
  petname: library for generating human-readable, random names
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
	"strings"
)

// Dict represents a petname dictionary consisting of 3 word lists:
// adverbs, adjectives and names.
//
// These word lists are imported from the project at:
//   - https://github.com/dustinkirkland/petname
// into the dict/small, dict/medium and dict/large packages.
type Dict interface {
	// Adverb returns a random adverb from a list of petname adverbs.
	Adverb() string

	// Adjective returns a random adjective from a list of petname adjectives.
	Adjective() string

	// Name returns a random name from a list of petname names.
	Name() string
}

// Generate generates and returns a random pet name.
// It takes two parameters:  the number of words in the name, and a separator token.
// If a single word is requested, simply a Name() is returned.
// If two words are requested, a Adjective() and a Name() are returned.
// If three or more words are requested, a variable number of Adverb() and a Adjective and a Name() is returned.
// The separator can be any charater, string, or the empty string.
func Generate(dict Dict, words int, separator string) string {
	if words == 1 {
		return dict.Name()
	} else if words == 2 {
		return dict.Adjective() + separator + dict.Name()
	}
	var petname []string
	for i := 0; i < words-2; i++ {
		petname = append(petname, dict.Adverb())
	}
	petname = append(petname, dict.Adjective(), dict.Name())
	return strings.Join(petname, separator)
}

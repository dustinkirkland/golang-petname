package small

import "github.com/dustinkirkland/golang-petname/dict"

var Dict = &dict.Dict{
	Adverbs: dict.WordSource{
		Words: adverbs,
		Count: 261,
	},
	Adjectives: dict.WordSource{
		Words: adjectives,
		Count: 449,
	},
	Names: dict.WordSource{
		Words: names,
		Count: 456,
	},
}

package large

import "github.com/dustinkirkland/golang-petname/dict"

var Dict = &dict.Dict{
	Adverbs: dict.WordSource{
		Words: adverbs,
		Count: 12545,
	},
	Adjectives: dict.WordSource{
		Words: adjectives,
		Count: 36577,
	},
	Names: dict.WordSource{
		Words: names,
		Count: 5899,
	},
}

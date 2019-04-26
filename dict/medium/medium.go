package medium

import "github.com/zippoxer/golang-petname/dict"

var Dict = &dict.Dict{
	Adverbs: dict.WordSource{
		Words: adverbs,
		Count: 1630,
	},
	Adjectives: dict.WordSource{
		Words: adjectives,
		Count: 1198,
	},
	Names: dict.WordSource{
		Words: names,
		Count: 1060,
	},
}

# PetName

This utility will generate "pet names", consisting of a random combination of an adverb, adjective, and proper name.  These are useful for unique hostnames, for instance.

As such, PetName tries to follow the tenets of Zooko's triangle.  Names are:

 - human meaningful
 - decentralized
 - secure

The default packaging includes:

 - 1,933 names
 - 1,315 adjectives
 - 3,656 adverbs

A 1-word PetName consists of one random name.  A 2-word Petname consists of a random adjective and a random name.  A 3-word (or more than 3 word) PetName consists of random adverb(s) and an adjective and a name.

 - 2-word PetNames yield 1,933 x 1,315 = 2,541,895 unique combinations
 - 3-word PetNames yield 1,933 x 1,315 x 3,656 = 9,293,168,120 unique combinations
 - 4-word PetNames yield 1,933 x 1,315 x 3,656 x 3,656 = 3.397582265×10¹³ unique combinations

## Command Line Usage

Command line help:

    usage: petname [--words INT] [--separator STR]

    optional arguments:
        --words            number of words in the name, default is 2
        --separator        string used to separate name words, default is '-'

## Command Line Examples

    $ petname
    wiggly-Anna

    $ petname --words 1
    Marco

    $ petname --words 3
    quickly-scornful-Johnathan

    $ petname --words 4
    dolorously-leisurely-wee-Susan

    $ petname --separator ":"
    hospitable:Isla

    $ petname --separator "" --words 3
    adeptlystaticNicole

## Golang Examples
```golang
package main

import (
        "flag"
        "fmt"
        "github.com/dustinkirkland/golang-petname"
)

var (
        words = flag.Int("words", 2, "The number of words in the pet name")
        separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func main() {
        flag.Parse()
        fmt.Println(petname.PetName(*words, *separator))
}
```

## Python Examples

See: https://pypi.python.org/pypi/petname

    $ pip install petname
    $ sudo apt-get install python-petname

```python
import argparse
from petname import *

parser = argparse.ArgumentParser(description='Generate human readable random names')
parser.add_argument('-w', '--words', help='Number of words in name, default=2', default=2)
parser.add_argument('-s', '--separator', help='Separator between words, default="-"', default="-")
parser.options = parser.parse_args()

print PetName(int(parser.options.words), parser.options.separator)
```

## Credits

This project is authored and maintained by Dustin Kirkland.


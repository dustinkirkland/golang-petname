# PetName

This utility will generate "pet names", consisting of a random combination of an adverb, adjective, and proper name.  These are useful for unique hostnames, for instance.

As such, PetName tries to follow the tenets of Zooko's triangle.  Names are:

 - human meaningful
 - decentralized
 - secure

The default packaging includes:

 - 5,931 names
 - 37,389 adjectives
 - 12,814 adverbs

A 1-word PetName consists of one random name.  A 2-word Petname consists of a random adjective and a random name.  A 3-word (or more than 3 word) PetName consists of random adverb(s) and an adjective and a name.

 - 2-word PetNames yield 5,931 x 37,389 = 221,754,159 unique combinations
 - 3-word PetNames yield 5,931 x 37,389 x 12,814 = 2.8415578x10^12 unique combinations

## Command Line Usage

Command line help:

    usage: petname [--words INT] [--separator STR]

    optional arguments:
        -w|--words            number of words in the name, default is 2
        -s|--separator        string used to separate name words, default is '-'
        -d|--dir              directory containing adverbs.txt, adjectives.txt, names.txt, default is \fI/usr/share/petname/\fP

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
        fmt.Println(petname.Generate(*words, *separator))
}
```

## Python Examples

See: https://pypi.python.org/pypi/petname

    $ pip install petname
    $ sudo apt-get install python-petname

```python
import argparse
import petname

parser = argparse.ArgumentParser(description='Generate human readable random names')
parser.add_argument('-w', '--words', help='Number of words in name, default=2', default=2)
parser.add_argument('-s', '--separator', help='Separator between words, default="-"', default="-")
parser.options = parser.parse_args()

print petname.Generate(int(parser.options.words), parser.options.separator)
```

## Credits

This project is authored and maintained by Dustin Kirkland.


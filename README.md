# PetName

This utility will generate "pet names", consisting of a random combination of an adverb, adjective, and proper name.  These are useful for unique hostnames, for instance.

The default packaging contains about 2000 names, 1300 adjectives, and 4000 adverbs, yielding nearly 10 billion unique combinations, covering over 32 bits of unique namespace.

As such, PetName tries to follow the tenets of Zooko's triangle.  Names are:

 - human meaningful
 - decentralized
 - secure


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

## Python Examples

## Editing Word Lists

## Credits

This project is authored and maintained by Dustin Kirkland.


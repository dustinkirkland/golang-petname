#!/bin/sh

set -e

PKG="petname"

for f in adverbs adjectives names; do
	rm -f "$f".txt.list
	for w in $(cat /usr/lib/petname/"$f".txt); do
		printf '"%s", ' "$w" >> "$f".txt.list
	done
	sed -i -e "s/, $//" "$f".txt.list
	sed -i -e "s/^        $f    = [...]string{.*$/        $f    = [...]string{$(cat ${f}.txt.list)}/" ${PKG}.go
	rm -f "$f".txt.list
done

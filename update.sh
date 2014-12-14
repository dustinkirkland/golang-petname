#!/bin/sh

set -e

PKG="petname"

for f in adverbs adjectives names; do
	rm -f "$f".txt.list
	for w in $(cat "$f"); do
		printf '"%s", ' "$w" >> "$f".txt.list
	done
	sed -i -e "s/, $//" "$f".txt.list
	sed -i -e "s/^        $f    = [...]string{.*$/        $f    = [...]string{$(cat ${f}.list)}/" ${PKG}.go.unf
	rm -f "$f".list
done

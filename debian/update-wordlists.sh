#!/bin/sh

# This script only needs to be run by the upstream package maintainer (Dustin Kirkland)
# if the upstream petname wordlists change

set -e

PKG="petname"

for f in adverbs adjectives names; do
	rm -f "$f".txt.list
	printf "	$f = [...]string{" > "$f".txt.list
	for w in $(wget -q -O- https://raw.githubusercontent.com/dustinkirkland/petname/master/usr/share/petname/small/${f}.txt); do
		printf '"%s", ' "$w" >> "$f".txt.list
	done
	sed -i -e "s/, $/}\n/" "$f".txt.list
	sed -i "/^\s\+${f}\s\+= \[\.\.\.\]string{.*$/d" ${PKG}.go
done
printf "\n)\n\n" >> "$f".txt.list
grep -B 1000 "^var (" ${PKG}.go > above
grep -A 1000 "^// End word lists" ${PKG}.go > below
cat above *.txt.list below > ${PKG}.go
go fmt ${PKG}.go
rm -f *.txt.list above below
wget -q -O- https://raw.githubusercontent.com/dustinkirkland/petname/master/README.md > README.md

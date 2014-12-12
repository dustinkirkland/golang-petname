#!/bin/sh

set -e

PKG="petname"

cp -f ${PKG}.go.in ${PKG}.go.unf

for f in ../../etc/$PKG/*; do
	filename=$(basename "$f")
	rm -f "$filename"
	for w in $(cat "$f"); do
		printf '"%s", ' "$w" >> "$filename"
	done
	sed -i -e "s/, $//" "$filename"
	sed -i -e "s/__${filename}__/$(cat ${filename})/" ${PKG}.go.unf
	gofmt ${PKG}.go.unf > ${PKG}.go
	rm -f "$filename"
done
rm ${PKG}.go.unf
go build ${PKG}.go

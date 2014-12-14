#!/bin/sh

set -e

PKG="petname"

cp -f ${PKG}.go.in ${PKG}.go.unf
cp -f ${PKG}.py.in ${PKG}.py.unf

for f in usr/lib/$PKG/*; do
	filename=$(basename "$f")
	rm -f "$filename"
	for w in $(cat "$f"); do
		printf '"%s", ' "$w" >> "$filename"
	done
	sed -i -e "s/, $//" "$filename"
	sed -i -e "s/__${filename}__/$(cat ${filename})/" ${PKG}.go.unf
	sed -i -e "s/__${filename}__/$(cat ${filename})/" ${PKG}.py.unf
	rm -f "$filename"
done
gofmt ${PKG}.go.unf > ${PKG}.go
autopep8 ${PKG}.py.unf > ${PKG}.py
rm -f ${PKG}.go.unf ${PKG}.py.unf
go build ${PKG}.go
python -m py_compile ${PKG}.py

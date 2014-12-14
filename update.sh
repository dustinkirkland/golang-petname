#!/bin/sh

set -e

PKG="petname"

cp -f ${PKG}.go.in ${PKG}.go.unf
cp -f ${PKG}/${PKG}.py.in ${PKG}/${PKG}.py.unf

for f in adverbs.txt adjectives.txt names.txt; do
	rm -f "$f".list
	for w in $(cat "$f"); do
		printf '"%s", ' "$w" >> "$f".list
	done
	sed -i -e "s/, $//" "$f".list
	sed -i -e "s/__${f}__/$(cat ${f}.list)/" ${PKG}.go.unf
	sed -i -e "s/__${f}__/$(cat ${f}.list)/" ${PKG}/${PKG}.py.unf
	rm -f "$f".list
done
gofmt ${PKG}.go.unf > ${PKG}.go
autopep8 ${PKG}/${PKG}.py.unf > ${PKG}/${PKG}.py
rm -f ${PKG}.go.unf ${PKG}/${PKG}.py.unf
go build ${PKG}.go
python -m py_compile ${PKG}/${PKG}.py
rm -f ${PKG}/*pyc

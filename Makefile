all:
	make -C src/petname
	make -C src/petname_cmd

clean:
	make -C src/petname clean
	make -C src/petname_cmd clean
	rm -rf pkg

.PHONY: all clean

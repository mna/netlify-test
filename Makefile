.DEFAULT=build
PUB_DIR=pub/
TPL_DIR=src/tpl/

.PHONY: build
build: clean
	go run src/cmd/build.go $(TPL_DIR) $(PUB_DIR)

.PHONY: clean
clean:
	rm -rf pub/*


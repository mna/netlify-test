.DEFAULT=build

PUB_DIR=pub/
TPL_DIR=src/tpl/
BIN_DIR=functions/
FUNC_DIR=src/funcs/

.PHONY: build
build: clean build-html build-lambda

.PHONY: build-html
build-html:
	go run src/cmd/build.go $(TPL_DIR) $(PUB_DIR)

.PHONY: build-lambda
build-lambda:
	mkdir -p $(FUNC_DIR)
	go get ./...
	go build -o $(BIN_DIR)hello $(FUNC_DIR)hello.go

.PHONY: clean
clean:
	rm -rf $(PUB_DIR)*
	rm -rf $(BIN_DIR)*


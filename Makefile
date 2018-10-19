.DEFAULT=build

PUB_DIR=./pub/
TPL_DIR=./src/tpl/
BIN_DIR=./functions/
FUNC_DIR=./src/funcs/

.PHONY: build
build: clean build-html build-lambda

.PHONY: build-html
build-html:
	go run src/cmd/build.go $(TPL_DIR) $(PUB_DIR)

.PHONY: build-lambda
build-lambda:
	mkdir -p $(BIN_DIR)
	mkdir -p $(PUB_DIR)
	go version
	go env
	go get ./...
	./scripts/build-go.sh $(FUNC_DIR) $(BIN_DIR)
	./scripts/build-js.sh $(FUNC_DIR) $(BIN_DIR)

.PHONY: clean
clean:
	rm -rf $(PUB_DIR)*
	rm -rf $(BIN_DIR)*


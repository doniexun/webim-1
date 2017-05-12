GO=$(shell which go)
GB=$(shell unalias gb;which gb)

vendor:
	$(GB) vendor fetch ${pkg}
	@echo "download" ${pkg} "success"
build:
	$(GO) build

test: build
	./webim -dbname testwebim -user root -pass root -dbport 3306

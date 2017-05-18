GO=$(shell which go)

build:
	GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo
build-dev: build
	docker build -t adolphlwq/webim:0.1.dev -f Dockerfile.dev .
push-dev: build-dev
	docker push adolphlwq/webim:0.1.dev
test:
	./webim -dbname testwebim -user root -pass root -dbport 3306 -

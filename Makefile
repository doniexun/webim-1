GO=$(shell which go)

build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o webim
build-lls:
	docker build -t adolphlwq/webim:lls -f Dockerfile.dev .
push-lls: build-dev
	docker push adolphlwq/webim:lls
test:
	./webim -dbname testwebim -user root -pass root -dbport 3306

GO_BUILD_FLAGS=
GO_TEST_FLAGS=
PKGS=$(shell go list ./... | grep -E -v "(vendor)")

all:
	go build $(GO_BUILD_FLAGS) -o webim
test:
	go test --cover $(GO_TEST_FLAGS) $(PKGS)
dev:
	./webim -config ./dev.properties
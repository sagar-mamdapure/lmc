PROFILE := local
VERSION       ?= $(shell git describe --tags --always --dirty=-SNAPSHOT-`git rev-parse --short HEAD`| tr -d 'v')

default: deploy

gofmt:
ifeq ($(shell uname -s), Darwin)
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofmt -l -d -s -w; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofumpt -l -d -w; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs gci write; sync
	find pkg/ -type f | grep go | egrep -v "mocks|gomock" | xargs goimports -l -d -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofmt -l -d -s -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gofumpt -l -d -w; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs gci write; sync
	find cmd/ -type f | grep go | egrep -v "mocks|gomock" | xargs goimports -l -d -w; sync
	gofmt -l -d -s -w main.go; sync
	gofumpt -l -d -w main.go; sync
	gci write main.go; sync
	goimports -l -d -w main.go; sync
endif

.PHONY:
deploy:
	@helm plugin uninstall lmc 2> /dev/null || true
	@goreleaser release --skip-publish --rm-dist  --snapshot
	@helm plugin install .
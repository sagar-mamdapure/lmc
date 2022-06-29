PROFILE := local
VERSION       ?= $(shell git describe --tags --always --dirty=-SNAPSHOT-`git rev-parse --short HEAD`| tr -d 'v')

default: deploy

.PHONY:
deploy:
	@helm plugin uninstall lmc 2> /dev/null || true
	@goreleaser release --skip-publish --rm-dist  --snapshot
	@helm plugin install .
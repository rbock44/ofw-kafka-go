GO=GO111MODULE=on go

all: update bench test

update:
	cd kafka && $(GO) mod vendor

generate:
	cd kafka && $(GO) generate -mod=vendor

bench: 
	cd kafka && $(GO) test -mod=vendor -bench=.

test:
	cd kafka && $(GO) test -v


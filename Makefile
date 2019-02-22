GO=GO111MODULE=on go
update:
	cd kafka && $(GO) mod vendor

generate:
	cd kafka && $(GO) generate -mod=vendor

bench: 
	cd kafka && $(GO) test -mod=vendor -bench=.

test:
	cd kafka && $(GO) test


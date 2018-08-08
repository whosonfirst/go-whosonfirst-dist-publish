CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-dist-publish
	cp *.go src/github.com/whosonfirst/go-whosonfirst-dist-publish/
	cp -r assets src/github.com/whosonfirst/go-whosonfirst-dist-publish/
	cp -r publisher src/github.com/whosonfirst/go-whosonfirst-dist-publish/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-dist"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-repo"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-aws/..."
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-bindata"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-bindata-html-template"
	@GOPATH=$(GOPATH) go get -u "github.com/dustin/go-humanize"
	mv src/github.com/whosonfirst/go-whosonfirst-dist/vendor/github.com/tidwall src/github.com/

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt publisher/*.go

assets: self
	@GOPATH=$(GOPATH) go build -o bin/go-bindata ./vendor/github.com/whosonfirst/go-bindata/go-bindata/
	rm -rf templates/*/*~
	rm -rf assets
	mkdir -p assets/html
	@GOPATH=$(GOPATH) bin/go-bindata -pkg html -o assets/html/html.go templates/html

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-dist-publish cmd/wof-dist-publish.go
	@GOPATH=$(GOPATH) go build -o bin/wof-dist-prune cmd/wof-dist-prune.go
	@GOPATH=$(GOPATH) go build -o bin/wof-dist-index cmd/wof-dist-index.go

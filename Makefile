fmt:
	# go fmt *.go
	go fmt cmd/*.go
	go fmt publisher/*.go

# See this? See the release numbers? I don't know why I need to do this
# but apparently 'go mod vendor' excludes things in a package's cmd folder
# Why... I have no idea. Perhaps I am just going it wrong...
# (20190606/thisisaaronland)

bindata:
	if test ! -d cmd/go-bindata; then mkdir -p cmd/go-bindata; fi
	if test ! -d cmd/go-bindata-assetfs; then mkdir -p cmd/go-bindata-assetfs; fi
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/v0.1.0/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/v1.0.1/cmd/go-bindata-assetfs/main.go
	@echo "This file was cloned from https://raw.githubusercontent.com/whosonfirst/go-bindata/v0.1.0/cmd/go-bindata/main.go" > cmd/go-bindata/README.md
	@echo "This file was cloned from https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/v0.1.0/cmd/go-bindata-assetfs/main.go" > cmd/go-bindata-assetfs/README.md

static:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	rm -rf templates/*/*~
	rm -rf assets
	mkdir -p assets/html
	mkdir -p assets/feed
	bin/go-bindata -pkg html -o assets/html/html.go templates/html
	bin/go-bindata -pkg feed -o assets/feed/feed.go templates/feed

tools:
	go build -mod vendor -o bin/wof-dist-publish cmd/wof-dist-publish/main.go
	go build -mod vendor -o bin/wof-dist-prune cmd/wof-dist-prune/main.go
	go build -mod vendor -o bin/wof-dist-index cmd/wof-dist-index/main.go
	go build -mod vendor -o bin/wof-bundles-index cmd/wof-bundles-index/main.go

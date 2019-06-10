fmt:
	# go fmt *.go
	go fmt cmd/*.go
	go fmt publisher/*.go

assets:
	go build -o bin/go-bindata ./vendor/github.com/whosonfirst/go-bindata/cmd/go-bindata/main.go
	rm -rf templates/*/*~
	rm -rf assets
	mkdir -p assets/html
	mkdir -p assets/feed
	bin/go-bindata -pkg html -o assets/html/html.go templates/html
	bin/go-bindata -pkg feed -o assets/feed/feed.go templates/feed

tools:
	go build -mod vendor -o bin/wof-dist-publish cmd/wof-dist-publish.go
	go build -mod vendor -o bin/wof-dist-prune cmd/wof-dist-prune.go
	go build -mod vendor -o bin/wof-dist-index cmd/wof-dist-index.go
	go build -mod vendor -o bin/wof-bundles-index cmd/wof-bundles-index.go

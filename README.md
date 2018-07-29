# go-whosonfirst-dist-publish

Go package for publishing Who's On First distributions.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

This doesn't work yet. Strictly speaking it does work but I might break everything still so all the usual caveats apply.

## Tools

### wof-dist-publish

```
./bin/wof-dist-publish -h
Usage of ./bin/wof-dist-publish:
  -publisher string
    	Valid publishers are: s3 (default "s3")
  -publisher-dsn string
    	A valid DSN string for your distribution publisher
  -workdir string
    	Where to read build files from. If empty the code will attempt to use the current working directory.
```

For example:

```
$> ./bin/wof-dist-publish -workdir /path/to/workdir -dsn 'bucket=BUCKET region={REGION} prefix={PREFIX} credentials={CREDENTIALS}' <repo>...
```

## DSN strings

### S3

```
bucket=BUCKET region={REGION} prefix={PREFIX} credentials={CREDENTIALS}
```

Valid credentials strings are:

* `env:`

* `iam:`

* `{PATH}:{PROFILE}`

## See also:

* https://github.com/whosonfirst/go-whosonfirst-dist
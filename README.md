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

## For example

```
#!/bin/sh

LIST_REPOS="/usr/local/whosonfirst/go-whosonfirst-github/bin/wof-list-repos"
BUILD_DIST="/usr/local/whosonfirst/go-whosonfirst-dist/bin/wof-dist-build"
PUBLISH_DIST="/usr/local/whosonfirst/go-whosonfirst-dist-publish/bin/wof-dist-publish"
PRUNE_DIST="/usr/local/whosonfirst/go-whosonfirst-dist-publish/bin/wof-dist-prune"

PUBLISHER="s3"

# I am too dumb to make this work...
# PUBLISHER_DSN="bucket=dist.whosonfirst.org region=us-east-2 prefix=test credentials=iam:"

# So we'll just do it the long way...

S3_BUCKET="dist.whosonfirst.org"
S3_REGION="us-east-2"
S3_PREFIX="test"
S3_CREDENTIALS="iam:"

WORKDIR="/usr/local/data/dist"

for REPO in `${LIST_REPOS} -not-forked -updated-since P1D`
do
    
    echo "rebuild distributions for ${REPO}"
    
    ${BUILD_DIST} -workdir ${WORKDIR} -timings -verbose -build-meta -build-bundle ${REPO}

    echo "publish distributions for ${REPO}"
    
    ${PUBLISH_DIST} -workdir ${WORKDIR} -publisher ${PUBLISHER} -publisher-dsn "bucket=${S3_BUCKET} region=${S3_REGION} prefix=${S3_PREFIX} credentials=${S3_CREDENTIALS}" ${REPO}

done

echo "prune distributions"

${PRUNE_DIST} -publisher ${PUBLISHER} -publisher-dsn "bucket=${S3_BUCKET} region=${S3_REGION} prefix=${S3_PREFIX} credentials=${S3_CREDENTIALS}" whosonfirst-data
    
echo "INDEX ME... (translation: please write me)"
```

## See also:

* https://github.com/whosonfirst/go-whosonfirst-dist
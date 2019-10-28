# go-whosonfirst-dist-publish

Go package for publishing Who's On First distributions.

## Install

You will need to have both `Go` (version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## For example

```
#!/bin/sh

LIST_REPOS="/usr/local/whosonfirst/go-whosonfirst-github/bin/wof-list-repos"
BUILD_DIST="/usr/local/whosonfirst/go-whosonfirst-dist/bin/wof-dist-build"
PUBLISH_DIST="/usr/local/whosonfirst/go-whosonfirst-dist-publish/bin/wof-dist-publish"
PRUNE_DIST="/usr/local/whosonfirst/go-whosonfirst-dist-publish/bin/wof-dist-prune"
INDEX_DIST="/usr/local/whosonfirst/go-whosonfirst-dist-publish/bin/wof-dist-index"

PUBLISHER="s3"

# I am too dumb to make this work...
# PUBLISHER_DSN="bucket=dist.whosonfirst.org region=us-east-2 prefix=test credentials=iam:"

# So we'll just do it the long way...

S3_BUCKET="dist.whosonfirst.org"
S3_REGION="us-east-2"
S3_PREFIX="test"
S3_CREDENTIALS="iam:"

WORKDIR="/usr/local/data/dist"
LOCKFILE="${WORKDIR}/.lock"

SINCEFILE="/usr/local/data/dist.txt"
SINCE="P1D"

if [ -f ${LOCKFILE} ]
then
    echo "lockfile '${LOCKFILE}' is present, exiting"
    exit 1
fi

if [ -f ${SINCEFILE} ]
then
    SINCE=`cat ${SINCEFILE}`
    echo "SINCE FROM SINCEFILE ${SINCE}"
fi

echo `date '+%s'` > ${SINCEFILE}

rm -rf ${WORKDIR}/*
echo `date` > ${LOCKFILE}

TO_PUBLISH=$@

if [ "$#" -eq 0 ]
then   
    TO_PUBLISH=`${LIST_REPOS} -not-forked -updated-since ${SINCE}`
fi

if [ "$1" = "all" ]
then
    echo "publish all not-forked repos"    
    TO_PUBLISH=`${LIST_REPOS} -not-forked`
fi

echo "publish '${TO_PUBLISH}'"

for REPO in ${TO_PUBLISH}
do
    
    echo "rebuild distributions for ${REPO}"
    
    ${BUILD_DIST} -workdir ${WORKDIR} -timings -verbose -build-meta -build-bundle ${REPO}

    if [ $? -ne 0 ]
    then
	echo "rebuild failed for ${REPO}"
	continue
    fi
       
    echo "publish distributions for ${REPO}"
    
    ${PUBLISH_DIST} -workdir ${WORKDIR} -publisher ${PUBLISHER} -publisher-dsn "bucket=${S3_BUCKET} region=${S3_REGION} prefix=${S3_PREFIX} credentials=${S3_CREDENTIALS}" ${REPO}

    if [ $? -ne 0 ]
    then
	echo "publish failed for ${REPO}"
	continue
    fi
    
    echo "prune distributions"

    ${PRUNE_DIST} -publisher ${PUBLISHER} -publisher-dsn "bucket=${S3_BUCKET} region=${S3_REGION} prefix=${S3_PREFIX} credentials=${S3_CREDENTIALS}" whosonfirst-data

    if [ $? -ne 0 ]
    then
	echo "pruning failed"
	continue
    fi
    
    echo "index distributions"
    
    ${INDEX_DIST} -publisher ${PUBLISHER} -publisher-dsn "bucket=${S3_BUCKET} region=${S3_REGION} prefix=${S3_PREFIX} credentials=${S3_CREDENTIALS}" whosonfirst-data

    if [ $? -ne 0 ]
    then
	echo "indexing failed"
	continue
    fi
    
done

rm -f ${LOCKFILE}
```

## See also:

* https://github.com/whosonfirst/go-whosonfirst-github
* https://github.com/whosonfirst/go-whosonfirst-dist
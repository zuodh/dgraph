#!/bin/bash

BENCHMARKS_REPO="$(pwd)/benchmarks"
BENCHMARKS_URL=https://github.com/dgraph-io/benchmarks/blob/master/data
SCHEMA_FILE="$BENCHMARKS_REPO/data/1million.schema"
DATA_FILE="$BENCHMARKS_REPO/data/1million.rdf.gz"

wget -O $SCHEMA_FILE $BENCHMARKS_URL/1million.schema?raw=true
wget -O $DATA_FILE $BENCHMARKS_URL/1million.rdf.gz?raw=true

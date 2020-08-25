#!/bin/bash

./start_1_zero_2_alpha.sh

DATA_DIR=$PWD/data
BENCHMARKS_URL=https://github.com/dgraph-io/benchmarks/blob/master/data
SCHEMA_FILE="$DATA_DIR/1million.schema"
DATA_FILE="$DATA_DIR/1million.rdf.gz"

wget -O $SCHEMA_FILE $BENCHMARKS_URL/1million.schema?raw=true
wget -O $DATA_FILE $BENCHMARKS_URL/1million.rdf.gz?raw=true


LOGS_DIR=$PWD"/logs"


echo "Running live loader "
dgraph live --cwd $DATA_DIR -f "$DATA_FILE" -s "$SCHEMA_FILE" &> "$LOGS_DIR/live_loader"

echo "Done loading data"

# smoke test kind of query to fetch current schema
echo "Reading schema"

curl -i -H 'Content-Type: application/graphql+-' -X POST -d 'schema{}' localhost:8080/query


# TODO: run go script here.


echo "All movies of Quentin Tarantino"

query=$(<quentin_movies.gql)

echo "\n Movies of Quentin Tarantino"

response=$(curl -i -H 'Content-Type: application/graphql+-' -X POST localhost:8080/query -d $query)
echo $response
if [[ $response == *"Django Unchained" ]]; then
    echo "Okay response"
else
    echo "Error in response" >> /dev/stderr
fi


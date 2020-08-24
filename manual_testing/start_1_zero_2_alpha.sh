#!/bin/bash

pkill -f dgraph;pkill -f dgraph;pkill -f dgraph

logs_dir=$PWD"/logs"
data_dir=$PWD"/data"

rm -rf $logs_dir; mkdir $logs_dir
rm -rf $data_dir; mkdir $data_dir

echo "Starting a zero and 2 alphas"

o="0"
dgraph zero -w "z$o" --cwd $data_dir &> "$logs_dir/z1" &

sleep 4s
grep "Dgraph version" "$logs_dir/z1"


o="0"
dgraph alpha -o $o -p "p$o" -w "w$o" --cwd $data_dir &> "$logs_dir/a$o" &

o="1"
dgraph alpha -o $o -p "p$o" -w "w$o" --cwd $data_dir &> "$logs_dir/a$o" &


sleep 10s


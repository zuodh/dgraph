#!/bin/bash

pkill -f dgraph;pkill -f dgraph;pkill -f dgraph

logs_dir=$PWD"/logs"
data_dir=$PWD"/data"

rm -rf $logs_dir; mkdir $logs_dir
rm -rf $data_dir; mkdir $data_dir

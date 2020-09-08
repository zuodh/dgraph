#!/usr/bin/env python

import logging
import os.path
import shutil
import subprocess
import time
from collections import Counter

import psutil
import pytest
import requests


# Compare two JSON objects, treating lists as unordered
def json_equals(expected, got):
    def process(obj):
        if isinstance(obj, dict):
            return sorted((k, process(v)) for k, v in obj.items())
        if isinstance(obj, list):
            return sorted(process(ele) for ele in obj)
        return obj

    logging.info(f"Expected: {expected}")
    logging.info(f"Got: {got}")

    assert process(expected) == process(got)


def killall(name: str):
    found = True

    while found:
        found = False

        for proc in psutil.process_iter():
            if proc.name() == name and proc.status() != psutil.STATUS_ZOMBIE:
                found = True
                proc.kill()

        time.sleep(1)


def test_bulk_loader():
    cwd = os.getcwd()

    data_path = os.path.join(cwd, "data")
    logs_path = os.path.join(cwd, "logs")
    dgraph_path = os.path.join(cwd, "dgraph")
    backup_path = os.path.join(cwd, "backup")

    db_schema_url = "https://github.com/dgraph-io/benchmarks/blob/master/data/1million.schema?raw=true"
    db_data_url = "https://github.com/dgraph-io/benchmarks/blob/master/data/1million.rdf.gz?raw=true"

    db_schema_path = os.path.join(data_path, "1million.schema")
    db_data_path = os.path.join(data_path, "1million.rdf.gz")

    dgraph_exe = "dgraph.exe" if os.name == "nt" else "dgraph"

    logging.info("Killing running Dgraph instances")
    killall(dgraph_exe)

    time.sleep(5)

    logging.info("Setting up directories")
    for path in (data_path, logs_path, dgraph_path, backup_path):
        try:
            shutil.rmtree(path)
        except FileNotFoundError:
            pass
        os.mkdir(path)

    logging.info("Downloading schema file")
    response = requests.get(db_schema_url)
    with open(db_schema_path, "wb") as f:
        f.write(response.content)

    logging.info("Downloading data file")
    response = requests.get(db_data_url)
    with open(db_data_path, "wb") as f:
        f.write(response.content)

    logging.info("Spawning Dgraph Zero instance in background")
    with open(os.path.join(logs_path, "zero"), "w") as f:
        subprocess.Popen(["dgraph", "zero", "--cwd", dgraph_path],
                         stdout=f,
                         stderr=f)

    time.sleep(5)

    logging.info("Running Dgraph bulk loader")
    with open(os.path.join(logs_path, "bulk-loader"), "w") as f:
        subprocess.run([
            "dgraph", "bulk", "--schema", db_schema_path, "--files",
            db_data_path, "--reduce_shards", "1", "--map_shards", "1"
        ],
                       stdout=f,
                       stderr=f,
                       check=True)

    logging.info("Copying bulk loader data")
    shutil.copytree(os.path.join(cwd, "out/0/p"),
                    os.path.join(dgraph_path, "p"))

    logging.info("Spawning Dgraph Alpha instance in background")
    with open(os.path.join(logs_path, "alpha"), "w") as f:
        subprocess.Popen(["dgraph", "alpha", "--cwd", dgraph_path],
                         stdout=f,
                         stderr=f)

    time.sleep(5)

    logging.info("Querying Dgraph schema")
    response = requests.post("http://localhost:8080/query",
                             "schema {}",
                             headers={"Content-Type": "application/graphql+-"})

    expected = {
        "schema": [{
            "predicate": "actor.film",
            "type": "uid",
            "count": True,
            "list": True
        }, {
            "predicate": "country",
            "type": "uid",
            "reverse": True,
            "list": True
        }, {
            "predicate": "cut.note",
            "type": "string",
            "lang": True
        }, {
            "predicate": "dgraph.acl.rule",
            "type": "uid",
            "list": True
        }, {
            "predicate": "dgraph.graphql.schema",
            "type": "string"
        }, {
            "predicate": "dgraph.graphql.xid",
            "type": "string",
            "index": True,
            "tokenizer": ["exact"],
            "upsert": True
        }, {
            "predicate": "dgraph.password",
            "type": "password"
        }, {
            "predicate": "dgraph.rule.permission",
            "type": "int"
        }, {
            "predicate": "dgraph.rule.predicate",
            "type": "string",
            "index": True,
            "tokenizer": ["exact"],
            "upsert": True
        }, {
            "predicate": "dgraph.type",
            "type": "string",
            "index": True,
            "tokenizer": ["exact"],
            "list": True
        }, {
            "predicate": "dgraph.user.group",
            "type": "uid",
            "reverse": True,
            "list": True
        }, {
            "predicate": "dgraph.xid",
            "type": "string",
            "index": True,
            "tokenizer": ["exact"],
            "upsert": True
        }, {
            "predicate": "director.film",
            "type": "uid",
            "reverse": True,
            "count": True,
            "list": True
        }, {
            "predicate": "email",
            "type": "string",
            "index": True,
            "tokenizer": ["exact"],
            "upsert": True
        }, {
            "predicate": "genre",
            "type": "uid",
            "reverse": True,
            "count": True,
            "list": True
        }, {
            "predicate": "initial_release_date",
            "type": "datetime",
            "index": True,
            "tokenizer": ["year"]
        }, {
            "predicate": "loc",
            "type": "geo",
            "index": True,
            "tokenizer": ["geo"]
        }, {
            "predicate": "name",
            "type": "string",
            "index": True,
            "tokenizer": ["hash", "term", "trigram", "fulltext"],
            "lang": True
        }, {
            "predicate": "performance.actor",
            "type": "uid",
            "list": True
        }, {
            "predicate": "performance.character",
            "type": "uid",
            "list": True
        }, {
            "predicate": "performance.character_note",
            "type": "string",
            "lang": True
        }, {
            "predicate": "performance.film",
            "type": "uid",
            "list": True
        }, {
            "predicate": "rated",
            "type": "uid",
            "reverse": True,
            "list": True
        }, {
            "predicate": "rating",
            "type": "uid",
            "reverse": True,
            "list": True
        }, {
            "predicate": "starring",
            "type": "uid",
            "count": True,
            "list": True
        }, {
            "predicate": "tagline",
            "type": "string",
            "lang": True
        }],
        "types": [{
            "fields": [{
                "name": "dgraph.graphql.schema"
            }, {
                "name": "dgraph.graphql.xid"
            }],
            "name":
            "dgraph.graphql"
        }]
    }
    got = response.json()["data"]
    json_equals(got, expected)

    logging.info("Querying Dgraph data")
    response = requests.post("http://localhost:8080/query",
                             """\
    {
        peterDinklage(func: eq(name@en, "Peter Dinklage")) {
            uid
            dgraph.type
        }
    }""",
                             headers={"Content-Type": "application/graphql+-"})

    expected = {
        "peterDinklage": [{
            "uid": "0x171f221e9c87314a",
            "dgraph.type": ["Person"]
        }]
    }
    got = response.json()["data"]
    json_equals(expected, got)

    logging.info("Exporting Dgraph data")
    response = requests.get("http://localhost:8080/admin/export")

    expected = {"code": "Success", "message": "Export completed."}
    got = response.json()
    json_equals(got, expected)

    logging.info("Backing up of Dgraph data")
    response = requests.post("http://localhost:8080/admin",
                             f"""\
    mutation {{
        backup(input: {{destination: "{backup_path}"}}) {{
            response {{
                message
                code
            }}
        }}
    }}""",
                             headers={"Content-Type": "application/graphql"})

    got = response = response.json()["data"]
    expected = {
        "backup": {
            "response": {
                "message": "Backup completed.",
                "code": "Success"
            }
        }
    }
    json_equals(got, expected)

    logging.info("Killing running Dgraph instances")
    killall(dgraph_exe)

    logging.info("Deleting Dgraph directory")
    shutil.rmtree(dgraph_path)
    os.mkdir(dgraph_path)

    logging.info("Spawning Dgraph Zero instance in background")
    with open(os.path.join(logs_path, "zero2"), "w") as f:
        subprocess.Popen(["dgraph", "zero", "--cwd", dgraph_path],
                         stdout=f,
                         stderr=f)

    with open(os.path.join(logs_path, "restore"), "w") as f:
        logging.info("Restoring Dgraph data")
        subprocess.run([
            "dgraph", "restore", "--location", backup_path, "--postings",
            dgraph_path, "--zero", "localhost:5080"
        ],
                       stdout=f,
                       stderr=f,
                       check=True)

    os.rename(os.path.join(dgraph_path, "p1"), os.path.join(dgraph_path, "p"))

    logging.info("Spawning Dgraph Alpha instance in background")
    with open(os.path.join(logs_path, "alpha"), "w") as f:
        subprocess.Popen(["dgraph", "alpha", "--cwd", dgraph_path],
                         stdout=f,
                         stderr=f)

    time.sleep(5)

    time.sleep(1000000)

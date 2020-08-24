import pydgraph
import json
query="""{
  q(func: allofterms(name@en, "Quentin Tarantino")) {
    name@en
    director.film {
      name@en
    }
  }
}
"""
client_stub = pydgraph.DgraphClientStub('localhost:9080')
client = pydgraph.DgraphClient(client_stub)
txn = client.txn()
try:
    res = txn.query(query)
    response_string = str(res.json)
    if 'Django Unchained' in response_string:
        print("Success")
except:
    print("Error occured while running quentin query")


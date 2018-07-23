package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func main() {
	runTxn(newClient())
}

type Ident struct {
	Ident string  `json:"ident"`
	RDep  []Ident `json:"~dependency"`
}

func runTxn(c *dgo.Dgraph) {
	ctx := context.Background()
	txn := c.NewTxn()
	defer txn.Discard(ctx)
	variables := map[string]string{"$ident": "core/elixir"}
	const q = `query Deps($ident: string) {
			deps(func: eq(ident, $ident)) @recurse(depth: 2, loop: false) {
				ident,
				~dependency
			}
		}
	`
	resp, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		Deps []Ident `json:"deps"`
	}
	var decode Root
	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}
	fmt.Println(decode)
}

func newClient() *dgo.Dgraph {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

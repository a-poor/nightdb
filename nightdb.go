package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

const defaultQuery = `SELECT a, b FROM t;`

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

func main() {
	log.Println("Starting...")
	q := defaultQuery
	if len(os.Args) > 1 {
		q = os.Args[1]
	}
	log.Printf("Parsing query: %q\n", q)

	log.Println("Parsing query...")
	astNode, err := parse(q)
	if err != nil {
		err = fmt.Errorf("parse error: %w", err)
		log.Panic(err)
	}
	log.Printf("Result: %+v\n", *astNode)
}

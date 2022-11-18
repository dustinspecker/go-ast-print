package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	filepath := flag.String("f", "", "filepath to print AST of")

	flag.Parse()

	if *filepath == "" {
		log.Fatalf("a filepath is required")
	}

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatalf("unexpected error opening %s: %s", *filepath, err.Error())
	}
	defer file.Close()

	fileset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fileset, *filepath, file, 0)
	if err != nil {
		log.Fatalf("unexpected error parsing file %s: %s", *filepath, err.Error())
	}

	ast.Print(fileset, parsedFile)
}

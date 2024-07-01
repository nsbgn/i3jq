// This tool has been adapted from:
// <https://github.com/itchyny/gojq/blob/0607aa5af33a4f980e3e769a1820db80e3cc7b23/_tools/gen_builtin.go>
// +build ignore

package main

import (
	"fmt"
	"go/printer"
	"go/token"
	"os"
	"log"
	"regexp"
	"strings"
	"path/filepath"

	"github.com/itchyny/astgen-go"
	"github.com/itchyny/gojq"
)

func main() {
	out, err := os.Create("builtin.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	out.WriteString(
`// Code generated by go generate; DO NOT EDIT.
package main
import . "github.com/itchyny/gojq"
import "github.com/itchyny/gojq"
`)

	cwd, err := os.Getwd()
	builtin := filepath.Join(cwd, "builtin")
	err = filepath.Walk(builtin, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(path, ".jq") {
			cnt, err := os.ReadFile(path)
			if err != nil {
				log.Fatalln(err)
			}
			q, err := gojq.Parse(string(cnt))
			if err != nil {
				log.Fatalln(err)
			}
			for _, fd := range q.FuncDefs {
				fd.Minify()
			}
			str, err := formatQuery(q)
			if err != nil {
				log.Fatalln(err)
			}
			out.WriteString(*str)
		}
		return nil
	})
}

func getName(q *gojq.Query) *string {
	if q.Meta != nil {
		for _, kv := range q.Meta.KeyVals {
			if kv.Key == "name" {
				return &kv.Val.Str
			}
		}
	}
	str := "unnamed"
	return &str
}

func formatQuery(q *gojq.Query) (*string, error) {
	// Get abstract syntax tree of query
	ast, err := astgen.Build(q)
	if err != nil {
		return nil, err
	}

	// Turn AST into a string
	var sb strings.Builder
	var name = getName(q)
	sb.WriteString("\nvar ")
	sb.WriteString(*name)
	sb.WriteString("Query = ")
	err = printer.Fprint(&sb, token.NewFileSet(), ast)
	if err != nil {
		return nil, err
	}
	sb.WriteString("\n\n")
	str := sb.String()

	// Convert integers to proper enums
	for op := gojq.OpPipe; op <= gojq.OpUpdateAlt; op++ {
		re := regexp.MustCompile(fmt.Sprintf(`\b((?:Update)?Op): %d\b`, op))
		str = re.ReplaceAllString(str, fmt.Sprintf("$1: %#v", op))
	}
	for t := gojq.TermTypeIdentity; t <= gojq.TermTypeQuery; t++ {
		re := regexp.MustCompile(fmt.Sprintf(`(Term{Type): %d\b`, t))
		str = re.ReplaceAllString(str, fmt.Sprintf("$1: %#v", t))
	}

	return &str, nil
}

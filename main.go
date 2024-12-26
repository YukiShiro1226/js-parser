package main

import (
	"fmt"
	"os"
	"time"

	"github.com/t14raptor/go-fast/parser"
	// "github.com/t14raptor/go-fast/resolver"
)

func main() {
	dat, err := os.ReadFile("in.js")
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now()

	a, err := parser.ParseFile(string(dat))

	elapsed := time.Since(start)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(elapsed)

	f, err := os.OpenFile("parsed.js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	formattedString := fmt.Sprintf("%s\n\n", a)
	if _, err := f.WriteString(formattedString); err != nil {
		fmt.Println(err)
	}
}

// type ExampleVisitor struct {
// 	ast.NoopVisitor
// }

// func (v *ExampleVisitor) VisitIdentifier(n *ast.Identifier) {
// 	fmt.Println(n.ToId())
// }

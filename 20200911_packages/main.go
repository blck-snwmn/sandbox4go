package main

import (
	"fmt"
	"go/ast"
	"os"

	"golang.org/x/tools/go/packages"
)

func do2() {
	fmt.Println("test")
}

func main() {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax}
	pkgs, err := packages.Load(cfg, "./testdata")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		// ast.Print(pkg.Fset, pkg)
		for _, s := range pkg.Syntax {
			ast.Inspect(s, func(n ast.Node) bool {
				switch n := n.(type) {
				case *ast.CallExpr:
					switch f := n.Fun.(type) {
					case *ast.SelectorExpr:
						fmt.Printf("f.`X` %v\n", f.X)
						fmt.Printf("f.`Sel` %+v\n", f.Sel)
					default:
						fmt.Printf("f.`X` %v", f)
					}
					fmt.Printf("args=%v\n", n.Args)
					fmt.Printf("Fun=%v\n", n.Fun)
					fmt.Printf("Lparen=%v\n", n.Lparen)
					fmt.Printf("Rparen=%v\n", n.Rparen)
				}
				return true
			})
		}
	}
	fmt.Println(len(pkgs))
	fmt.Println(len(pkgs[0].Syntax))
}

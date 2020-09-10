package main

import (
	"fmt"
	"go/ast"
	"os"

	"golang.org/x/tools/go/packages"
)

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
					fmt.Println("----")
					switch f := n.Fun.(type) {
					case *ast.SelectorExpr:
						// f.X is package name or variable name
						fmt.Printf("call %+v.%v(%q) [*ast.SelectorExpr]\n", f.X, f.Sel, n.Args)
					default:
						fmt.Printf("call %+v(%q), \n", f, n.Args)
					}
				}
				return true
			})
		}
	}
	fmt.Println("----")
	fmt.Println(len(pkgs))
	fmt.Println(len(pkgs[0].Syntax))
}

package main

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
)

func main() {
	// content, err := os.ReadFile("testdata/merge_changes/current/api_user.go")
	// if err != nil {
	// 	panic(err)
	// }
	// f, err := decorator.Parse(string(content))
	// if err != nil {
	// 	panic(err)
	// }

	// // list := f.Decls[0].(*dst.FuncDecl).Body.List
	// // list[0], list[1] = list[1], list[0]
	// // dst.Inspect(f, inspectFunc)
	// dstutil.Apply(f, applyFunc, nil)
}

func applyFunc(c *dstutil.Cursor) bool {
	node := c.Node()
	// Use a switch-case construct based on the node "type"
	// This is a very useful of navigating the AST
	switch n := node.(type) {
	case (*dst.FuncDecl):
		// Pretty print the Node AST
		fmt.Println("\n\n---------------\n\n")
		// dst.Print(n)
		dst.Print(n.Body)
	}

	return true
}

func inspectFunc(node dst.Node) bool {
	if node == nil {
		return false
	}

	before, after, points := dstutil.Decorations(node)

	var info string

	if before != dst.None {
		info += fmt.Sprintf("- Before: %s\n", before)
	}

	for _, point := range points {
		if len(point.Decs) == 0 {
			continue
		}

		info += fmt.Sprintf("- %s: [", point.Name)

		for i, dec := range point.Decs {
			if i > 0 {
				info += ", "
			}
			info += fmt.Sprintf("%q", dec)
		}
		info += "]\n"
	}

	if after != dst.None {
		info += fmt.Sprintf("- After: %s\n", after)
	}

	if info != "" {
		fmt.Printf("%T\n%s\n", node, info)
	}

	return true
}

// see https://github.com/dave/dst/blob/master/dstutil/rewrite_test.go
// for some usage patterns

func newNonImplementedHandlerBody() dst.Field {
	return dst.BlockStmt{List: []dst.Stmt{dst.ExprStmt{}}}
	// Body: *ast.BlockStmt {
	// 	Lbrace: foo:10:48
	// 	List: []ast.Stmt (len = 1) {
	// 		0: *ast.ExprStmt {
	// 			X: *ast.CallExpr {
	// 				Fun: *ast.SelectorExpr {
	// 					X: *ast.Ident {
	// 						NamePos: foo:11:2
	// 						Name: "c"
	// 						Obj: *(obj @ 72)
	// 					}
	// 					Sel: *ast.Ident {
	// 						NamePos: foo:11:4
	// 						Name: "String"
	// 						Obj: nil
	// 					}
	// 				}
	// 				Lparen: foo:11:10
	// 				Args: []ast.Expr (len = 2) {
	// 					0: *ast.SelectorExpr {
	// 						X: *ast.Ident {
	// 							NamePos: foo:11:11
	// 							Name: "http"
	// 							Obj: nil
	// 						}
	// 						Sel: *ast.Ident {
	// 							NamePos: foo:11:16
	// 							Name: "StatusNotImplemented"
	// 							Obj: nil
	// 						}
	// 					}
	// 					1: *ast.BasicLit {
	// 						ValuePos: foo:11:38
	// 						Kind: STRING
	// 						Value: "\"501 not implemented\""
	// 					}
	// 				}
	// 				Ellipsis: -
	// 				Rparen: foo:11:59
	// 			}
	// 		}
	// 	}
	// 	Rbrace: foo:12:1
	// }
}

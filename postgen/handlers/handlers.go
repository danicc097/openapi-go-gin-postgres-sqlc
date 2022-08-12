package main

import (
	"fmt"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

func main() {
	content, err := os.ReadFile("testdata/merge_changes/current/api_user.go")
	if err != nil {
		panic(err)
	}
	f, err := decorator.Parse(string(content))
	if err != nil {
		panic(err)
	}

	// list := f.Decls[0].(*dst.FuncDecl).Body.List
	// list[0], list[1] = list[1], list[0]
	// dst.Inspect(f, inspectFunc)
	// dstutil.Apply(f, applyFunc, nil)
	inspectStruct(f)
}

func applyFunc(c *dstutil.Cursor) bool {
	node := c.Node()
	// Use a switch-case construct based on the node "type"
	// This is a very useful of navigating the AST
	//
	switch n := node.(type) {
	case (*dst.FuncDecl):
		// Pretty print the Node AST
		//
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

type HandlerAST struct {
	Tag         string
	OperationId string
	Fn          *dst.BlockStmt
}

// https://github.com/dave/dst/blob/master/dstutil/rewrite_test.go
// for some usage patterns
// https://developers.mattermost.com/blog/instrumenting-go-code-via-ast-2/
// TODO inspect gen and current and save blocks.
// if any method from gen is not in current, append the method
// if viceversa, do nothing
func inspectStruct(f dst.Node) {
	structName := "User"
	fmt.Printf("structName: %s\n", structName)

	dst.Inspect(f, func(n dst.Node) bool {
		if fn, isFn := n.(*dst.FuncDecl); isFn {
			if fn.Recv != nil && len(fn.Recv.List) == 1 {
				// Check that the receiver is actually the struct we want
				if r, rok := fn.Recv.List[0].Type.(*dst.StarExpr); rok &&
					r.X.(*dst.Ident).Name == structName {
					fmt.Printf("found method %v for struct %s\n", fn.Name, structName)
				}
			}
		}

		return true
	})
}

// func newNonImplementedHandlerBody() dst.BlockStmt {
// 	return dst.BlockStmt{
// 		List: []dst.Stmt{
// 			dst.ExprStmt{
// 				X: dst.CallExpr{
// 					Fun: dst.SelectorExpr{
// 						X: dst.Ident{
// 							Name: "c",
// 						},
// 						Sel: dst.Ident{
// 							Name: "String",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// 		List: []dst.Stmt {
// 			0: dst.ExprStmt {
// 				X: dst.CallExpr {
// 					Fun: dst.SelectorExpr {
// 						X: dst.Ident {
// 							Name: "c"
// 						}
// 						Sel: dst.Ident {
// 							Name: "String"
// 						}
// 					}
// 					Lparen: foo:11:10
// 					Args: []dst.Expr (len = 2) {
// 						0: dst.SelectorExpr {
// 							X: dst.Ident {
// 								NamePos: foo:11:11
// 								Name: "http"
// 								Obj: nil
// 							}
// 							Sel: dst.Ident {
// 								NamePos: foo:11:16
// 								Name: "StatusNotImplemented"
// 								Obj: nil
// 							}
// 						}
// 						1: dst.BasicLit {
// 							ValuePos: foo:11:38
// 							Kind: STRING
// 							Value: "\"501 not implemented\""
// 						}
// 					}
// 					Ellipsis: -
// 					Rparen: foo:11:59
// 				}
// 			}
// 		}
// 	}
// }

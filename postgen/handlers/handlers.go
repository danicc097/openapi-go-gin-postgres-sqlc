package main

import (
	"fmt"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

func main() {
	/*
		Read:
		https://go.dev/src/go/ast
		https://pkg.go.dev/go/ast


		Rationale:
		We dont want users to manually add handlers and routes to handlers/*
		If we let them, we wouldnt know at plain sight what was in the spec and what wasnt
		and parsing will become a bigger mess.
		Users can still add new methods, but the routes slice in Register will have
		all items containing a route () that wasnt generated removed.
		If we need a new route that cant be defined in the spec, e.g. fileserver,
		we purposely want that out of the generated handler struct, so its clear that
		its outside the spec.
		It can still remain in handlers/* as long as its not api_*(!_test).go, e.g. fileserver.go
		and still follow the same handlers struct pattern for all we care, it wont be touched.

		flow:
		glob current/api_*(!_test).go -> currentBasenames
		glob gen/api_*(!_test).go -> genBasenames
		For each gen basename:
			If current basename doesnt exist, cp as is.
			Else:
				1. parse gen:
				- extract slice of routes, which contains all relevant info we will need
				to merge -> genRoutes.
				genRoutes is a map indexed by Route.Name (operation ids are unique).
				TODO Can we easily load a struct ast node into the struct itself?
				- get list of struct methods (inspectStruct) --> genHandlers
				2. parse current:
				- extract slice of routes in the same way --> currentRoutes
				- get list of struct methods (inspectStruct) --> currentHandlers
			While merging:
					Based on assumption that users have not modified Register() (clearly indicated).
					if key of genRoutes is not in currentRoutes:
						- append gen slice node value to current routes slice
						- append gen method (501 status) to current struct.
					IMPORTANT: if a method already exists in current but has no routes item (meaning
					its probably some handler helper method created afterwards) then panic and alert
					the user to rename. it shouldve been unexported or a function in the first place anyway.

	*/

	currentContent, err := os.ReadFile("testdata/merge_changes/current/api_user.go")
	if err != nil {
		panic(err)
	}
	f, err := decorator.Parse(string(currentContent))
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

	switch n := node.(type) {
	case (*dst.FuncDecl):
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

type Handler struct {
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
					fmt.Printf("body is %v\n", fn.Body)
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

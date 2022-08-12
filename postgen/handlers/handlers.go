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

	dst.Inspect(f, func(node dst.Node) bool {
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
	})
}

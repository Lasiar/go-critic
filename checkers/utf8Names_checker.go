package checkers

import (
	"go/ast"
	"unicode"

	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unnecessaryDefer"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects redundantly deferred calls"
	info.Before = `
func() {
	defer os.Remove(filename)
}`
	info.After = `
func() {
	os.Remove(filename)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &utf8Namxes{ctx: ctx}, nil
	})
}

type utf8Names struct {
	ctx    *linter.CheckerContext
	isFunc bool
}

func (c *utf8Names) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			c.checkName(decl.Name)
		}
	}
}

func (c *utf8Names) isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

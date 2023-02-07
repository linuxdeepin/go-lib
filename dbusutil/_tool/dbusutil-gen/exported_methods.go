// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"unicode"

	"github.com/linuxdeepin/go-lib/strv"

	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (g *Generator) genExportedMethods(types strv.Strv) {
	fSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fSet, ".", func(info os.FileInfo) bool {
		if info.Name() == filepath.Base(_outputFile) {
			return false
		}
		return true
	}, 0)
	if err != nil {
		log.Fatal(err)
	}
	result := make(map[string]dbusutil.ExportedMethods)
	for _, typ := range types {
		result[typ] = dbusutil.ExportedMethods{}
	}
	for _, pkg := range pkgs {

		for _, file := range pkg.Files {
			dbusIdent := getDBusIdent(file.Imports)

			ast.Inspect(file, func(node ast.Node) bool {
				funcDecl, ok := node.(*ast.FuncDecl)
				if !ok {
					return true
				}

				if funcDecl.Recv != nil && len(funcDecl.Recv.List) == 1 {
					var recvTypeName string
					recv := funcDecl.Recv.List[0]
					starExpr, ok := recv.Type.(*ast.StarExpr)
					if ok {
						if ident, ok := starExpr.X.(*ast.Ident); ok {
							recvTypeName = ident.Name
							if !types.Contains(recvTypeName) {
								return false
							}
						} else {
							return false
						}
					} else {
						return false
					}
					if isTargetFunc(funcDecl, dbusIdent) {
						in, out := getArgs(funcDecl, dbusIdent)
						result[recvTypeName] = append(result[recvTypeName], dbusutil.ExportedMethod{
							Name:    funcDecl.Name.Name,
							InArgs:  in,
							OutArgs: out,
						})
					}

				}

				return false
			})
		}

	}
	g.genGetExportedMethodsFn(result)
}

func (g *Generator) genGetExportedMethodsFn(result map[string]dbusutil.ExportedMethods) {
	typeNames := make([]string, 0, len(result))
	for typeName := range result {
		typeNames = append(typeNames, typeName)
	}
	sort.Strings(typeNames)
	for _, typeName := range typeNames {
		methods := result[typeName]
		sort.Sort(methods)
		g.printf("func (v *%v) GetExportedMethods() dbusutil.ExportedMethods {\n", typeName)
		if len(methods) == 0 {
			g.printf("return nil\n")
		} else {
			g.printf("return dbusutil.ExportedMethods{\n")
			for _, method := range methods {
				g.printf("{\n")
				g.printf("Name: %q,\n", method.Name)
				g.printf("Fn: v.%v,\n", method.Name)
				if len(method.InArgs) > 0 {
					g.printf("InArgs: %#v,\n", method.InArgs)
				}
				if len(method.OutArgs) > 0 {
					g.printf("OutArgs: %#v,\n", method.OutArgs)
				}
				g.printf("},\n")
			}
			g.printf("}\n") // end struct
		}
		g.printf("}\n") // end func
	}
}

func getDBusIdent(imports []*ast.ImportSpec) string {
	for _, spec := range imports {
		if spec.Path.Value == `"github.com/godbus/dbus/v5"` {
			if spec.Name != nil {
				return spec.Name.Name
			}
		}
	}
	return "dbus"
}

func isTargetFunc(funcDecl *ast.FuncDecl, dbusIdent string) bool {
	if !unicode.IsUpper(rune(funcDecl.Name.Name[0])) {
		// 非导出的
		return false
	}
	if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) > 0 {
		resultList := funcDecl.Type.Results.List
		lastField := resultList[len(resultList)-1]
		// 检查最后一个字段的类型是 *dbus.Error
		if starExpr, ok := lastField.Type.(*ast.StarExpr); ok {
			if selExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
				if selExpr.Sel.Name == "Error" {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if ident.Name == dbusIdent {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func getArgs(funcDel *ast.FuncDecl, dbusIdent string) (in, out []string) {
	getArg := func(idx *int, name string, in bool) string {
		if name != "" && name != "_" {
			return name
		}
		dir := "out"
		if in {
			dir = "in"
		}
		name = fmt.Sprintf("%vArg%v", dir, *idx)
		*idx++
		return name
	}

	if funcDel.Type.Params != nil {
		params := funcDel.Type.Params.List
		var argIdx int
		for _, param := range params {
			if selExpr, ok := param.Type.(*ast.SelectorExpr); ok {
				if selExpr.Sel.Name == "Sender" || selExpr.Sel.Name == "Message" {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if ident.Name == dbusIdent {
							// 跳过 dbus.Sender 和 dbus.Message 类型的参数
							continue
						}
					}
				}
			}

			var argName string
			if len(param.Names) > 0 {
				for _, name := range param.Names {
					argName = getArg(&argIdx, name.Name, true)
					in = append(in, argName)
				}
			} else {
				argName = getArg(&argIdx, "", true)
				in = append(in, argName)
			}
		}
	}

	if funcDel.Type.Results != nil {
		results := funcDel.Type.Results.List
		if len(results) > 1 {
			var argIdx int
			for _, result := range results[:len(results)-1] {
				var argName string
				if len(result.Names) > 0 {
					for _, name := range result.Names {
						argName = getArg(&argIdx, name.Name, false)
						out = append(out, argName)
					}
				} else {
					argName = getArg(&argIdx, "", false)
					out = append(out, argName)
				}
			}
		}
	}

	return in, out
}

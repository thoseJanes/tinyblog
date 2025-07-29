package gencode

import (
	"bytes"
	"fmt"
	"go/ast"
	"html/template"
	"strings"

	"go/parser"
	"go/token"
	"strconv"
)


type paramType struct {
	name string
	typ string
}

type functionType struct {
	inputs []paramType
	outputs []paramType
	name string
}

// type implType struct {
// 	implName string
// 	ifaceName string
// 	functions []functionType
// }

type ifaceType struct {
	name string
	functions []functionType
}

type ImportType struct {
	Name string
	Path string
}

type fileInfo struct {
	pkgName string
	importPkgs []ImportType
	ifaces []ifaceType
}

const(
	// tmplString = `
	// {{range $name, $funcs := .}}
	// type {{$name}} struct {
	// {{- range $field := $funcs}}
	// 	{{$field}} string
	// {{- end}}
	// }
	// {{end}}
	// `
	tmplString = `package {{.package}}

import(
{{- range .importPkgs}}
	{{- if gt (len .Name) 0}} 
	{{.Name}} "{{.Path}}"
	{{- else}} 
	"{{.Path}}"
	{{- end}}
{{- end}}
)
{{range $i, $name := .implNames}}
//implement
{{- range $field :=  (index $.implFuncs $i)}}
//{{$field}}
{{- end}}
type {{$name}} struct {

}

var _ {{(index $.ifaceNames $i)}} = (*{{$name}})(nil)


{{range $field :=  (index $.implFuncs $i)}}
func ({{slice $name 0 1}} *{{$name}}) {{$field}} {
	
}

{{end}}
{{end}}
`
)

func GenerateFromInterface(inputPath, outputPath string, mode Mode, reflection map[string]string) error {
	if outputPath == "" {
		outputPath, _ = strings.CutSuffix(inputPath, ".go")
		outputPath += ".gen.go"
	}

	// pkgInfo, err := build.ImportDir(".", 0);
	fileAst, err := getFileInfo(inputPath)
	if err != nil {
		return err
	}
	var implFuncs = [][]string{}
	var ifaceNames = []string{}
	var implNames = []string{}
	for _, iface := range fileAst.ifaces{
		if _, ok := reflection[iface.name]; !ok {
			continue
		}
		var funcParams []string
		for _, function := range iface.functions {
			funcParams = append(funcParams, formatFunctionParams(function))
		}
		implFuncs = append(implFuncs, funcParams)
		ifaceNames = append(ifaceNames, iface.name)
		implNames = append(implNames, reflection[iface.name])
	}

	info := map[string]interface{}{
		"implFuncs":implFuncs,
		"ifaceNames":ifaceNames,
		"implNames":implNames,
		"importPkgs":fileAst.importPkgs,
		"package":fileAst.pkgName,
	}

	tmpl, err := template.New("").Parse(tmplString)
	
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, info)
	if err != nil {
		return err
	}
	
	return WriteToFile(buff.Bytes(), outputPath, mode)
}


func getFileInfo(filePath string) (*fileInfo, error) {
	var fileAst fileInfo
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, filePath, nil, 0)
	if err != nil {
		return nil, err;
	}
	fmt.Printf("%#v\n", filePath)

	interfaceInfo := []ifaceType{}
	packageInfo := []ImportType{}
	ast.Inspect(f, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		if decl.Tok == token.IMPORT {
			for _, spec := range decl.Specs {
				importSpec, ok := spec.(*ast.ImportSpec)
				if !ok {
					continue
				}
				importPath, err := strconv.Unquote(importSpec.Path.Value)
				if err != nil {
					fmt.Print("process import path " + importSpec.Path.Value + " failed." )
					continue
				}
				//importPath = `"` + importPath + `"`
				if importSpec.Name == nil {
					packageInfo = append(packageInfo, ImportType{"", importPath})
				}else{
					packageInfo = append(packageInfo, ImportType{importSpec.Name.Name, importPath})
				}
				
			}
		}

		if decl.Tok == token.TYPE {
			for _, spec := range decl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				iface, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}

				var functions []functionType
				for _, method := range iface.Methods.List {
					funcType, ok := method.Type.(*ast.FuncType)
					if !ok {
						continue
					}
					var inputs[]paramType
					var outputs[]paramType
					for _, param := range funcType.Params.List {
						for _, paramName := range param.Names {
							inputs = append(inputs, paramType{paramName.Name, formatParamType(param.Type)})
							// fmt.Println(nameIdent.X.(*ast.SelectorExpr).)
						}
						if len(param.Names) == 0 {
							inputs = append(inputs, paramType{"", formatParamType(param.Type)})
						}
					}
					for _, param := range funcType.Results.List {
						for _, paramName := range param.Names {
							outputs = append(outputs, paramType{paramName.Name, formatParamType(param.Type)})
						}
						if len(param.Names) == 0 {
							outputs = append(outputs, paramType{"", formatParamType(param.Type)})
						}
					}
					functions = append(functions, functionType{
						inputs: inputs,
						outputs: outputs,
						name: method.Names[0].Name,
					})
					fmt.Printf("method name:%#v\n", method.Names[0].Name)
					//fmt.Printf("type names:%#v\n", funcType.Params)
				}

				interfaceInfo = append(interfaceInfo, 
				ifaceType{
					name: typeSpec.Name.Name, 
					functions: functions,
				})
			}
		}
		return true
	})

	fileAst.pkgName = f.Name.Name
	fileAst.ifaces = interfaceInfo
	fileAst.importPkgs = packageInfo
	return &fileAst, nil
}


func formatParamType(paramType ast.Expr) string {
	if nameIdent, ok := paramType.(*ast.Ident); ok {
		return nameIdent.Name
	}else if pointerType, ok := paramType.(*ast.StarExpr); ok {
		return "*" + formatParamType(pointerType.X)
	}else if selectorType, ok := paramType.(*ast.SelectorExpr); ok {
		return formatParamType(selectorType.X) + "." + selectorType.Sel.Name
	}else if arrayType, ok := paramType.(*ast.ArrayType); ok {
		if arrayType.Len == nil {
			return "[]" + formatParamType(arrayType.Elt)
		}
		return "[" + formatParamType(arrayType.Len) + "]" + formatParamType(arrayType.Elt)
	}else if basicLit, ok := paramType.(*ast.BasicLit); ok {
		return basicLit.Value
	}else{
		fmt.Printf("%#v", paramType)
		nameIdent = paramType.(*ast.Ident)
		
		panic("can't find a conversion")
	}
}

func formatFunctionParams(f functionType) string {
	var inputString string
	var outputString string
	inputString += "("
	for i, param := range f.inputs {
		if param.name == "" {
			inputString += param.typ
		}else{
			inputString += param.name + " " + param.typ
		}
		
		if i != len(f.inputs) - 1 {
			inputString += ", "
		}
	}
	inputString += ")"
	if len(f.outputs) == 1 && f.outputs[0].name == "" {
		outputString = " " + f.outputs[0].typ
	}else if len(f.outputs) != 0{
		outputString += " ("
		for i, param := range f.outputs {
			if param.name == "" {
				outputString += param.typ
			}else{
				outputString += param.name + " " + param.typ
			}
			if i != len(f.outputs) - 1 {
				outputString += ", "
			}
		}
		outputString += ")"
	}

	return f.name + inputString + outputString
}
package generator

import (
	"strings"
	"text/template"
)

var interfaceFuncs = template.FuncMap{
	"ToLower":    strings.ToLower,
	"UnExport":   unexport,
	"Replace":    strings.Replace,
	"IsExported": isExported,
}

const interfaceTemplate string = `// Code generated by counterfeiter. DO NOT EDIT.
package {{.DestinationPackage}}

import (
	{{- range $index, $import := .Imports.ByAlias}}
	{{$import}}
	{{- end}}
)

type {{.Name}} struct {
	{{- range .Methods}}
	{{.Name}}Stub func({{.Params.AsArgs}}) {{.Returns.AsReturnSignature}}
	{{UnExport .Name}}Mutex sync.RWMutex
	{{UnExport .Name}}ArgsForCall []struct{
		{{- range .Params}}
		{{.Name}} {{if .IsVariadic}}{{Replace .Type "..." "[]" -1}}{{else}}{{.Type}}{{end}}
		{{- end}}
	}
	{{- if .Returns.HasLength}}
	{{UnExport .Name}}Returns struct{
		{{- range .Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}
	{{UnExport .Name}}ReturnsOnCall map[int]struct{
		{{- range .Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

{{range .Methods -}}
func (fake *{{$.Name}}) {{.Name}}({{.Params.AsNamedArgsWithTypes}}) {{.Returns.AsReturnSignature}} {
	{{- range .Params.Slices}}
	var {{UnExport .Name}}Copy {{.Type}}
	if {{UnExport .Name}} != nil {
		{{UnExport .Name}}Copy = make({{.Type}}, len({{UnExport .Name}}))
		copy({{UnExport .Name}}Copy, {{UnExport .Name}})
	}
	{{- end}}
	fake.{{UnExport .Name}}Mutex.Lock()
	{{- if .Returns.HasLength}}
	ret, specificReturn := fake.{{UnExport .Name}}ReturnsOnCall[len(fake.{{UnExport .Name}}ArgsForCall)]
	{{- end}}
	fake.{{UnExport .Name}}ArgsForCall = append(fake.{{UnExport .Name}}ArgsForCall, struct{
		{{- range .Params}}
		{{.Name}} {{if .IsVariadic}}{{Replace .Type "..." "[]" -1}}{{else}}{{.Type}}{{end}}
		{{- end}}
	}{ {{- .Params.AsNamedArgs -}} })
	fake.recordInvocation("{{.Name}}", []interface{}{ {{- if .Params.HasLength}}{{.Params.AsNamedArgs}}{{end -}} })
	fake.{{UnExport .Name}}Mutex.Unlock()
	if fake.{{.Name}}Stub != nil {
		{{- if .Returns.HasLength}}
		return fake.{{.Name}}Stub({{.Params.AsNamedArgsForInvocation}}){{else}}fake.{{.Name}}Stub({{.Params.AsNamedArgsForInvocation}})
		{{- end}}
	}
	{{- if .Returns.HasLength}}
	if specificReturn {
		return {{.Returns.WithPrefix "ret."}}
	}
	fakeReturns := fake.{{UnExport .Name}}Returns
	return {{.Returns.WithPrefix "fakeReturns."}}
	{{- end}}
}

func (fake *{{$.Name}}) {{.Name}}CallCount() int {
	fake.{{UnExport .Name}}Mutex.RLock()
	defer fake.{{UnExport .Name}}Mutex.RUnlock()
	return len(fake.{{UnExport .Name}}ArgsForCall)
}

func (fake *{{$.Name}}) {{.Name}}Calls(stub func({{.Params.AsArgs}}) {{.Returns.AsReturnSignature}}) {
	fake.{{UnExport .Name}}Mutex.Lock()
	defer fake.{{UnExport .Name}}Mutex.Unlock()
	fake.{{.Name}}Stub = stub
}

{{if .Params.HasLength -}}
func (fake *{{$.Name}}) {{.Name}}ArgsForCall(i int) {{.Params.AsReturnSignature}} {
	fake.{{UnExport .Name}}Mutex.RLock()
	defer fake.{{UnExport .Name}}Mutex.RUnlock()
	argsForCall := fake.{{UnExport .Name}}ArgsForCall[i]
	return {{.Params.WithPrefix "argsForCall."}}
}
{{- end}}

{{if .Returns.HasLength -}}
func (fake *{{$.Name}}) {{.Name}}Returns({{.Returns.AsNamedArgsWithTypes}}) {
	fake.{{UnExport .Name}}Mutex.Lock()
	defer fake.{{UnExport .Name}}Mutex.Unlock()
	fake.{{.Name}}Stub = nil
	fake.{{UnExport .Name}}Returns = struct {
		{{- range .Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}{ {{- .Returns.AsNamedArgs -}} }
}

func (fake *{{$.Name}}) {{.Name}}ReturnsOnCall(i int, {{.Returns.AsNamedArgsWithTypes}}) {
	fake.{{UnExport .Name}}Mutex.Lock()
	defer fake.{{UnExport .Name}}Mutex.Unlock()
	fake.{{.Name}}Stub = nil
	if fake.{{UnExport .Name}}ReturnsOnCall == nil {
		fake.{{UnExport .Name}}ReturnsOnCall = make(map[int]struct {
			{{- range .Returns}}
			{{UnExport .Name}} {{.Type}}
			{{- end}}
		})
	}
	fake.{{UnExport .Name}}ReturnsOnCall[i] = struct {
		{{- range .Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}{ {{- .Returns.AsNamedArgs -}} }
}

{{end -}}
{{end}}

func (fake *{{.Name}}) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	{{- range .Methods}}
	fake.{{UnExport .Name}}Mutex.RLock()
	defer fake.{{UnExport .Name}}Mutex.RUnlock()
	{{- end}}
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *{{.Name}}) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

{{if IsExported .TargetName -}}
var _ {{.TargetAlias}}.{{.TargetName}} = new({{.Name}})
{{- end}}
`

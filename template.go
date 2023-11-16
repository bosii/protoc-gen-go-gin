package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"html/template"
	"strings"
)

//go:embed template.tpl
var httpTemplate string

type methodDesc struct {
	Name   string // 函数名
	Desc   string // 注释
	Method string // 请求方法
	Path   string // 请求路径
	Req    string // 请求参数名
	Resp   string // 请求返回名

}

type serviceDesc struct {
	ServiceName string
	Methods     []*methodDesc
}

func (s *serviceDesc) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	content := strings.Trim(buf.String(), "\r\n")
	formatted, err := format.Source([]byte(content))
	if err != nil {
		panic(err)
	}
	return string(formatted)
}

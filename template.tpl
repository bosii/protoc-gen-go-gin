type {{.ServiceName}}Handler interface {
     {{range .Methods}}
        {{.Desc}}
        {{.Name}}(r *gin.Engine, req *{{.Req}}, resp *{{.Resp}}) gin.HandlerFunc
    {{- end}}
}

func New{{.ServiceName}} (r *gin.Engine,handler {{.ServiceName}}Handler) {
    {{range .Methods}}
        r.{{.Method}}("{{.Path}}",handler.{{.Name}}(r, &{{.Req}}{}, &{{.Resp}}{}))
    {{- end}}
}

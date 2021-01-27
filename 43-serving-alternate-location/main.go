var t *template.Template
var l = flag.String("location", "http://localhost:8080", "A location")

var tpl = `
<link rel="stylesheet" href="{{.Location}}/styles.css"
`

func servePage(res http.ResponseWriter, req *http.Request) {
	data := struct{ Location *string } {
		Location: 1, 
	}
	t.Execute(res, data)
}

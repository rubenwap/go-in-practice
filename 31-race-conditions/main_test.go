package main

import (
	"testing"
	"bytes"
	"text/template"
)
func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string {
		"Name" : "World",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}

// race detection go test -bench . -race -cpu=1,2,4      

package main

import (
	"fmt"
	"html/template"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"Name":     "Bob",
		"UserName": "bob92",
		"Roles":    []string{"dbteam", "uiteam", "tester"},
	}

	t := template.Must(template.New("email").Parse(emailTmpl))
	builder := &strings.Builder{}
	if err := t.Execute(builder, data); err != nil {
		panic(err)
	}
	s := builder.String()
	fmt.Println(s)

}

const emailTmpl = `Hi {{.Name}}!
Your account is ready, your user name is: {{.UserName}}
You have the following roles assigned:
{{range $i, $r := .Roles}}{{if $i}}, {{end}}{{.}}{{end}}
`

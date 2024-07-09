package render

import (
	"bytes"
	"git.array2d.com/cncf/tca/pkg"
	"text/template"
)

func TextTemplate(tmpl string, kinds map[string]pkg.AnyStruct) (txt string, err error) {
	t := template.New("")
	t, err = t.Parse(tmpl)
	if err != nil {
		return "", err
	}
	sh := bytes.Buffer{}

	err = t.Execute(&sh, kinds)
	if err != nil {
		return "", err
	}
	txt = sh.String()
	return
}

package tca

import (
	"bytes"
	"git.array2d.com/cncf/tca/pkg/shell"
	"text/template"
)

type TemplateTable struct {
	ID     int    `gorm:"primaryKey;`
	Kind   string `gorm:"type:varchar(80)"`
	Method string `gorm:"type:varchar(40)"`
	Shell  string
	Argf1  string
	Sql    string
}

func (tmpl *TemplateTable) BuildAndRunShellArgf(kinds map[string]AnyStruct, in map[string]any) (output string, err error) {
	t, _ := template.New(tmpl.Kind).Parse(tmpl.Shell)
	sh := bytes.Buffer{}
	values := kinds
	values["in"] = in
	err = t.Execute(&sh, values)
	if err != nil {
		return "", err
	}
	var code int
	code, output, err = shell.BashC(sh.String())
	if code == 0 {

	}
	return
}
func (tmpl *TemplateTable) BuildAndRunSql(kinds map[string]AnyStruct, in map[string]string) (output string, err error) {
	return
}

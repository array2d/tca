package tca

import (
	"git.array2d.com/cncf/tca/pkg"
	"git.array2d.com/cncf/tca/pkg/render"
	"git.array2d.com/cncf/tca/pkg/shell"
)

type TemplateTable struct {
	ID     int    `gorm:"primaryKey;`
	Kind   string `gorm:"type:varchar(80)"`
	Method string `gorm:"type:varchar(40)"`
	Shell  string
	Argf1  string
	Sql    string
}

func (tmpl *TemplateTable) BuildAndRunShellArgf(kinds map[string]pkg.AnyStruct, in pkg.AnyStruct) (output string, err error) {
	var sh string
	sh, err = render.TextTemplate(tmpl.Shell, kinds, in)
	var code int
	code, output, err = shell.BashC(sh)
	if code == 0 {

	}
	return
}

package tca

import (
	"git.array2d.com/cncf/tca/pkg"
	"git.array2d.com/cncf/tca/pkg/render"
	"git.array2d.com/cncf/tca/pkg/shell"
	log "github.com/sirupsen/logrus"
	"os"
)

type TemplateTable struct {
	ID     int    `gorm:"primaryKey;`
	Kind   string `gorm:"type:varchar(80)"`
	Method string `gorm:"type:varchar(40)"`
	Shell  string
	File1  string
	Sql    string
}

func (tmpl *TemplateTable) BuildAndRunShellArgf(kinds map[string]pkg.AnyStruct) (code int, output string, err error) {
	var files []*os.File
	var fileTexts []string
	defer func() {
		for _, f := range files {
			if err = f.Close(); err != nil {
				log.WithError(err).Errorln("tmp failed")
			}
		}
	}()

	if tmpl.File1 != "" {
		files = append(files, &os.File{})
		fileTexts = append(fileTexts, "")
		fileTexts[0], err = render.TextTemplate(tmpl.File1, kinds)
		if err != nil {
			log.WithError(err).Errorln("render argf- failed")
			return 500, "", err
		}
		files[0], err = os.Create("/tmp/tca-argf-" + tmpl.Kind + "-" + tmpl.Method + "-0")
		if err != nil {
			log.WithError(err).Errorln("create argf- failed")
			return 500, "", err
		}

		if _, err = files[0].Write([]byte(fileTexts[0])); err != nil {
			log.WithError(err).Errorln("write argf- failed")
			return 500, "", err
		}

	}

	var sh string
	sh, err = render.TextTemplate(tmpl.Shell, kinds)
	log.Debugln(sh)
	code, output, err = shell.BashFile(sh)
	if code == 0 {

	}

	return
}

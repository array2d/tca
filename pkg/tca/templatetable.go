package tca

import (
	"fmt"
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

	shDir := "/var/tmp/tca"
	shDir += string(os.PathSeparator) + tmpl.Kind
	os.MkdirAll(shDir, os.ModePerm)
	if tmpl.File1 != "" {
		fileText, err := render.TextTemplate(tmpl.File1, kinds)
		if err != nil {
			log.WithError(err).Errorln("render argf- failed")
			return 500, "", err
		}
		argfname := fmt.Sprintf("%s.file1", tmpl.Method)
		f, err := os.Create(shDir + string(os.PathSeparator) + argfname)
		if err != nil {
			log.Println("Failed to create arg file:", err)
			return 500, "", err
		}
		if _, err = f.Write([]byte(fileText)); err != nil {
			log.WithError(err).Errorln("write argf- failed")
			return 500, "", err
		}
		files = append(files, f)
		fileTexts = append(fileTexts, fileText)
	}

	var sh string
	sh, err = render.TextTemplate(tmpl.Shell, kinds)

	var shfile *os.File
	shpath := shDir + string(os.PathSeparator) + fmt.Sprintf("%s.sh", tmpl.Method)
	shfile, err = os.Create(shpath)
	if _, err = shfile.Write([]byte(sh)); err != nil {
		log.WithError(err).Errorln("write tmp failed")
	}
	code, output, err = shell.BashFile(shpath)
	if code == 0 {

	}

	return
}

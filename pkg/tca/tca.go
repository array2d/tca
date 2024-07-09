package tca

import (
	"errors"
	"fmt"
	"git.array2d.com/cncf/tca/pkg"
	"git.array2d.com/cncf/tca/pkg/db"
	"git.array2d.com/cncf/tca/pkg/render"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Tca struct {
	db *gorm.DB
}

func New() (t *Tca) {
	t = &Tca{}
	var err error
	t.db, err = db.Getdb()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Fatalln("config error")
		return nil
	}

	models := []any{
		new(TemplateTable),
	}
	err = t.db.Migrator().AutoMigrate(models...)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Fatalln("migrate error")
	}
	return
}

func (t *Tca) TemplateTable(kind, method string) (template TemplateTable, err error) {
	err = t.db.Model(&template).Where("kind = ? and method = ?", kind, method).First(&template).Error
	if err != nil {
		log.WithFields(
			log.Fields{
				"kind":   kind,
				"method": method,
			}).Errorln("template not found")
		err = errors.New("template not found:" + method + " " + kind)
	}
	return
}

func (t *Tca) ComplateKinds(kindsid map[string]string) (kinds map[string]pkg.AnyStruct, err error) {
	kinds = make(map[string]pkg.AnyStruct)
	for kind, id := range kindsid {
		var a pkg.AnyStruct
		err = t.db.Table(kind).Table(" CAST(id AS CHAR)  = ?", id).Find(&a).Error
		if err != nil {
			log.WithFields(
				log.Fields{
					"kind": kind,
					"id":   id,
				}).Errorln("kindsid not found")
			return
		}
		kinds[kind] = a
	}
	return
}

func (t *Tca) Method(kind, method string, kindsid map[string]string, in pkg.AnyStruct) (code int, stdouterr string) {
	tmpl, err := t.TemplateTable(kind, method)
	if err != nil {
		return 500, err.Error()
	}
	kinds, err := t.ComplateKinds(kindsid)
	if err != nil {
		return 500, err.Error()
	}
	var output string
	output, err = tmpl.BuildAndRunShellArgf(kinds, in)
	if err != nil {
		return 500, fmt.Sprint(output)
	}
	if tmpl.Sql != "" {
		var sql string
		sql, err = render.TextTemplate(tmpl.Sql, kinds, in)
		err = t.db.Exec(sql).Error

		if err != nil {
			log.WithFields(
				log.Fields{
					"kind": kind,
					"sql":  tmpl.Sql,
				}).Errorln("sql execute failed")
		}
	}
	return 200, fmt.Sprint(output)
}

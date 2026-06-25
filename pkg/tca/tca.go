package tca

import (
	"encoding/json"
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
		log.WithError(err).WithFields(
			log.Fields{
				"err": err,
			}).Fatalln("migrate error")
	}
	return
}

func (t *Tca) TemplateTable(kind, method string) (template TemplateTable, err error) {
	err = t.db.Model(&template).Where("kind = ? and method = ?", kind, method).First(&template).Error
	if err != nil {
		log.WithError(err).WithFields(
			log.Fields{
				"kind":   kind,
				"method": method,
			}).Errorln("template not found")
		err = errors.New("template not found:" + method + " " + kind)
	}
	return
}

func (t *Tca) CompleteKinds(kindsid map[string]string) (kinds map[string]pkg.AnyStruct, err error) {
	kinds = make(map[string]pkg.AnyStruct)
	for kind, id := range kindsid {
		var as []map[string]interface{}
		err = t.db.Table(kind).Where("CAST(id AS CHAR)  = ?", id).Find(&as).Error
		if err != nil {
			log.WithError(err).WithFields(
				log.Fields{
					"kind": kind,
					"id":   id,
				}).Errorln("kindsid not found")
			return
		}
		if len(as) == 1 {
			kinds[kind] = as[0]
		}
	}
	return
}

func (t *Tca) Method(kind, method string, kindsid map[string]string, in pkg.AnyStruct) (code int, stdouterr string) {
	tmpl, err := t.TemplateTable(kind, method)
	if err != nil {
		return 500, err.Error()
	}
	kinds, err := t.CompleteKinds(kindsid)
	if err != nil {
		return 500, err.Error()
	}
	kinds["in"] = in
	var out = make(map[string]any)
	out["code"], out["output"], out["err"] = tmpl.BuildAndRunShellArgf(kinds)
	if out["err"] != nil {
		return 500, fmt.Sprint(out["output"])
	}
	if tmpl.Sql != "" {
		var sql string
		kinds["out"] = out
		var jsonout pkg.AnyStruct
		jsonerr := json.Unmarshal([]byte(out["output"].(string)), &jsonout)
		if jsonerr == nil {
			out["output"] = jsonout
		} else {
			log.WithError(jsonerr).Errorln("json unmarshal output failed")
		}
		sql, err = render.TextTemplate(tmpl.Sql, kinds)
		if err != nil {
			log.WithError(err).WithFields(
				log.Fields{
					"kind": kind,
					"sql":  tmpl.Sql,
				}).Errorln("sql template failed")
			return 500, err.Error()
		}
		log.Debugln(sql)
		err = t.db.Exec(sql).Error
		if err != nil {
			log.WithError(err).WithFields(
				log.Fields{
					"kind": kind,
					"sql":  tmpl.Sql,
				}).Errorln("sql execute failed")
			return 500, err.Error()
		}
	}
	return 200, fmt.Sprint(out["output"])
}

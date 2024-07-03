package tca

import (
	"errors"
	"git.array2d.com/cncf/tca/pkg/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

var (
	templateTable string = "template"
)

func init() {
	t := os.Getenv("TEMPLATE_TABLE")
	if t != "" {
		templateTable = t
		log.WithFields(
			log.Fields{
				"TEMPLATE_TABLE": t,
			}).Infoln("templateTable updated")
	}
}

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
	t.db.Migrator().AutoMigrate(new(TemplateTable))
	return
}
func (t *Tca) Method(kind, method, id string, values map[string]any) (code int, stdouterr string) {
	var err error
	var template TemplateTable
	err = t.db.Table(templateTable).Where("kind = ? and method = ?", kind, method).Find(&template).Error
	if err != nil {
		log.WithFields(
			log.Fields{
				"kind":   kind,
				"method": method,
			}).Errorln("template not found")
		err = errors.New("template not found:" + method + " " + kind)
		code = 500
		stdouterr = err.Error()
	}

	if id == "" {
		err = t.db.Table(kind).Create(values).Error
		if err != nil {
			log.WithFields(
				log.Fields{
					"kind":   kind,
					"method": method,
				}).Errorln("id=null ")
			err = errors.New("id=null " + method + " " + kind)
			code = 500
			stdouterr = err.Error()
		}
	} else {
		err = t.db.Table(kind).Find(values).Error
		if err != nil {
			log.WithFields(
				log.Fields{
					"kind":   kind,
					"method": method,
					"id":     id,
				}).Errorln("select failed")
			err = errors.New("select " + kind + "'s value failed" + method + " " + id)
			code = 500
			stdouterr = err.Error()
		}
	}

	return
}

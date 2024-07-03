package tca

type TemplateTable struct {
	ID     int    `gorm:"primaryKey;`
	Kind   string `gorm:"type:varchar(80)"`
	Method string `gorm:"type:varchar(40)"`
	Shell  string
	Argf1  string
}

//var (
//	templateTable string = "template"
//)
//
//func init() {
//	t := os.Getenv("TEMPLATE_TABLE")
//	if t != "" {
//		templateTable = t
//		log.WithFields(
//			log.Fields{
//				"TEMPLATE_TABLE": t,
//			}).Infoln("templateTable updated")
//	}
//}

package tca

type TemplateTable struct {
	ID     int `gorm:"primaryKey;`
	Kind   string
	Method string
	Shell  string
	Argf1  string
}

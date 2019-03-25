package models

type Goods struct {
	Id_RENAME int    `orm:"column(id);null"`
	Name      string `orm:"column(name);size(128)"`
	Image     string `orm:"column(image);size(128)"`
}

// Code generated by go-gendb, DO NOT EDIT.
// go-gendb version: 0.3
// source: samples/orm-mgo/item.go
package item

const (
	SkuAttrFieldPropId        = "prop_id"
	SkuAttrFieldPropValueId   = "prop_value_id"
	SkuAttrFieldImg           = "img"
	SkuAttrFieldPropName      = "prop_name"
	SkuAttrFieldPropValueName = "prop_value_name"
)

type SkuAttr struct {
	PropId        string            `bson:"prop_id" json:"prop_id"`
	PropValueId   string            `bson:"prop_value_id" json:"prop_value_id"`
	Img           string            `bson:"img" json:"img"`
	PropName      map[string]string `bson:"prop_name" json:"prop_name"`
	PropValueName map[string]string `bson:"prop_value_name" json:"prop_value_name"`
}
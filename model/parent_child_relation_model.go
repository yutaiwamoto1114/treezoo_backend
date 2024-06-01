// /model/parent_child_relation_model.go
package model

type ParentChildRelation struct {
	ChildId  int `json:"child_id"`
	ParentId int `json:"parent_id"`
}

package model

type ParentChildRelation struct {
	ChildId  int `json:"child_id"`
	ParentId int `json:"parent_id"`
}

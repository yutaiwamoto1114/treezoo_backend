// /model/partner_relation_model.go
package model

type PartnerRelation struct {
	PartnerRelationId int    `json:"partner_relation_id"`
	Animal1Id         int    `json:"animal1_id"`
	Animal2Id         int    `json:"animal2_id"`
	Status            string `json:"status"`
	Notes             string `json:"notes"`
}

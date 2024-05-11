// lib/model/picture_model.go
package model

import "database/sql"

// 動物プロフィール写真モデル
type AnimalProfilePicture struct {
	AnimalId  int `json:"animal_id"`
	PictureId int `json:"picture_id"`
	// PictureData []byte         `json:"picture_data"`
	PictureData string         `json:"picture_data"`
	Description sql.NullString `json:"description"`
	IsProfile   sql.NullBool   `json:"is_profile"`
	Priority    sql.NullInt64  `json:"priority"`
	ShootDate   sql.NullTime   `json:"shoot_date"`
	UploadDate  sql.NullTime   `json:"upload_date"`
}

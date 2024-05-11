package model

import (
	"database/sql"
)

// 写真モデル: 画像のバイナリデータとメタデータを保持する
type Picture struct {
	PictureId   int            `json:"picture_id"`
	PictureData []byte         `json:"picture_data"`
	Description sql.NullString `json:"description"`
	IsProfile   sql.NullBool   `json:"is_profile"`
	Priority    sql.NullInt64  `json:"priority"`
	ShootDate   sql.NullTime   `json:"shoot_date"`
	UploadDate  sql.NullTime   `json:"upload_date"`
}

CREATE TABLE `pictures` (
  `picture_id` int NOT NULL AUTO_INCREMENT,
  `picture_data` mediumblob NOT NULL COMMENT '写真データ',
  `description` varchar(255) DEFAULT NULL COMMENT '説明',
  `is_profile` tinyint(1) DEFAULT '0' COMMENT 'プロフィール写真フラグ',
  `priority` int DEFAULT NULL COMMENT '表示優先度',
  `shoot_date` date DEFAULT NULL COMMENT '撮影日',
  `upload_date` date DEFAULT NULL COMMENT 'アップロード日',
  PRIMARY KEY (`picture_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
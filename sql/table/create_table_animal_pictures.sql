CREATE TABLE `animal_pictures` (
  `animal_id` int NOT NULL,
  `picture_id` int NOT NULL,
  `is_profile` tinyint(1) DEFAULT '0' COMMENT 'プロフィール写真フラグ',
  PRIMARY KEY (`animal_id`,`picture_id`),
  KEY `fk_animal_pictures_picture` (`picture_id`),
  CONSTRAINT `fk_animal_pictures_animal` FOREIGN KEY (`animal_id`) REFERENCES `animals` (`animal_id`),
  CONSTRAINT `fk_animal_pictures_picture` FOREIGN KEY (`picture_id`) REFERENCES `pictures` (`picture_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
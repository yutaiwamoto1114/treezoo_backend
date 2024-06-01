CREATE TABLE `zoo_pictures` (
  `zoo_id` int NOT NULL,
  `picture_id` int NOT NULL,
  PRIMARY KEY (`zoo_id`,`picture_id`),
  KEY `fk_zoo_pictures_picture` (`picture_id`),
  CONSTRAINT `fk_zoo_pictures_picture` FOREIGN KEY (`picture_id`) REFERENCES `pictures` (`picture_id`),
  CONSTRAINT `fk_zoo_pictures_zoo` FOREIGN KEY (`zoo_id`) REFERENCES `zoos` (`zoo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
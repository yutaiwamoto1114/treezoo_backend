CREATE TABLE `animals` (
  `animal_id` int NOT NULL AUTO_INCREMENT COMMENT '動物ID',
  `name` varchar(255) NOT NULL COMMENT '名前',
  `species` varchar(255) NOT NULL COMMENT '種類',
  `birth_zoo_id` int NOT NULL COMMENT '生まれた動物園ID',
  `current_zoo_id` int NOT NULL COMMENT '現在いる動物園ID',
  `birthday` date DEFAULT NULL COMMENT '誕生日',
  `age` int DEFAULT NULL COMMENT '年齢',
  `gender` varchar(10) DEFAULT NULL COMMENT '性別',
  PRIMARY KEY (`animal_id`),
  KEY `birth_zoo_id` (`birth_zoo_id`),
  KEY `current_zoo_id` (`current_zoo_id`),
  CONSTRAINT `animals_ibfk_1` FOREIGN KEY (`birth_zoo_id`) REFERENCES `zoos` (`zoo_id`),
  CONSTRAINT `animals_ibfk_2` FOREIGN KEY (`current_zoo_id`) REFERENCES `zoos` (`zoo_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
CREATE TABLE `zoos` (
  `zoo_id` int NOT NULL AUTO_INCREMENT COMMENT '動物園ID',
  `name` varchar(255) NOT NULL COMMENT '名前',
  `location` varchar(255) NOT NULL COMMENT '場所',
  PRIMARY KEY (`zoo_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
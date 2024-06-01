CREATE TABLE `partner_relations` (
    `partner_relation_id` int NOT NULL AUTO_INCREMENT,
    `animal1_id` int NOT NULL,
    `animal2_id` int NOT NULL,
    `status` varchar(50) NOT NULL,
    `notes` text,
    PRIMARY KEY (`partner_relation_id`),
    KEY `animal1_id` (`animal1_id`),
    KEY `animal2_id` (`animal2_id`),
    CONSTRAINT `partner_relations_ibfk_1` FOREIGN KEY (`animal1_id`) REFERENCES `animals` (`animal_id`),
    CONSTRAINT `partner_relations_ibfk_2` FOREIGN KEY (`animal2_id`) REFERENCES `animals` (`animal_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
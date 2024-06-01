CREATE TABLE `parent_child_relations` (
  `child_id` int NOT NULL COMMENT '子供ID',
  `parent_id` int NOT NULL COMMENT '親ID',
  `breeding_zoo_id` int NOT NULL COMMENT '配偶が行われた動物園ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC
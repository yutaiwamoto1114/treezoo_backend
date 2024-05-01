CREATE OR REPLACE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `view_parent_child_relations` AS
select `child`.`animal_id` AS `child_id`,
    `child`.`name` AS `child_name`,
    `child`.`species` AS `child_species`,
    `parent`.`animal_id` AS `parent_id`,
    `parent`.`name` AS `parent_name`,
    `parent`.`species` AS `parent_species`,
    `zoos`.`name` AS `zoo_name`,
    `zoos`.`location` AS `zoo_location`
from (
        (
            (
                `parent_child_relations`
                join `animals` `child` on(
                    (
                        `parent_child_relations`.`child_id` = `child`.`animal_id`
                    )
                )
            )
            join `animals` `parent` on(
                (
                    `parent_child_relations`.`parent_id` = `parent`.`animal_id`
                )
            )
        )
        join `zoos` on(
            (
                `parent_child_relations`.`breeding_zoo_id` = `zoos`.`zoo_id`
            )
        )
    )
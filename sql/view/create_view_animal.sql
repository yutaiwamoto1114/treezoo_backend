CREATE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `view_animals` AS
select `a`.`animal_id` AS `animal_id`,
    `a`.`name` AS `animal_name`,
    `a`.`species` AS `species`,
    `a`.`birthday` AS `birthday`,
    `a`.`age` AS `age`,
    `a`.`gender` AS `gender`,
    `b`.`zoo_id` AS `birth_zoo_id`,
    `b`.`name` AS `birth_zoo_name`,
    `b`.`location` AS `birth_zoo_location`,
    `c`.`zoo_id` AS `current_zoo_id`,
    `c`.`name` AS `current_zoo_name`,
    `c`.`location` AS `current_zoo_location`
from (
        (
            `animals` `a`
            left join `zoos` `b` on((`a`.`birth_zoo_id` = `b`.`zoo_id`))
        )
        left join `zoos` `c` on((`a`.`current_zoo_id` = `c`.`zoo_id`))
    )
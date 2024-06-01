CREATE OR REPLACE ALGORITHM = UNDEFINED DEFINER = `root` @`localhost` SQL SECURITY DEFINER VIEW `view_animal_pictures` AS
select `a`.`animal_id` AS `animal_id`,
    `p`.`picture_id` AS `picture_id`,
    `p`.`picture_data` AS `picture_data`,
    `p`.`description` AS `description`,
    `ap`.`is_profile` AS `is_profile`,
    `p`.`priority` AS `priority`,
    `p`.`shoot_date` AS `shoot_date`,
    `p`.`upload_date` AS `upload_date`
from (
        (
            `animals` `a`
            join `animal_pictures` `ap` on((`a`.`animal_id` = `ap`.`animal_id`))
        )
        join `pictures` `p` on((`ap`.`picture_id` = `p`.`picture_id`))
    )
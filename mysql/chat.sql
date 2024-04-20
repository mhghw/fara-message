CREATE TABLE `messenger`.`chat` (
  `chat_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `chat_name` VARCHAR(45) NOT NULL,
  UNIQUE INDEX `chat_id_UNIQUE` (`chat_id` ASC) VISIBLE,
  PRIMARY KEY (`chat_id`));

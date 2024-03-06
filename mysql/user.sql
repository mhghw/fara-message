
CREATE TABLE `user`.`users` (
  `User_ID` int unsigned NOT NULL AUTO_INCREMENT,
  `Username` varchar(45) NOT NULL,
  `Firstname` varchar(45) NOT NULL,
  `Lastname` varchar(45) NOT NULL,
  `Password` varchar(45) NOT NULL,
  `Gender` varchar(45) NOT NULL,
  `Date_Of_Birth` datetime NOT NULL,
  `Created_Time` datetime NOT NULL,
  PRIMARY KEY (`User_ID`),
  UNIQUE KEY `User_ID_UNIQUE` (`User_ID`)
)
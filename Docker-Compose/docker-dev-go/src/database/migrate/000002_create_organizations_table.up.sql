CREATE TABLE IF NOT EXISTS `organizations` (
  `org_id`    int(11) NOT NULL AUTO_INCREMENT,
  `company`   VARCHAR(191) NULL,
  `address`   VARCHAR(191) NULL,
  `phone`     VARCHAR(191) NULL,
  PRIMARY KEY (`org_id`)
);
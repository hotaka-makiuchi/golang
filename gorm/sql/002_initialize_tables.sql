CREATE TABLE `gormsample`.`events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `summary` varchar(255) DEFAULT NULL,
  `start_date` timestamp DEFAULT NULL,
  `end_date` timestamp DEFAULT NULL,
  `count` int default 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `gormsample`.`events` (`summary`, `start_date`, `end_date`) values ('hoge1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
insert into `gormsample`.`events` (`summary`, `start_date`, `end_date`) values ('hoge2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
insert into `gormsample`.`events` (`summary`, `start_date`, `end_date`) values ('hoge3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
insert into `gormsample`.`events` (`summary`, `start_date`, `end_date`) values ('hoge4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
insert into `gormsample`.`events` (`summary`, `start_date`, `end_date`) values ('hoge5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

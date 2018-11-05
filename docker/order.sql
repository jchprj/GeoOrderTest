CREATE TABLE if not exists `order` (
  `order_i_d` bigint(20) NOT NULL AUTO_INCREMENT,
  `distance` int(11) DEFAULT NULL,
  `status` varchar(200) DEFAULT NULL,
  `start_latitude` varchar(200) DEFAULT NULL,
  `start_longitude` varchar(200) DEFAULT NULL,
  `end_latitude` varchar(200) DEFAULT NULL,
  `end_longitude` varchar(200) DEFAULT NULL,
  `createTime` datetime DEFAULT NULL,
  `takenTime` datetime DEFAULT NULL,
  PRIMARY KEY (`order_i_d`),
  KEY `status` (`status`) USING BTREE,
  KEY `createTime` (`createTime`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


--
-- MariaDB
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET latin1 DEFAULT NULL,
  `email` varchar(50) CHARACTER SET latin1 DEFAULT NULL,
  `pwd` varchar(50) CHARACTER SET latin1 DEFAULT NULL,
  `crtTm` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;


--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
INSERT INTO `user` VALUES (1,'wenj1991',NULL,'654321',NULL),(2,'wenj1991','sss@qq.com','12345678','2019-08-26 17:41:04'),(4,'wenj1991',NULL,NULL,NULL),(5,'wenj1991',NULL,NULL,NULL),(6,'wenj1991',NULL,NULL,NULL),(7,'wenj1991',NULL,NULL,NULL),(8,'wenj1991',NULL,NULL,NULL);
UNLOCK TABLES;


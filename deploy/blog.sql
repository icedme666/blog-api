# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.5.13-MariaDB)
# Database: blog
# Generation Time: 2022-08-31 13:20:59 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table blog_article
# ------------------------------------------------------------

DROP TABLE IF EXISTS `blog_article`;

CREATE TABLE `blog_article` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `tag_id` int(10) unsigned DEFAULT 0 COMMENT '标签ID',
                                `title` varchar(100) DEFAULT '' COMMENT '文章标题',
                                `desc` varchar(255) DEFAULT '' COMMENT '简述',
                                `content` text DEFAULT NULL,
                                `created_on` int(11) DEFAULT NULL,
                                `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
                                `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int(10) unsigned DEFAULT 0,
                                `state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态 0为禁用1为启用',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';



# Dump of table blog_auth
# ------------------------------------------------------------

DROP TABLE IF EXISTS `blog_auth`;

CREATE TABLE `blog_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `blog_auth` WRITE;
/*!40000 ALTER TABLE `blog_auth` DISABLE KEYS */;

INSERT INTO `blog_auth` (`id`, `username`, `password`)
VALUES
(1,'test','test123456');

/*!40000 ALTER TABLE `blog_auth` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table blog_tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `blog_tag`;

CREATE TABLE `blog_tag` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(100) DEFAULT '' COMMENT '标签名称',
                            `created_on` int(10) unsigned DEFAULT 0 COMMENT '创建时间',
                            `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                            `modified_on` int(10) unsigned DEFAULT 0 COMMENT '修改时间',
                            `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                            `deleted_on` int(10) unsigned DEFAULT 0,
                            `state` tinyint(3) unsigned DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

/*
Navicat MySQL Data Transfer

Source Server         : kabao
Source Server Version : 50549
Source Host           : 172.98.201.182:3306
Source Database       : shadowsocks

Target Server Type    : MYSQL
Target Server Version : 50549
File Encoding         : 65001

Date: 2016-07-28 23:13:59
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `serverid` int(11) NOT NULL,
  `containerid` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `password` varchar(255) NOT NULL,
  `userid` int(11) NOT NULL,
  `createtime` datetime DEFAULT NULL,
  `expiretime` datetime DEFAULT NULL,
  `cycle` tinyint(1) NOT NULL DEFAULT '0',
  `active` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1正常，2suspended 3 deleted',
  `firstprice` decimal(12,2) DEFAULT NULL,
  `recurringprice` decimal(12,2) NOT NULL DEFAULT '20.00',
  PRIMARY KEY (`id`),
  KEY `server` (`serverid`),
  KEY `user` (`userid`),
  CONSTRAINT `server` FOREIGN KEY (`serverid`) REFERENCES `server` (`id`),
  CONSTRAINT `user` FOREIGN KEY (`userid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for bill
-- ----------------------------
DROP TABLE IF EXISTS `bill`;
CREATE TABLE `bill` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountid` int(11) NOT NULL,
  `price` decimal(12,2) NOT NULL,
  `createtime` datetime NOT NULL,
  `expiretime` datetime NOT NULL,
  `ispaid` tinyint(1) NOT NULL DEFAULT '0',
  `active` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1有效0无效',
  `payno` varchar(25) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `acc` (`accountid`)
) ENGINE=MyISAM AUTO_INCREMENT=99 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `auth` varchar(255) NOT NULL,
  `driver` varchar(32) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `month` decimal(12,2) DEFAULT '20.00',
  `quarter` decimal(12,2) DEFAULT NULL,
  `year` decimal(12,2) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `have` int(11) DEFAULT NULL,
  `isonline` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `firstname` varchar(32) DEFAULT NULL,
  `lastname` varchar(32) DEFAULT NULL,
  `company` varchar(32) DEFAULT NULL,
  `address1` varchar(64) DEFAULT NULL,
  `address2` varchar(64) DEFAULT NULL,
  `city` varchar(64) DEFAULT NULL,
  `province` varchar(64) DEFAULT NULL,
  `zipcode` varchar(32) DEFAULT NULL,
  `country` varchar(32) DEFAULT NULL,
  `mobile` varchar(64) DEFAULT NULL,
  `issubscribe` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

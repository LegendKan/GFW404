/*
Navicat MySQL Data Transfer

Source Server         : BWG
Source Server Version : 50173
Source Host           : 107.182.177.241:3306
Source Database       : shadowsocks

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2015-12-24 00:05:07
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
  `active` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `server` (`serverid`),
  KEY `user` (`userid`),
  CONSTRAINT `server` FOREIGN KEY (`serverid`) REFERENCES `server` (`id`),
  CONSTRAINT `user` FOREIGN KEY (`userid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES ('11', '1', '', '0', '', '1', '2015-01-29 05:40:20', null, '1', '0');
INSERT INTO `account` VALUES ('12', '1', '', '0', '', '1', '2015-01-29 05:42:21', null, '1', '0');
INSERT INTO `account` VALUES ('13', '1', '', '0', '', '1', '2015-01-29 05:42:23', null, '3', '0');
INSERT INTO `account` VALUES ('14', '1', '', '0', '', '1', '2015-01-29 05:46:13', null, '3', '0');
INSERT INTO `account` VALUES ('15', '1', '', '0', '', '1', '2015-01-31 09:02:35', '2015-03-03 09:02:35', '1', '0');
INSERT INTO `account` VALUES ('16', '1', '', '0', '', '1', '2015-01-31 09:06:54', '2015-05-01 08:06:54', '2', '0');
INSERT INTO `account` VALUES ('17', '1', '', '0', '', '1', '2015-01-31 09:11:04', '2015-03-03 09:11:04', '1', '0');
INSERT INTO `account` VALUES ('18', '1', '', '0', '', '1', '2015-01-31 09:12:12', '2015-05-01 08:12:12', '2', '0');
INSERT INTO `account` VALUES ('19', '1', '', '0', '', '1', '2015-01-31 09:22:58', '2016-01-31 09:22:58', '3', '0');
INSERT INTO `account` VALUES ('20', '1', '', '0', '', '1', '2015-01-31 14:35:06', '2015-05-01 12:35:06', '2', '0');
INSERT INTO `account` VALUES ('21', '1', '', '0', '', '1', '2015-01-31 14:39:36', '2015-05-01 12:39:36', '2', '0');
INSERT INTO `account` VALUES ('22', '1', '', '0', '', '1', '2015-01-31 14:50:35', '2015-05-01 12:50:35', '2', '0');
INSERT INTO `account` VALUES ('23', '1', '', '0', '', '1', '2015-01-31 15:01:36', '2015-05-01 13:01:36', '2', '0');
INSERT INTO `account` VALUES ('24', '1', '', '0', '', '1', '2015-01-31 10:18:18', '2015-05-01 09:18:18', '2', '0');
INSERT INTO `account` VALUES ('25', '1', '', '0', '', '1', '2015-01-31 10:27:22', '2015-05-01 09:27:22', '2', '0');
INSERT INTO `account` VALUES ('26', '1', '', '0', '', '1', '2015-01-31 10:30:47', '2015-05-01 09:30:47', '2', '0');
INSERT INTO `account` VALUES ('27', '1', '', '0', '', '1', '2015-01-31 15:33:45', '2015-05-01 13:33:45', '2', '0');
INSERT INTO `account` VALUES ('28', '1', '', '0', '', '1', '2015-01-31 14:20:54', '2015-03-03 14:20:54', '1', '0');
INSERT INTO `account` VALUES ('29', '1', '', '0', '', '1', '2015-01-31 19:21:52', '2015-03-03 19:21:52', '1', '1');
INSERT INTO `account` VALUES ('30', '1', '', '0', '', '1', '2015-01-31 19:32:02', '2015-03-03 19:32:02', '1', '0');
INSERT INTO `account` VALUES ('31', '1', '29168792929a63c0de6acb49181f9b20f790569c12564e7ce52b28737cb0037f', '49158', 'hTHctcuA', '1', '2015-01-31 19:35:38', '2015-03-03 19:35:38', '1', '1');
INSERT INTO `account` VALUES ('32', '1', '', '0', '', '1', '2015-12-22 06:11:21', '2016-01-22 06:11:21', '1', '0');
INSERT INTO `account` VALUES ('33', '1', '', '0', '', '1', '2015-12-22 06:48:17', '2016-01-22 06:48:17', '1', '0');
INSERT INTO `account` VALUES ('34', '1', '', '0', '', '1', '2015-12-22 14:18:34', '2016-01-22 14:18:34', '1', '0');
INSERT INTO `account` VALUES ('35', '1', '', '0', '', '1', '2015-12-22 14:21:11', '2016-01-22 14:21:11', '1', '0');
INSERT INTO `account` VALUES ('36', '1', '', '0', '', '1', '2015-12-22 14:22:58', '2016-01-22 14:22:58', '1', '0');
INSERT INTO `account` VALUES ('37', '1', '', '0', '', '1', '2015-12-22 14:42:25', '2016-01-22 14:42:25', '1', '0');
INSERT INTO `account` VALUES ('38', '1', '', '0', '', '1', '2015-12-22 14:44:54', '2016-01-22 14:44:54', '1', '0');
INSERT INTO `account` VALUES ('39', '1', '', '0', '', '1', '2015-12-22 14:45:56', '2016-01-22 14:45:56', '1', '0');
INSERT INTO `account` VALUES ('40', '1', '', '0', '', '1', '2015-12-22 14:56:22', '2016-01-22 14:56:22', '1', '0');
INSERT INTO `account` VALUES ('41', '1', '', '0', '', '1', '2015-12-23 09:32:03', '2016-01-23 09:32:03', '1', '0');
INSERT INTO `account` VALUES ('42', '1', '', '0', '', '1', '2015-12-23 10:35:50', '2016-01-23 10:35:50', '1', '0');
INSERT INTO `account` VALUES ('43', '1', '', '0', '', '1', '2015-12-23 10:43:58', '2016-01-23 10:43:58', '1', '0');
INSERT INTO `account` VALUES ('44', '1', '', '0', '', '1', '2015-12-23 11:08:10', '2016-01-23 11:08:10', '1', '0');
INSERT INTO `account` VALUES ('45', '1', '', '0', '', '1', '2015-12-23 11:10:17', '2016-01-23 11:10:17', '1', '0');
INSERT INTO `account` VALUES ('46', '1', '', '0', '', '1', '2015-12-23 11:12:51', '2016-01-23 11:12:51', '1', '0');
INSERT INTO `account` VALUES ('47', '1', '', '0', '', '1', '2015-12-23 11:14:51', '2016-01-23 11:14:51', '1', '0');
INSERT INTO `account` VALUES ('48', '1', '', '0', '', '1', '2015-12-23 13:16:11', '2016-01-23 13:16:11', '1', '0');

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
  `payno` varchar(25) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `acc` (`accountid`)
) ENGINE=MyISAM AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bill
-- ----------------------------
INSERT INTO `bill` VALUES ('9', '11', '20.00', '2015-01-29 05:40:20', '0000-00-00 00:00:00', '0', '1422510020');
INSERT INTO `bill` VALUES ('10', '12', '20.00', '2015-01-29 05:42:22', '0000-00-00 00:00:00', '0', '1422510141');
INSERT INTO `bill` VALUES ('11', '13', '100.00', '2015-01-29 05:42:23', '0000-00-00 00:00:00', '0', '1422510141');
INSERT INTO `bill` VALUES ('12', '14', '100.00', '2015-01-29 05:46:14', '0000-00-00 00:00:00', '0', '1422510373');
INSERT INTO `bill` VALUES ('13', '15', '20.00', '2015-01-31 09:02:35', '2015-02-05 09:02:35', '0', '1422694955');
INSERT INTO `bill` VALUES ('14', '16', '40.00', '2015-01-31 09:06:54', '2015-02-05 09:06:54', '0', '1422695214');
INSERT INTO `bill` VALUES ('15', '17', '20.00', '2015-01-31 09:11:04', '2015-02-05 09:11:04', '0', '1422695464');
INSERT INTO `bill` VALUES ('16', '18', '40.00', '2015-01-31 09:12:12', '2015-02-05 09:12:12', '0', '1422695532');
INSERT INTO `bill` VALUES ('17', '19', '100.00', '2015-01-31 09:22:58', '2015-02-05 09:22:58', '0', '1422696178');
INSERT INTO `bill` VALUES ('18', '20', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('19', '21', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('20', '22', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('21', '23', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('22', '24', '40.00', '2015-01-31 10:18:18', '2015-02-05 10:18:18', '0', '1422699498');
INSERT INTO `bill` VALUES ('23', '25', '40.00', '2015-01-31 10:27:22', '2015-02-05 10:27:22', '0', '1422700042');
INSERT INTO `bill` VALUES ('24', '26', '40.00', '2015-01-31 10:30:47', '2015-02-05 10:30:47', '0', '1422700247');
INSERT INTO `bill` VALUES ('25', '27', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('26', '28', '20.00', '2015-01-31 14:20:54', '2015-02-05 14:20:54', '0', '1422714054');
INSERT INTO `bill` VALUES ('27', '29', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('28', '30', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('29', '31', '0.00', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '1', '');
INSERT INTO `bill` VALUES ('30', '32', '20.00', '2015-12-22 06:11:21', '2015-12-27 06:11:21', '0', '1450764681');
INSERT INTO `bill` VALUES ('31', '33', '20.00', '2015-12-22 06:48:17', '2015-12-27 06:48:17', '0', '1450766897');
INSERT INTO `bill` VALUES ('32', '34', '20.00', '2015-12-22 14:18:35', '2015-12-27 14:18:34', '0', '1450793914');
INSERT INTO `bill` VALUES ('33', '35', '20.00', '2015-12-22 14:21:11', '2015-12-27 14:21:11', '0', '1450794071');
INSERT INTO `bill` VALUES ('34', '36', '20.00', '2015-12-22 14:22:59', '2015-12-27 14:22:58', '0', '1450794178');
INSERT INTO `bill` VALUES ('35', '37', '20.00', '2015-12-22 14:42:25', '2015-12-27 14:42:25', '0', '1450795345');
INSERT INTO `bill` VALUES ('36', '38', '20.00', '2015-12-22 14:44:54', '2015-12-27 14:44:54', '0', '1450795494');
INSERT INTO `bill` VALUES ('37', '39', '20.00', '2015-12-22 14:45:56', '2015-12-27 14:45:56', '0', '1450795556');
INSERT INTO `bill` VALUES ('38', '40', '20.00', '2015-12-22 14:56:22', '2015-12-27 14:56:22', '0', '1450796182');
INSERT INTO `bill` VALUES ('39', '41', '20.00', '2015-12-23 09:32:03', '2015-12-28 09:32:03', '0', '1450863123');
INSERT INTO `bill` VALUES ('40', '42', '20.00', '2015-12-23 10:35:51', '2015-12-28 10:35:50', '0', '1450866950');
INSERT INTO `bill` VALUES ('41', '43', '20.00', '2015-12-23 10:43:58', '2015-12-28 10:43:58', '0', '1450867438');
INSERT INTO `bill` VALUES ('42', '44', '20.00', '2015-12-23 11:08:11', '2015-12-28 11:08:10', '0', '1450868890');
INSERT INTO `bill` VALUES ('43', '45', '20.00', '2015-12-23 11:10:17', '2015-12-28 11:10:17', '0', '1450869017');
INSERT INTO `bill` VALUES ('44', '46', '20.00', '2015-12-23 11:12:52', '2015-12-28 11:12:51', '0', '1450869171');
INSERT INTO `bill` VALUES ('45', '47', '20.00', '2015-12-23 11:14:52', '2015-12-28 11:14:51', '0', '1450869291');
INSERT INTO `bill` VALUES ('46', '48', '20.00', '2015-12-23 13:16:12', '2015-12-28 13:16:11', '0', '1450876571');

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `auth` varchar(255) NOT NULL,
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of server
-- ----------------------------
INSERT INTO `server` VALUES ('1', '198.12.87.172', '8080', 'LoveWYN1008', 'Seattle', 'Shadowsocks.com 普通版第四期', '3 个美国节点，分别位于 Fremont 和 Los Angeles，高带宽|4 个亚洲节点，分别位于新加坡和日本，低延迟|不限制流量|同一时间同一个账号仅限一个终端使用，如需多个账号请联系销售部门获取优惠。', '20.00', '40.00', '100.00', '20', '0', '1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '看传奇', '19901124', 'surpesb@163.com');
INSERT INTO `user` VALUES ('2', 'helloworld', '123456', '123@123.com');
INSERT INTO `user` VALUES ('3', 'lelekan', 'kanchuanqi', '981088636@qq.com');

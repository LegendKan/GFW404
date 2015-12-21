/*
Navicat MySQL Data Transfer

Source Server         : BWG
Source Server Version : 50173
Source Host           : 107.182.177.241:3306
Source Database       : shadowsocks

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2015-12-21 10:51:58
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;

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
) ENGINE=MyISAM AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;

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
  `remain` int(11) DEFAULT NULL,
  `isonline` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of server
-- ----------------------------
INSERT INTO `server` VALUES ('1', '198.12.87.172', '8080', 'LoveWYN1008', 'Seattle', 'Shadowsocks.com 普通版第四期', '3 个美国节点，分别位于 Fremont 和 Los Angeles，高带宽|4 个亚洲节点，分别位于新加坡和日本，低延迟|不限制流量|同一时间同一个账号仅限一个终端使用，如需多个账号请联系销售部门获取优惠。', '20.00', '40.00', '100.00', '20', '20', '1');
INSERT INTO `server` VALUES ('2', '19812', '8081', '', 'Seattle', 'Shadowsocks.com 高级版', '可用所有普通版的节点，另增加 VIP 节点：|香港 Rackspace 机房|不限制流量|同一个账号同时支持 5 个终端使用（不支持分享帐号给多人使用），如需更多帐号请联系销售部门获取优惠。', '0.00', '0.00', '0.00', '20', '0', '1');

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

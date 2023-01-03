/*
 Navicat MySQL Data Transfer

 Source Server         : ucloud
 Source Server Type    : MySQL
 Source Server Version : 50740 (5.7.40)
 Source Host           : 106.75.13.27:3307
 Source Schema         : asm

 Target Server Type    : MySQL
 Target Server Version : 50740 (5.7.40)
 File Encoding         : 65001

 Date: 27/12/2022 17:09:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sub_domain_items
-- ----------------------------
DROP TABLE IF EXISTS `sub_domain_items`;
CREATE TABLE `sub_domain_items`  (
  `domain` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `alive` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `cdn` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `asn` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `addr` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `isp` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `source` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `title` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `reason` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `ip` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `url` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `cname` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `subdomain` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `port` int(11) NULL DEFAULT NULL,
  `Favicon Hash` bigint(255) NULL DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `level` int(11) NOT NULL,
  `request` int(11) NULL DEFAULT NULL,
  `resolve` int(11) NULL DEFAULT NULL,
  `org` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `banner` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `cidr` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `first_time` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `task_id` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 221 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

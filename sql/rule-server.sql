/*
 Navicat Premium Data Transfer

 Source Server         : 腾讯云Rule-Server-MySql
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 1.12.242.39:13306
 Source Schema         : rule-server

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 24/02/2022 19:04:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for aggregator
-- ----------------------------
DROP TABLE IF EXISTS `aggregator`;
CREATE TABLE `aggregator`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '聚合器名称',
  `metric` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '聚合指标',
  `function_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '关联的聚合函数id',
  `rule_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '关联的告警规则id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '聚合器表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of aggregator
-- ----------------------------
INSERT INTO `aggregator` VALUES (10, '内存占用WARNING最大值监测', 'memory_used', 1, 2);
INSERT INTO `aggregator` VALUES (11, '内存占用FATAL最大值监测', 'memory_used', 2, 3);
INSERT INTO `aggregator` VALUES (12, '内存占用平均值监测', 'memory_used', 3, 3);
INSERT INTO `aggregator` VALUES (13, '内存占用求和监测', 'memory_used', 5, 3);
INSERT INTO `aggregator` VALUES (20, 'CPU占用WARNING最大值监测', 'cpu_rate', 1, 2);
INSERT INTO `aggregator` VALUES (21, 'CPU占用FATAL最大值监测', 'cpu_rate', 2, 3);
INSERT INTO `aggregator` VALUES (22, 'CPU占用平均值监测', 'cpu_rate', 3, 3);
INSERT INTO `aggregator` VALUES (23, 'CPU占用求和监测', 'cpu_rate', 5, 3);
INSERT INTO `aggregator` VALUES (30, '磁盘占用WARNING最大值监测', 'disk_used', 1, 2);
INSERT INTO `aggregator` VALUES (31, '磁盘占用FATAL最大值监测', 'disk_used', 2, 3);
INSERT INTO `aggregator` VALUES (32, '磁盘占用平均值监测', 'disk_used', 3, 3);
INSERT INTO `aggregator` VALUES (33, '磁盘占用求和监测', 'disk_used', 5, 3);
INSERT INTO `aggregator` VALUES (41, 'CPU与内存', 'cpu_mem', 6, 3);

-- ----------------------------
-- Table structure for email
-- ----------------------------
DROP TABLE IF EXISTS `email`;
CREATE TABLE `email`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `address_unique`(`address`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户邮箱' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of email
-- ----------------------------
INSERT INTO `email` VALUES (1, 'athena_client@163.com');

-- ----------------------------
-- Table structure for function
-- ----------------------------
DROP TABLE IF EXISTS `function`;
CREATE TABLE `function`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类型',
  `threshold` double NOT NULL COMMENT '阈值',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '聚合函数表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of function
-- ----------------------------
INSERT INTO `function` VALUES (1, 'MAX', 60, 'WARNING型最大值类型函数，聚合数据的最大值不可大于阈值');
INSERT INTO `function` VALUES (2, 'MAX', 90, 'FATAL型最大值函数，聚合数据的最大值不可大于阈值');
INSERT INTO `function` VALUES (3, 'AVG', 90, '平均值类型函数，聚合数据的平均值值不可大于阈值');
INSERT INTO `function` VALUES (4, 'MIN', 10, '最小值类型函数，聚合数据的最小值不可小于阈值');
INSERT INTO `function` VALUES (5, 'SUM', 450, '求和类型函数，聚合结果值不可大于阈值');
INSERT INTO `function` VALUES (6, 'LOGIC', 1, '用于复杂的聚合函数');

-- ----------------------------
-- Table structure for rule
-- ----------------------------
DROP TABLE IF EXISTS `rule`;
CREATE TABLE `rule`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `level` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '告警等级',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '告警行为',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '告警规则表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rule
-- ----------------------------
INSERT INTO `rule` VALUES (1, 'INFO', 'MESSAGE', '特定情况发送短信通知');
INSERT INTO `rule` VALUES (2, 'WARNING', 'EMAIL', '特定情况发送邮件到指定邮箱');
INSERT INTO `rule` VALUES (3, 'FATAL', 'PHONE', '特定情况打电话');

-- ----------------------------
-- Table structure for smtp
-- ----------------------------
DROP TABLE IF EXISTS `smtp`;
CREATE TABLE `smtp`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SMTP服务器',
  `port` int(0) NULL DEFAULT 465 COMMENT '端口',
  `isSSL` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '是否使用SSL加密 0 为禁用、1 为启用',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SMTP登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SMTP密码',
  `from` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '发送邮箱地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'SMTP服务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of smtp
-- ----------------------------
INSERT INTO `smtp` VALUES (1, 'smtp.126.com', 465, 1, 'athena_server@126.com', 'RJPNHMTNLHTKDRGH', 'athena_server@126.com');
INSERT INTO `smtp` VALUES (2, 'smtp.126.com', 465, 1, 'athena_server1@126.com', 'HYUSOPXJJJIEJRRW', 'athena_server1@126.com');
INSERT INTO `smtp` VALUES (3, 'smtp.126.com', 465, 1, 'athena_server2@126.com', 'BTFJVDJKSTTUVUFU', 'athena_server2@126.com');

SET FOREIGN_KEY_CHECKS = 1;

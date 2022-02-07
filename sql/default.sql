/*
 Navicat Premium Data Transfer

 Source Server         : rule-server
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : rule-server

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 03/02/2022 16:22:31
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
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '聚合器表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of aggregator
-- ----------------------------
INSERT INTO `aggregator` VALUES (11, '内存占用最大值监测', 'memory_used', 1, 2);
INSERT INTO `aggregator` VALUES (12, '内存占用平均值监测', 'memory_used', 2, 2);
INSERT INTO `aggregator` VALUES (13, '内存占用最小值监测', 'memory_used', 3, 1);
INSERT INTO `aggregator` VALUES (14, '内存占用求和监测', 'memory_used', 4, 3);
INSERT INTO `aggregator` VALUES (21, 'CPU占用最大值监测', 'cpu_rate', 1, 2);
INSERT INTO `aggregator` VALUES (22, 'CPU占用平均值监测', 'cpu_rate', 2, 2);
INSERT INTO `aggregator` VALUES (23, 'CPU占用最小值监测', 'cpu_rate', 3, 1);
INSERT INTO `aggregator` VALUES (24, 'CPU占用求和监测', 'cpu_rate', 4, 3);
INSERT INTO `aggregator` VALUES (31, '磁盘占用最大值监测', 'disk_used', 1, 2);
INSERT INTO `aggregator` VALUES (32, '磁盘占用平均值监测', 'disk_used', 2, 2);
INSERT INTO `aggregator` VALUES (33, '磁盘占用最小值监测', 'disk_used', 3, 1);
INSERT INTO `aggregator` VALUES (34, '磁盘占用求和监测', 'disk_used', 4, 3);
INSERT INTO `aggregator` VALUES (41, 'CPU与内存', 'cpu_mem', 5, 2);

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
INSERT INTO `function` VALUES (1, 'MAX', 90.0, '最大值类型函数，聚合数据的最大值不可大于阈值');
INSERT INTO `function` VALUES (2, 'AVG', 90.0, '平均值类型函数，聚合数据的平均值值不可大于阈值');
INSERT INTO `function` VALUES (3, 'MIN', 10.0, '最小值类型函数，聚合数据的最小值不可小于阈值');
INSERT INTO `function` VALUES (4, 'SUM', 450.0, '求和类型函数，聚合结果值不可大于阈值');
INSERT INTO `function` VALUES (5, 'LOGIC', 1.0, '用于复杂的聚合函数');

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

SET FOREIGN_KEY_CHECKS = 1;

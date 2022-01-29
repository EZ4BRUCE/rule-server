
CREATE DATABASE IF NOT EXISTS `rule-server` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

USE `rule-server`; 

CREATE TABLE IF NOT EXISTS `function` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) NOT NULL COMMENT '类型',
  `threshold` double NOT NULL COMMENT '阈值',
  `description` varchar(255) COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='聚合函数表';

CREATE TABLE IF NOT EXISTS `rule` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `level` varchar(255) NOT NULL COMMENT '告警等级',
  `action` varchar(255) NOT NULL COMMENT '告警行为',
  `description` varchar(255) DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则表';

-- 不设外键，因为用的时候还是要查
CREATE TABLE IF NOT EXISTS `aggregator` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) UNIQUE COMMENT '聚合器名称',
  `metric` varchar(255) DEFAULT '' COMMENT '聚合指标',
  `function_id` int(10) unsigned COMMENT '关联的聚合函数id',
  `rule_id` int(10) unsigned  COMMENT '关联的告警规则id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='聚合器表';

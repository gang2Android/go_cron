/*
Navicat MySQL Data Transfer

Source Server         : cron_task
Source Server Version : 50644
Source Host           : 127.0.0.1:3306
Source Database       : cron_task

Target Server Type    : MYSQL
Target Server Version : 50644
File Encoding         : 65001

Date: 2021-06-16 19:26:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for planning_tasks
-- ----------------------------
DROP TABLE IF EXISTS `planning_tasks`;
CREATE TABLE `planning_tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL COMMENT '任务名称',
  `type` int(1) DEFAULT '1' COMMENT '任务类型:1url,2数据库',
  `spec` varchar(255) DEFAULT NULL COMMENT '执行周期',
  `enable` int(1) DEFAULT '1' COMMENT '是否启用1是2否',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `content` longtext NOT NULL COMMENT '内容',
  `task_no` varchar(255) NOT NULL DEFAULT '' COMMENT '任务id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='计划任务';

-- ----------------------------
-- Records of planning_tasks
-- ----------------------------
INSERT INTO `planning_tasks` VALUES ('1', '百度', '1', '0 0/1 * * * ?', '1', '2021-06-16 18:46:55', 'https://www.baidu.com/', '1');
INSERT INTO `planning_tasks` VALUES ('2', '商城数据库备份', '2', '0 0 0/1 * * ?', '1', '2021-06-16 18:41:31', '{\"name\":\"shop\",\"host\":\"127.0.0.1\",\"db\":\"db\",\"user\":\"user\",\"pwd\":\"pwd\",\"backPath\":\"/www/wwwroot/a.xxx.com/db/\",\"mysqlPath\":\"/www/server/mysql/bin/\",\"retain\":\"5\"}', '11');

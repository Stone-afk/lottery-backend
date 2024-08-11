/*
 Navicat Premium Data Transfer

 Source Server         : gozero-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : lottery

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 17/01/2024 23:07:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lottery
-- ----------------------------
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery`  (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL DEFAULT 0 COMMENT '发起抽奖用户ID',
                            `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等奖名称',
                            `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等经配图',
                            `publish_time` datetime NULL DEFAULT NULL COMMENT '发布抽奖时间',
                            `join_number` int NOT NULL DEFAULT 0 COMMENT '自动开奖人数',
                            `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '抽奖说明',
                            `award_deadline` datetime NOT NULL COMMENT '领奖截止时间',
                            `is_selected` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否精选 1是 0否',
                            `announce_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '开奖设置：1按时间开奖 2按人数开奖 3即抽即中',
                            `announce_time` datetime NOT NULL DEFAULT NULL COMMENT '开奖时间',
                            `del_state` tinyint NOT NULL DEFAULT '0',
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                            `is_announced` tinyint(1) NULL DEFAULT 0 COMMENT '是否开奖：0未开奖；1已经开奖',
                            `sponsor_id` int NOT NULL DEFAULT 0 COMMENT '发起抽奖赞助商ID',
                            `is_clocked` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启打卡任务：0未开启；1已开启',
                            `clock_task_id` int NOT NULL DEFAULT 0 COMMENT '打卡任务任务ID',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 111 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '抽奖表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;


-- ----------------------------
-- Table structure for prize
-- ----------------------------
DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize`  (
                          `id` int(0) NOT NULL AUTO_INCREMENT,
                          `lottery_id` int(0) NOT NULL DEFAULT 0 COMMENT '抽奖ID',
                          `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包',
                          `name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
                          `level` int(0) NOT NULL DEFAULT 1 COMMENT '几等奖 默认1',
                          `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品图',
                          `count` int(0) NOT NULL DEFAULT 0 COMMENT '奖品份数',
                          `grant_type` tinyint(1) NOT NULL COMMENT '奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序',
                          `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '奖品表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for lottery_participation
-- ----------------------------
CREATE TABLE lottery_participation
(
    id         BIGINT AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    lottery_id INT     NOT NULL COMMENT '参与的抽奖的id',
    user_id    INT     NOT NULL COMMENT '用户id',
    is_won     TINYINT NOT NULL COMMENT '中奖了吗？',
    prize_id   BIGINT  NOT NULL COMMENT '中奖id',
    del_state tinyint NOT NULL DEFAULT '0',
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT index_lottery_user UNIQUE (lottery_id, user_id)
)
    COMMENT '参与抽奖' COLLATE = utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for clock_task
-- ----------------------------
DROP TABLE IF EXISTS `clock_task`;
CREATE TABLE `clock_task` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `lottery_id` int NOT NULL COMMENT '抽奖ID',
                              `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '任务类型 1: 体验小程序 2： 浏览指定公众号文章 3: 浏览图片（微信图片二维码等） 4： 浏览视频号视频',
                              `seconds` int NOT NULL DEFAULT '0' COMMENT '任务秒数',
                              `applet_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'type=1时该字段才有意义，1小程序链接 2小程序路径',
                              `page_link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=1时 该字段才有意义，配置要跳转小程序的页面链接 （如 #小程序://抽奖/oM....）',
                              `app_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序AppID',
                              `page_path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序路径（如：/pages/index）',
                              `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=3时 该字段才有意义，添加要查看的图片',
                              `video_account_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=4时 该字段才有意义，视频号ID',
                              `video_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=4时 该字段才有意义，视频ID',
                              `article_link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=2时 该字段才有意义，公众号文章链接',
                              `copywriting` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '引导参与者完成打卡任务的文案',
                              `chance_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '概率类型 1: 随机 2: 指定 ',
                              `increase_multiple` int NOT NULL DEFAULT '1' COMMENT 'chance_type=2时 该字段才有意义，概率增加倍数',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='打卡任务表';

SET FOREIGN_KEY_CHECKS = 1;


-- ----------------------------
-- Table structure for clock_task_record
-- ----------------------------
DROP TABLE IF EXISTS `clock_task_record`;
CREATE TABLE `clock_task_record` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `lottery_id` int NOT NULL COMMENT '抽奖ID',
                                     `clock_task_id` int NOT NULL COMMENT '打卡任务ID',
                                     `user_id` int NOT NULL COMMENT '用户id',
                                     `increase_multiple` int NOT NULL DEFAULT '1' COMMENT '概率增加倍数',
                                     `del_state` tinyint(1) NOT NULL DEFAULT '0',
                                     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='打卡任务记录表';

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `award_record`;
CREATE TABLE `award_record` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `lottery_id` int NOT NULL COMMENT '抽奖ID',
                                     `user_id` int NOT NULL COMMENT '用户id',
                                     `prize_id`   BIGINT  NOT NULL COMMENT '奖品id',
                                     `del_state` tinyint NOT NULL DEFAULT '0',
                                     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='领取奖品记录表';
SET FOREIGN_KEY_CHECKS = 1;
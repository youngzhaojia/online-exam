-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `userId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'userId',
  `mobile` varchar(64) DEFAULT '' COMMENT '手机号',
  `name` varchar(64) DEFAULT '' COMMENT '姓名',
  `password` varchar(64) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='用户表';

INSERT INTO `online_exam`.`t_user` (`mobile`, `password`) VALUES ('13100000000', '123456');

-- ----------------------------
-- Table structure for t_question
-- ----------------------------
DROP TABLE IF EXISTS `t_question`;
CREATE TABLE `t_question` (
  `questionId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'questionId',
  `userId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类: 1选择题2问答题',
  `title` varchar(255) unsigned NOT NULL DEFAULT '' COMMENT '题目',
  `remark` int(255) unsigned NOT NULL DEFAULT '' COMMENT '备注',
  `option` varchar(255) unsigned NOT NULL DEFAULT '' COMMENT '选择题选项',
  `createTime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`questionId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='试题表';

-- ----------------------------
-- Table structure for t_exam
-- ----------------------------
DROP TABLE IF EXISTS `t_exam`;
CREATE TABLE `t_exam` (
  `examId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'examId',
  `userId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '试卷名称',
  `time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '考试时长(秒)',
  `questionList` varchar(255) unsigned NOT NULL DEFAULT '' COMMENT '试题列表,逗号隔开',
  `createTime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`examId`),
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='试卷表';

-- ----------------------------
-- Table structure for t_answer
-- ----------------------------
DROP TABLE IF EXISTS `t_answer`;
CREATE TABLE `t_answer` (
  `answerId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'answerId',
  `userId` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `answerList` text unsigned NOT NULL DEFAULT '' COMMENT '答案列表',
  `createTime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`answerId`),
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='答题表';

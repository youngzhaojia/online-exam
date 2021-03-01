
use online_exam;
-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'user_id',
  `mobile` varchar(64) DEFAULT '' COMMENT '手机号',
  `name` varchar(64) DEFAULT '' COMMENT '姓名',
  `password` varchar(64) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Table structure for t_question
-- ----------------------------
DROP TABLE IF EXISTS `t_question`;
CREATE TABLE `t_question` (
  `question_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'question_id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '老师用户id',
  `type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类: 1选择题2问答题',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '题目',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `option` varchar(255) NOT NULL DEFAULT '' COMMENT '选择题选项',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`question_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='试题表';

-- ----------------------------
-- Table structure for t_exam
-- ----------------------------
DROP TABLE IF EXISTS `t_exam`;
CREATE TABLE `t_exam` (
  `exam_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'exam_id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '老师用户id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '试卷名称',
  `time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '考试时长(秒)',
  `question_list` varchar(255) NOT NULL DEFAULT '' COMMENT '试题列表,逗号隔开',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`exam_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='试卷表';

-- ----------------------------
-- Table structure for t_answer
-- ----------------------------
DROP TABLE IF EXISTS `t_answer`;
CREATE TABLE `t_answer` (
  `answer_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'answer_id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '学生id',
  `exam_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '试卷id',
  `teacher_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '老师用户id',
  `answer_list` text NOT NULL COMMENT '答案列表',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`answer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='答题表';

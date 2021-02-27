SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`         bigint UNSIGNED                                           NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)                                               NULL DEFAULT NULL,
    `updated_at` datetime(3)                                               NULL DEFAULT NULL,
    `deleted_at` datetime(3)                                               NULL DEFAULT NULL,
    `name`       longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_role_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `user_id` bigint UNSIGNED NOT NULL,
    `role_id` bigint UNSIGNED NOT NULL,
    PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
    INDEX `fk_user_role_role` (`role_id`) USING BTREE,
    CONSTRAINT `fk_user_role_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_user_role_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint UNSIGNED                                                NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3)                                                    NULL DEFAULT NULL,
    `updated_at` datetime(3)                                                    NULL DEFAULT NULL,
    `deleted_at` datetime(3)                                                    NULL DEFAULT NULL,
    `username`   longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci      NULL,
    `password`   longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci      NULL,
    `avatar`     varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_user_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stories
-- ----------------------------
DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories`
(
    `id`             int                                                           NOT NULL AUTO_INCREMENT,
    `time_stamp`     bigint                                                        NULL DEFAULT NULL,
    `name`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `story_describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`
(
    `id`          int                                                           NOT NULL AUTO_INCREMENT,
    `title`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `article_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for images
-- ----------------------------
DROP TABLE IF EXISTS `images`;
CREATE TABLE `images`
(
    `id`         int                                                           NOT NULL AUTO_INCREMENT,
    `image_type` int                                                           NULL DEFAULT NULL,
    `image_url`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for members
-- ----------------------------
DROP TABLE IF EXISTS `members`;
CREATE TABLE `members`
(
    `id`           int                                                           NOT NULL AUTO_INCREMENT,
    `name`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `url_avatar`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `school_year`  int                                                           NULL DEFAULT NULL,
    `describe`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `member_type`  int                                                           NULL DEFAULT NULL,
    `teacher_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for passages
-- ----------------------------
DROP TABLE IF EXISTS `passages`;
CREATE TABLE `passages`
(
    `id`      int                                                           NOT NULL AUTO_INCREMENT,
    `passage` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`
(
    `id`               int                                                           NOT NULL AUTO_INCREMENT,
    `name`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `little_describe`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `describe`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `partner`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `project_type`     int                                                           NULL DEFAULT NULL,
    `url_avatar`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `url_partner_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `url_background`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `url_screenshot`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `url_pro_code`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for interviews
-- ----------------------------
DROP TABLE IF EXISTS `interviews`;
CREATE TABLE `interviews`
(
    `id`          int                                                       NOT NULL AUTO_INCREMENT,
    `name`        longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
    `create_time` bigint                                                    NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for interviewers
-- ----------------------------
DROP TABLE IF EXISTS `interviewers`;
CREATE TABLE `interviewers`
(
    `id`              int                                                           NOT NULL AUTO_INCREMENT,
    `name`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `interview_id`    int                                                           NULL DEFAULT NULL,
    `serverchan_id`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `webhook_url`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `department_type` int                                                           NULL DEFAULT NULL,
    `status`          int                                                           NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for resumes
-- ----------------------------
DROP TABLE IF EXISTS `resumes`;
CREATE TABLE `resumes`
(
    `id`                       int                                                           NOT NULL AUTO_INCREMENT,
    `describe`                 longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NULL,              -- 简历内容文字
    `file_url`                 varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 附件url
    `college_major`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试者院系专业
    `name`                     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试者姓名
    `gender`                   boolean                                                       NULL default null, -- 面试者性别
    `grade`                    int                                                           NULL DEFAULT NULL, -- 面试者年级
    `department_type`          int                                                           NULL DEFAULT NULL, -- 部门类型
    `experience`               boolean                                                       NULL default null, -- 面试者有无项目经历
    `telephone_number`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试者电话号码
    `qq_number`                varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试者QQ号
    `wechat_number`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试者微信联系方式
    `interview_id`             int                                                           NULL DEFAULT NULL, -- 面试项目id
    `interviewer_id`           int                                                           NULL DEFAULT NULL, -- 面试官id
    `send_time`                bigint                                                        NULL DEFAULT NULL, -- 上传简历时间
    `initial_screening_result` int                                                           NULL DEFAULT NULL, -- 初筛结果
    `initial_screening_time`   bigint                                                        NULL DEFAULT NULL,-- 初筛结果提交时间
    `interview_result`         int                                                           NULL DEFAULT NULL, -- 面试结果
    `interview_evaluation`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL, -- 面试评价
    `interview_time`           bigint                                                        NULL DEFAULT NULL, -- 面试时间
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`
(
    `id`          int                                                       NOT NULL AUTO_INCREMENT,
    `content`     longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
    `read_status` int                                                       NULL DEFAULT NULL,
    `resume_id`   int                                                       NULL DEFAULT NULL,
    `create_time` bigint                                                    NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `logs`;
CREATE TABLE `logs`
(
    `id`             int                                                           NOT NULL AUTO_INCREMENT,
    `time_stamp`     bigint                                                        NULL,
    `request_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `username`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `request_url`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `response_code`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `model_name`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
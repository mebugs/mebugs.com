/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 06/06/2025 09:15:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
                            `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                            `account` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登陆账号',
                            `encryption` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码密文',
                            `salt_key` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码盐值（AES加密KEY）',
                            `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
                            `dept_id` bigint(0) NOT NULL COMMENT '部门ID',
                            `is_leader` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门职位，枚举：0_部门员工 1_部门领导',
                            `permission_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局',
                            `pwd_exp_time` datetime(0) NULL DEFAULT NULL COMMENT '密码过期时间，创建即过期',
                            `last_login_time` datetime(0) NULL DEFAULT NULL COMMENT '最后登陆时间，为空表示创建',
                            `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识，枚举：0_可变更 1_禁止变更',
                            `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                            `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                            `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE INDEX `account_uni`(`account`) USING BTREE COMMENT '登陆账号唯一',
                            INDEX `dept_query`(`dept_id`) USING BTREE COMMENT '部门查询'
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登陆账号' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES (2, 'admin', 's/ZFL210W3mQt+qGEy2jFA==', 'gWJoO3l0K0Zrnl42', '管理员', 1, '1', '3', '2025-06-06 09:14:38', NULL, '0', '0', '2024-07-08 08:56:33', '2025-06-06 09:14:38');

-- ----------------------------
-- Table structure for account_role
-- ----------------------------
DROP TABLE IF EXISTS `account_role`;
CREATE TABLE `account_role`  (
                                 `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                                 `account_id` bigint(0) NOT NULL COMMENT '账号ID',
                                 `role_id` bigint(0) NOT NULL COMMENT '角色ID',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE INDEX `account_role_uni`(`account_id`, `role_id`) USING BTREE COMMENT '账号角色唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '账号角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account_role
-- ----------------------------
INSERT INTO `account_role` VALUES (10, 2, 5);

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`  (
                           `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                           `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
                           `tag` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签（仅展示）',
                           `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '类型：0Banner 1Page',
                           `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类地址',
                           `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                           `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                           `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                           `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                           `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           UNIQUE INDEX `title_uni`(`title`) USING BTREE COMMENT '名称唯一',
                           UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一',
                           INDEX `update_time`(`update_at`) USING BTREE COMMENT '分页索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Banner' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of banner
-- ----------------------------

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
                             `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                             `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
                             `url` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类地址',
                             `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                             `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                             `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                             `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                             `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `title_uni`(`title`) USING BTREE COMMENT '名称唯一',
                             UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分组分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept`  (
                         `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                         `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
                         `pid` bigint(0) NOT NULL DEFAULT 0 COMMENT '父级部门ID，租户创建时默认创建根部门，父级ID=0',
                         `sort` int(0) NOT NULL DEFAULT 0 COMMENT '同级部门排序',
                         `permission_type` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局 4_指定部门 5_指定人',
                         `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识，枚举：0_可变更 1_禁止变更',
                         `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                         `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                         `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `dept_uni`(`name`) USING BTREE COMMENT '部门名唯一',
                         INDEX `dept_pid_query`(`pid`) USING BTREE COMMENT '部门树构建'
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集团部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dept
-- ----------------------------
INSERT INTO `dept` VALUES (1, '站线集团', 0, 0, '0', '0', '0', NULL, NULL);

-- ----------------------------
-- Table structure for dict
-- ----------------------------
DROP TABLE IF EXISTS `dict`;
CREATE TABLE `dict`  (
                         `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                         `group_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典分组KEY',
                         `label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段名称',
                         `label_en` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段名称（英文）',
                         `choose` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '是否可被选择 0可选择 1不可选择',
                         `val` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典值（字符型）',
                         `pid` bigint(0) NULL DEFAULT NULL COMMENT '父级字典ID 默认 0（根数据）',
                         `sort` smallint(0) NULL DEFAULT NULL COMMENT '字典排序',
                         `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典描述',
                         `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
                         `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                         `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                         `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `dict_uni`(`group_key`, `val`) USING BTREE COMMENT '同一分组值唯一',
                         INDEX `dict_pgs`(`group_key`, `pid`, `sort`) USING BTREE COMMENT '父级ID，分组过滤，排序'
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict
-- ----------------------------
INSERT INTO `dict` VALUES (1, 'accountPermission', '继承部门', 'Same department', '0', '0', 1, 0, '账号权限：继承部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (2, 'accountPermission', '本部门', 'Main department', '0', '1', 1, 0, '账号权限：本部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (3, 'accountPermission', '本人', 'Personal', '0', '2', 1, 0, '账号权限：本人', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (4, 'accountPermission', '全部', 'All', '0', '3', 1, 0, '账号权限：全部', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (5, 'accountStatus', '正常', 'Normal', '0', '0', 1, 0, '账号状态：正常', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (6, 'accountStatus', '锁定', 'Locked', '0', '1', 1, 0, '账号状态：锁定', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (7, 'deptAccount', '部门员工', 'Department staff', '0', '0', 1, 0, '部门成员：部门员工', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (8, 'deptAccount', '部门领导', 'Department heads', '0', '1', 1, 0, '部门成员：部门领导', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (9, 'deptPermission', '本部门及子部门', 'Main and sub departments', '0', '0', 1, 0, '部门权限：本部门与子部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (10, 'deptPermission', '本部门', 'Main department', '0', '1', 1, 0, '部门权限：本部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (11, 'deptPermission', '本人', 'Personal', '0', '2', 1, 0, '部门权限：本人', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (12, 'deptPermission', '全部', 'All', '0', '3', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (13, 'deptPermission', '指定部门', 'Choose department', '1', '4', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (14, 'deptPermission', '指定人', 'Choose person', '1', '5', 1, 0, '字典分组：部门权限类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (15, 'deptToType', '部门合入', 'Fit into', '0', '0', 1, 0, '合并模式：部门合入', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (16, 'deptToType', '部门并入', 'Merge into', '0', '1', 1, 0, '合并模式：部门并入', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (17, 'dictGroup', '账号数据权限', 'Account data permissions', '0', 'accountPermission', 1, 0, '字典分组：账号数据权限', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (18, 'dictGroup', '账号状态', 'Account status', '0', 'accountStatus', 1, 0, '字典分组：账号状态', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (19, 'dictGroup', '部门账号类型', 'Department account type', '0', 'deptAccount', 1, 0, '字典分组：部门账号类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (20, 'dictGroup', '部门数据权限', 'deptPermission', '0', 'deptPermission', 1, 0, '字典分组：部门数据权限', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (21, 'dictGroup', '部门合并类型', 'Department merger type', '0', 'deptToType', 1, 0, '字典分组：部门合并类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (22, 'dictGroup', '字典分组', 'Dict Group', '0', 'dictGroup', 1, 0, '字典分组：字典分组', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (23, 'dictGroup', '启用状态', 'Open Status', '0', 'openStatus', 1, 2, '字典分组：启用状态', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (24, 'dictGroup', '权限等级', 'Permission Level', '0', 'permissionLevel', 1, 3, '字典分组：权限等级', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (25, 'dictGroup', '响应类型', 'Response Type', '0', 'responseType', 1, 4, '字典分组：响应类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (26, 'dictGroup', '路由类型', 'Router Type', '0', 'routerType', 1, 5, '字典分组：路由类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (27, 'dictGroup', '业务编码', 'Service code', '0', 'serviceCode', 1, 1, '字典分组：业务编码', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (28, 'dictGroup', '时间周期', 'Time unit', '0', 'timeUnit', 1, 6, '字典分组：时间周期', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (29, 'openStatus', '启用', 'Enable', '0', '0', 1, 0, '启用状态：启用', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (30, 'openStatus', '禁用', 'Disable', '0', '1', 1, 0, '启用状态：禁用', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (31, 'permissionLevel', '系统', 'System', '1', '0', 1, 0, '权限等级：系统', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (32, 'permissionLevel', '模块', 'Model', '1', '1', 1, 0, '权限等级：模块', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (33, 'permissionLevel', '页面', 'Page', '1', '2', 1, 0, '权限等级：页面', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (34, 'permissionLevel', '按钮', 'Button', '1', '3', 1, 0, '权限等级：按钮', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (35, 'responseType', '异常', 'Error', '0', 'E', 1, 3, '响应类型：异常', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (36, 'responseType', '失败', 'Fail', '0', 'F', 1, 2, '响应类型：失败', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (37, 'responseType', '成功', 'Success', '0', 'S', 1, 0, '响应类型：成功', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (38, 'routerType', '鉴权路由', 'Authentication Router', '0', '0', 1, 0, '路由类型：需要登陆授权', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (39, 'routerType', '白名单路由', 'Whitelist Router', '0', '1', 1, 0, '路由类型：白名单', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (40, 'serviceCode', '框架', 'Frame', '0', '0', 1, 0, '模块：基础框架', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (41, 'serviceCode', '字典', 'Dict', '0', '1', 1, 7, '模块：字典', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (42, 'serviceCode', '响应', 'Response', '0', '2', 1, 6, '模块：响应配置', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (43, 'serviceCode', '路由', 'Router', '0', '3', 1, 5, '模块：接口路由', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (44, 'serviceCode', '权限', 'Permissions', '0', '4', 1, 4, '模块：权限集', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (45, 'serviceCode', '账号', 'Account', '0', '5', 1, 3, '模块：账号', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (46, 'serviceCode', '角色', 'Role', '0', '6', 1, 2, '模块：角色', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (47, 'serviceCode', '部门', 'Dept', '0', '7', 1, 1, '模块：部门', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (48, 'serviceCode', '配置', 'Config', '0', '8', 1, 8, '模块：配置', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (49, 'timeUnit', '默认(毫秒)', 'Default (milliseconds)', '0', '0', 1, 0, '时间单元：默认(毫秒)', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (50, 'timeUnit', '秒', 'Second', '0', '1', 1, 0, '时间单元：秒', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (51, 'timeUnit', '分', 'Point', '0', '2', 1, 0, '时间单元：分', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (52, 'timeUnit', '时', 'Hour', '0', '3', 1, 0, '时间单元：时', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (53, 'timeUnit', '日', 'Day', '0', '4', 1, 0, '时间单元：日', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (54, 'timeUnit', '月', 'Month', '0', '5', 1, 0, '时间单元：月', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (55, 'timeUnit', '年', 'Year', '0', '6', 1, 0, '时间单元：年', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (56, 'dictGroup', '资源类型', '资源类型', '0', 'sourceType', 1, 0, '资源类型', '0', '0', '2024-11-28 18:39:36', '2024-11-28 18:39:36');
INSERT INTO `dict` VALUES (57, 'sourceType', '文章主图', '文章主图', '0', '0', 1, 0, '资源图片类型', '0', '0', '2024-11-28 18:43:47', '2024-11-28 18:43:47');
INSERT INTO `dict` VALUES (58, 'sourceType', '方形图标', '方形图标', '0', '1', 1, 0, '图标', '0', '0', '2024-11-28 18:44:20', '2024-11-28 18:44:20');
INSERT INTO `dict` VALUES (59, 'sourceType', '单张资源', '单张资源', '0', '2', 1, 0, '正文', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (60, 'dictGroup', '文章状态', 'Post status', '0', 'postStatus', 1, 7, '字典分组：文章状态', '0', '0', NULL, NULL);
INSERT INTO `dict` VALUES (62, 'postStatus', '草稿', '草稿', '0', '0', 1, 0, '文章状态：0草稿', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (63, 'postStatus', '发布', '发布', '0', '1', 1, 1, '文章状态：1发布', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (64, 'postStatus', '封存', '封存', '0', '2', 1, 2, '文章状态：2发布', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (65, 'bannerType', '画廊', '画廊', '0', '0', 1, 0, 'Banner状态：0画廊', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (66, 'bannerType', '页面', '页面', '0', '1', 1, 1, 'Banner状态：1页面', '0', '0', '2024-11-28 18:44:42', '2024-11-28 18:44:42');
INSERT INTO `dict` VALUES (67, 'dictGroup', '链接类型', 'Banner Type', '0', 'bannerType', 1, 6, '字典分组：链接类型', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (68, 'serviceCode', '分类', 'Category', '0', '9', 1, 9, '模块：分类', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (69, 'serviceCode', '资源', 'Source', '0', '10', 1, 10, '模块：资源', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (70, 'serviceCode', '标签', 'Tag', '0', '11', 1, 11, '模块：标签', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (71, 'serviceCode', '专题', 'Topic', '0', '12', 1, 12, '模块：专题', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (72, 'serviceCode', '文章', 'Post', '0', '13', 1, 13, '模块：文章', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (73, 'serviceCode', '链接', 'Banner', '0', '14', 1, 14, '模块：链接', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (74, 'serviceCode', '评论', 'PostComment', '0', '15', 1, 15, '模块：评论', '1', '0', NULL, NULL);
INSERT INTO `dict` VALUES (75, 'serviceCode', '友链', 'Links', '0', '16', 1, 16, '模块：友链', '1', '0', NULL, NULL);

-- ----------------------------
-- Table structure for links
-- ----------------------------
DROP TABLE IF EXISTS `links`;
CREATE TABLE `links`  (
                          `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                          `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
                          `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '友链地址',
                          `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                          `summary` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                          `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                          `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                          `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '友链\r\n' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of links
-- ----------------------------

-- ----------------------------
-- Table structure for login_record
-- ----------------------------
DROP TABLE IF EXISTS `login_record`;
CREATE TABLE `login_record`  (
                                 `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                                 `account_id` bigint(0) NOT NULL COMMENT '账号ID',
                                 `login_type` smallint(0) NULL DEFAULT NULL COMMENT '登陆类型 1平台账号登录',
                                 `login_time` datetime(0) NULL DEFAULT NULL COMMENT '登陆时间',
                                 `token` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆Token',
                                 `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0登陆成功 1主动登出 2被动登出',
                                 `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                                 `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                 `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 INDEX `login_time_query`(`mark`, `account_id`, `login_time`) USING BTREE COMMENT '登陆时间定时任务批量查询'
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '登陆记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_record
-- ----------------------------

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
                               `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                               `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称，界面展示，建议与界面导航一致',
                               `alias` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限别名，英文，规范如下：sys，sysAccount sysAccountAdd',
                               `level` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现',
                               `pid` bigint(0) NOT NULL DEFAULT 0 COMMENT '父级ID，默认为1',
                               `sort` int(0) NOT NULL DEFAULT 0 COMMENT '字典排序',
                               `static` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾',
                               `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
                               `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                               `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                               `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `alias_uni`(`alias`) USING BTREE COMMENT '权限别名唯一',
                               UNIQUE INDEX `name_uni`(`name`) USING BTREE COMMENT '权限名唯一',
                               INDEX `permission_pid_query`(`pid`, `sort`) USING BTREE COMMENT '权限父级ID查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '根权限', 'ROOT', '0', 0, 1, '0', '1', '0', NULL, NULL);
INSERT INTO `permission` VALUES (2, '通用接口', 'Open', '1', 1, 0, '0', '1', '0', NULL, '2025-05-28 11:05:51');
INSERT INTO `permission` VALUES (3, '主页', 'Center', '2', 1, 1, '0', '1', '0', NULL, '2024-07-31 08:39:09');
INSERT INTO `permission` VALUES (4, '平台管理', 'Plat', '1', 1, 2, '1', '1', '0', NULL, NULL);
INSERT INTO `permission` VALUES (7, '角色', 'PlatRole', '2', 4, 2, '1', '0', '0', '2024-05-28 08:38:24', '2024-07-31 21:23:30');
INSERT INTO `permission` VALUES (9, '账号', 'PlatAccount', '2', 4, 0, '1', '0', '0', '2024-05-28 08:50:32', '2024-07-31 14:57:57');
INSERT INTO `permission` VALUES (10, '权限', 'PlatPermission', '2', 4, 3, '1', '0', '0', '2024-05-28 08:50:44', '2024-07-31 09:01:57');
INSERT INTO `permission` VALUES (11, '工作中心', 'CenterIndex', '3', 3, 0, '1', '0', '0', '2024-06-13 16:17:38', '2024-07-31 08:37:39');
INSERT INTO `permission` VALUES (12, '部门', 'PlatDept', '2', 4, 1, '1', '0', '0', '2024-06-21 18:24:48', '2024-07-31 08:54:46');
INSERT INTO `permission` VALUES (13, '账号新建', 'PlatAccountAdd', '3', 9, 1, '1', '0', '0', '2024-07-31 08:47:53', '2024-07-31 21:27:54');
INSERT INTO `permission` VALUES (14, '账号编辑', 'PlatAccountEdit', '3', 9, 2, '1', '0', '0', '2024-07-31 08:49:01', '2024-07-31 08:49:01');
INSERT INTO `permission` VALUES (15, '账号删除', 'PlatAccountDel', '3', 9, 3, '1', '0', '0', '2024-07-31 08:49:24', '2024-07-31 08:49:24');
INSERT INTO `permission` VALUES (16, '密码重置', 'PlatAccountReset', '3', 9, 4, '1', '0', '0', '2024-07-31 08:49:45', '2024-07-31 08:49:45');
INSERT INTO `permission` VALUES (17, '部门新建', 'PlatDeptAdd', '3', 12, 1, '1', '0', '0', '2024-07-31 08:55:56', '2024-07-31 08:55:56');
INSERT INTO `permission` VALUES (18, '部门编辑', 'PlatDeptEdit', '3', 12, 2, '1', '0', '0', '2024-07-31 08:56:18', '2024-07-31 08:56:18');
INSERT INTO `permission` VALUES (19, '部门删除', 'PlatDeptDel', '3', 12, 3, '1', '0', '0', '2024-07-31 08:56:43', '2024-07-31 08:56:43');
INSERT INTO `permission` VALUES (20, '部门排序', 'PlatDeptSort', '3', 12, 5, '1', '0', '0', '2024-07-31 08:57:15', '2024-07-31 09:04:10');
INSERT INTO `permission` VALUES (21, '部门迁移', 'PlatDeptMerge', '3', 12, 4, '1', '0', '0', '2024-07-31 08:57:46', '2024-07-31 08:57:46');
INSERT INTO `permission` VALUES (22, '角色新建', 'PlatRoleAdd', '3', 7, 1, '1', '0', '0', '2024-07-31 08:59:57', '2024-07-31 08:59:57');
INSERT INTO `permission` VALUES (23, '角色编辑', 'PlatRoleEdit', '3', 7, 2, '1', '0', '0', '2024-07-31 09:00:14', '2024-07-31 09:00:14');
INSERT INTO `permission` VALUES (24, '角色删除', 'PlatRoleDel', '3', 7, 3, '1', '0', '0', '2024-07-31 09:00:38', '2024-07-31 09:00:38');
INSERT INTO `permission` VALUES (25, '权限新建', 'PlatPermissionAdd', '3', 10, 1, '1', '0', '0', '2024-07-31 09:02:36', '2024-07-31 09:02:46');
INSERT INTO `permission` VALUES (26, '权限编辑', 'PlatPermissionEdit', '3', 10, 2, '1', '0', '0', '2024-07-31 09:03:13', '2024-07-31 09:03:13');
INSERT INTO `permission` VALUES (27, '权限删除', 'PlatPermissionDel', '3', 10, 3, '1', '0', '0', '2024-07-31 09:03:36', '2024-07-31 09:03:36');
INSERT INTO `permission` VALUES (28, '权限排序', 'PlatPermissionSort', '3', 10, 4, '1', '0', '0', '2024-07-31 09:04:02', '2024-07-31 09:04:02');
INSERT INTO `permission` VALUES (29, '路由', 'PlatRouter', '2', 4, 4, '1', '0', '0', '2024-07-31 09:05:03', '2024-07-31 15:03:18');
INSERT INTO `permission` VALUES (30, '路由新建', 'PlatRouterAdd', '3', 29, 1, '1', '0', '0', '2024-07-31 09:05:31', '2024-07-31 09:05:31');
INSERT INTO `permission` VALUES (31, '路由编辑', 'PlatRouterEdit', '3', 29, 2, '1', '0', '0', '2024-07-31 09:05:50', '2024-07-31 09:05:50');
INSERT INTO `permission` VALUES (32, '路由删除', 'PlatRouterDel', '3', 29, 3, '1', '0', '0', '2024-07-31 09:06:12', '2024-07-31 09:06:12');
INSERT INTO `permission` VALUES (33, '响应', 'PlatResponse', '2', 4, 5, '1', '0', '0', '2024-07-31 09:06:57', '2024-07-31 09:06:57');
INSERT INTO `permission` VALUES (34, '字典', 'PlatDict', '2', 4, 6, '1', '0', '0', '2024-07-31 09:07:38', '2024-07-31 09:07:38');
INSERT INTO `permission` VALUES (35, '配置', 'PlatConfig', '2', 4, 7, '1', '0', '0', '2024-07-31 09:08:13', '2024-08-23 16:14:41');
INSERT INTO `permission` VALUES (36, '响应码新建', 'PlatResponseAdd', '3', 33, 1, '1', '0', '0', '2024-07-31 09:09:09', '2024-07-31 09:09:09');
INSERT INTO `permission` VALUES (37, '响应码编辑', 'PlatResponseEdit', '3', 33, 2, '1', '0', '0', '2024-07-31 09:09:31', '2024-07-31 21:18:39');
INSERT INTO `permission` VALUES (38, '响应码删除', 'PlatResponseDel', '3', 33, 3, '1', '0', '0', '2024-07-31 09:09:49', '2024-07-31 09:09:49');
INSERT INTO `permission` VALUES (39, '字典新建', 'PlatDictAdd', '3', 34, 1, '1', '0', '0', '2024-07-31 09:10:37', '2024-07-31 09:10:37');
INSERT INTO `permission` VALUES (40, '字典编辑', 'PlatDictEdit', '3', 34, 2, '1', '0', '0', '2024-07-31 09:10:59', '2024-07-31 09:10:59');
INSERT INTO `permission` VALUES (41, '字典删除', 'PlatDictDel', '3', 34, 3, '1', '0', '0', '2024-07-31 09:11:18', '2024-07-31 09:11:18');
INSERT INTO `permission` VALUES (42, '字典排序', 'PlatDictSort', '3', 34, 4, '1', '0', '0', '2024-07-31 09:11:43', '2024-07-31 09:11:43');
INSERT INTO `permission` VALUES (43, '配置编辑', 'PlatConfigEdit', '3', 35, 1, '1', '0', '0', '2024-07-31 09:12:12', '2024-07-31 09:12:12');
INSERT INTO `permission` VALUES (44, '账号详情', 'PlatAccountView', '3', 9, 0, '1', '0', '0', '2024-07-31 14:55:20', '2024-07-31 14:55:20');
INSERT INTO `permission` VALUES (45, '部门详情', 'PlatDeptView', '3', 12, 0, '1', '0', '0', '2024-07-31 14:58:40', '2024-07-31 14:58:40');
INSERT INTO `permission` VALUES (46, '角色详情', 'PlatRoleView', '3', 7, 0, '1', '0', '0', '2024-07-31 14:59:29', '2024-07-31 14:59:29');
INSERT INTO `permission` VALUES (47, '权限详情', 'PlatPermissionView', '3', 10, 0, '1', '0', '0', '2024-07-31 15:01:36', '2024-07-31 15:01:36');
INSERT INTO `permission` VALUES (48, '路由详情', 'PlatRouterView', '3', 29, 0, '1', '0', '0', '2024-07-31 15:02:12', '2024-07-31 15:02:12');
INSERT INTO `permission` VALUES (49, '响应码详情', 'PlatResponseView', '3', 33, 0, '1', '0', '0', '2024-07-31 15:02:51', '2024-07-31 15:02:51');
INSERT INTO `permission` VALUES (50, '字典详情', 'PlatDictView', '3', 34, 0, '1', '0', '0', '2024-07-31 15:03:55', '2024-07-31 15:03:55');
INSERT INTO `permission` VALUES (52, '配置详情', 'PlatConfigView', '3', 35, 0, '1', '0', '0', '2024-07-31 21:24:18', '2024-07-31 21:24:18');
INSERT INTO `permission` VALUES (53, '博客管理', 'Blog', '1', 1, 3, '1', '0', '0', '2025-05-27 14:39:15', '2025-05-28 11:27:32');
INSERT INTO `permission` VALUES (56, '资源', 'BlogSource', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (57, '链接', 'BlogBanner', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (58, '分类', 'BlogCategory', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (59, '标签', 'BlogTag', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (60, '文章', 'BlogPost', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (61, '评论', 'BlogComment', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');
INSERT INTO `permission` VALUES (62, '友链', 'BlogLinks', '2', 53, 0, '1', '0', '0', '2025-05-28 11:09:53', '2025-05-28 11:09:53');

-- ----------------------------
-- Table structure for permission_router
-- ----------------------------
DROP TABLE IF EXISTS `permission_router`;
CREATE TABLE `permission_router`  (
                                      `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                                      `permission_id` bigint(0) NOT NULL COMMENT '权限ID',
                                      `router_id` bigint(0) NOT NULL COMMENT '路由ID',
                                      PRIMARY KEY (`id`) USING BTREE,
                                      UNIQUE INDEX `permission_router_uni`(`permission_id`, `router_id`) USING BTREE COMMENT '权限和路由组合唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 173 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限路由' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission_router
-- ----------------------------
INSERT INTO `permission_router` VALUES (95, 2, 2);
INSERT INTO `permission_router` VALUES (96, 2, 54);
INSERT INTO `permission_router` VALUES (97, 2, 67);
INSERT INTO `permission_router` VALUES (88, 7, 45);
INSERT INTO `permission_router` VALUES (72, 9, 8);
INSERT INTO `permission_router` VALUES (73, 9, 17);
INSERT INTO `permission_router` VALUES (74, 9, 25);
INSERT INTO `permission_router` VALUES (75, 9, 44);
INSERT INTO `permission_router` VALUES (36, 10, 25);
INSERT INTO `permission_router` VALUES (38, 10, 31);
INSERT INTO `permission_router` VALUES (37, 10, 33);
INSERT INTO `permission_router` VALUES (22, 12, 14);
INSERT INTO `permission_router` VALUES (23, 12, 17);
INSERT INTO `permission_router` VALUES (24, 12, 25);
INSERT INTO `permission_router` VALUES (90, 13, 4);
INSERT INTO `permission_router` VALUES (16, 14, 6);
INSERT INTO `permission_router` VALUES (17, 15, 5);
INSERT INTO `permission_router` VALUES (18, 16, 9);
INSERT INTO `permission_router` VALUES (25, 17, 10);
INSERT INTO `permission_router` VALUES (26, 18, 13);
INSERT INTO `permission_router` VALUES (27, 19, 12);
INSERT INTO `permission_router` VALUES (45, 20, 11);
INSERT INTO `permission_router` VALUES (46, 20, 15);
INSERT INTO `permission_router` VALUES (30, 21, 16);
INSERT INTO `permission_router` VALUES (33, 22, 40);
INSERT INTO `permission_router` VALUES (34, 23, 42);
INSERT INTO `permission_router` VALUES (35, 24, 41);
INSERT INTO `permission_router` VALUES (40, 25, 27);
INSERT INTO `permission_router` VALUES (41, 26, 30);
INSERT INTO `permission_router` VALUES (42, 27, 29);
INSERT INTO `permission_router` VALUES (44, 28, 28);
INSERT INTO `permission_router` VALUES (43, 28, 32);
INSERT INTO `permission_router` VALUES (82, 29, 25);
INSERT INTO `permission_router` VALUES (83, 29, 50);
INSERT INTO `permission_router` VALUES (50, 30, 46);
INSERT INTO `permission_router` VALUES (51, 31, 48);
INSERT INTO `permission_router` VALUES (52, 32, 47);
INSERT INTO `permission_router` VALUES (53, 33, 25);
INSERT INTO `permission_router` VALUES (54, 33, 39);
INSERT INTO `permission_router` VALUES (57, 34, 22);
INSERT INTO `permission_router` VALUES (56, 34, 24);
INSERT INTO `permission_router` VALUES (55, 34, 25);
INSERT INTO `permission_router` VALUES (94, 35, 25);
INSERT INTO `permission_router` VALUES (61, 36, 34);
INSERT INTO `permission_router` VALUES (60, 36, 38);
INSERT INTO `permission_router` VALUES (85, 37, 36);
INSERT INTO `permission_router` VALUES (63, 38, 35);
INSERT INTO `permission_router` VALUES (65, 39, 18);
INSERT INTO `permission_router` VALUES (64, 39, 23);
INSERT INTO `permission_router` VALUES (66, 40, 21);
INSERT INTO `permission_router` VALUES (67, 41, 20);
INSERT INTO `permission_router` VALUES (68, 42, 19);
INSERT INTO `permission_router` VALUES (69, 42, 26);
INSERT INTO `permission_router` VALUES (70, 43, 51);
INSERT INTO `permission_router` VALUES (71, 44, 7);
INSERT INTO `permission_router` VALUES (76, 45, 14);
INSERT INTO `permission_router` VALUES (77, 46, 43);
INSERT INTO `permission_router` VALUES (79, 47, 31);
INSERT INTO `permission_router` VALUES (80, 48, 49);
INSERT INTO `permission_router` VALUES (81, 49, 37);
INSERT INTO `permission_router` VALUES (84, 50, 22);
INSERT INTO `permission_router` VALUES (89, 52, 52);
INSERT INTO `permission_router` VALUES (134, 53, 55);
INSERT INTO `permission_router` VALUES (171, 53, 56);
INSERT INTO `permission_router` VALUES (135, 53, 57);
INSERT INTO `permission_router` VALUES (136, 53, 58);
INSERT INTO `permission_router` VALUES (137, 53, 59);
INSERT INTO `permission_router` VALUES (138, 53, 60);
INSERT INTO `permission_router` VALUES (139, 53, 61);
INSERT INTO `permission_router` VALUES (140, 53, 62);
INSERT INTO `permission_router` VALUES (141, 53, 63);
INSERT INTO `permission_router` VALUES (142, 53, 64);
INSERT INTO `permission_router` VALUES (143, 53, 65);
INSERT INTO `permission_router` VALUES (172, 53, 66);
INSERT INTO `permission_router` VALUES (144, 53, 68);
INSERT INTO `permission_router` VALUES (145, 53, 69);
INSERT INTO `permission_router` VALUES (146, 53, 70);
INSERT INTO `permission_router` VALUES (147, 53, 71);
INSERT INTO `permission_router` VALUES (148, 53, 72);
INSERT INTO `permission_router` VALUES (149, 53, 73);
INSERT INTO `permission_router` VALUES (150, 53, 74);
INSERT INTO `permission_router` VALUES (151, 53, 75);
INSERT INTO `permission_router` VALUES (170, 53, 76);
INSERT INTO `permission_router` VALUES (152, 53, 77);
INSERT INTO `permission_router` VALUES (153, 53, 78);
INSERT INTO `permission_router` VALUES (154, 53, 79);
INSERT INTO `permission_router` VALUES (155, 53, 80);
INSERT INTO `permission_router` VALUES (156, 53, 81);
INSERT INTO `permission_router` VALUES (157, 53, 82);
INSERT INTO `permission_router` VALUES (158, 53, 83);
INSERT INTO `permission_router` VALUES (159, 53, 84);
INSERT INTO `permission_router` VALUES (160, 53, 85);
INSERT INTO `permission_router` VALUES (173, 53, 86);
INSERT INTO `permission_router` VALUES (161, 53, 87);
INSERT INTO `permission_router` VALUES (162, 53, 88);
INSERT INTO `permission_router` VALUES (163, 53, 89);
INSERT INTO `permission_router` VALUES (164, 53, 90);
INSERT INTO `permission_router` VALUES (165, 53, 91);
INSERT INTO `permission_router` VALUES (166, 53, 92);
INSERT INTO `permission_router` VALUES (167, 53, 93);
INSERT INTO `permission_router` VALUES (168, 53, 94);
INSERT INTO `permission_router` VALUES (169, 53, 95);

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
                         `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                         `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
                         `url` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文章地址',
                         `summary` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '摘要',
                         `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                         `post_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文章类型，枚举：0_页面 1_文章 ',
                         `category_id` bigint(0) NULL DEFAULT NULL COMMENT '分组ID，文章可用',
                         `push_at` datetime(0) NULL DEFAULT NULL COMMENT '发布时间，可以自动任务抓取，不配代表不自动发布',
                         `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_草稿 1_发布 2_锁定',
                         `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                         `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                         `views` bigint(0) NULL DEFAULT NULL COMMENT '总浏览量，访问+1',
                         `goods` bigint(0) NULL DEFAULT NULL COMMENT '总支持量，打开页面持续30秒+1，持续120秒+5，持续300秒+10',
                         `hots` bigint(0) NULL DEFAULT NULL COMMENT '近期热度，浏览量+，每周零点/100',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `title_uni`(`title`) USING BTREE COMMENT '名称唯一',
                         UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章或页面' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------

-- ----------------------------
-- Table structure for post_comment
-- ----------------------------
DROP TABLE IF EXISTS `post_comment`;
CREATE TABLE `post_comment`  (
                                 `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                                 `post_id` bigint(0) NULL DEFAULT NULL COMMENT '文章ID',
                                 `level` smallint(0) NULL DEFAULT NULL COMMENT '评论等级，枚举：0_评论 1_回复',
                                 `uid` bigint(0) NULL DEFAULT NULL COMMENT '评论者ID',
                                 `rid` bigint(0) NULL DEFAULT NULL COMMENT '回复的评论者ID（来自被回复的数据）',
                                 `info` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '评论内容（暂不支持HTML）',
                                 `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                                 `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                 `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 INDEX `query`(`post_id`, `create_at`) USING BTREE COMMENT '查询'
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章评论' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_comment
-- ----------------------------

-- ----------------------------
-- Table structure for post_more
-- ----------------------------
DROP TABLE IF EXISTS `post_more`;
CREATE TABLE `post_more`  (
                              `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                              `toc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文章导航',
                              `md` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'MD源文件',
                              `html` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'HTML源文件',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章关联' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_more
-- ----------------------------

-- ----------------------------
-- Table structure for post_source
-- ----------------------------
DROP TABLE IF EXISTS `post_source`;
CREATE TABLE `post_source`  (
                                `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                                `post_id` bigint(0) NULL DEFAULT NULL COMMENT '对象ID',
                                `source_id` bigint(0) NULL DEFAULT NULL COMMENT '资源ID',
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `info_uni`(`post_id`, `source_id`) USING BTREE COMMENT '文章和资源唯一关联'
) ENGINE = InnoDB AUTO_INCREMENT = 104 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章引用资源' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_source
-- ----------------------------

-- ----------------------------
-- Table structure for post_tag
-- ----------------------------
DROP TABLE IF EXISTS `post_tag`;
CREATE TABLE `post_tag`  (
                             `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                             `post_id` bigint(0) NULL DEFAULT NULL COMMENT '文章ID',
                             `tag_id` bigint(0) NULL DEFAULT NULL COMMENT '标签ID',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `info_uni`(`post_id`, `tag_id`) USING BTREE COMMENT '文章和标签唯一关联'
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章引用标签' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_tag
-- ----------------------------

-- ----------------------------
-- Table structure for post_topic
-- ----------------------------
DROP TABLE IF EXISTS `post_topic`;
CREATE TABLE `post_topic`  (
                               `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                               `post_id` bigint(0) NULL DEFAULT NULL COMMENT '文章ID',
                               `topic_id` bigint(0) NULL DEFAULT NULL COMMENT '专题ID',
                               `sort` int(0) NULL DEFAULT NULL COMMENT '排序序号',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `info_uni`(`post_id`, `topic_id`) USING BTREE COMMENT '文章和资源唯一关联'
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章引用主题' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_topic
-- ----------------------------

-- ----------------------------
-- Table structure for response_code
-- ----------------------------
DROP TABLE IF EXISTS `response_code`;
CREATE TABLE `response_code`  (
                                  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                                  `code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应码',
                                  `service_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务ID，来源于字典，指定响应码归属业务',
                                  `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应类型，该字段用于筛选，可配置2和5',
                                  `zh_cn` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '中文响应文言',
                                  `en_us` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '英文响应文言',
                                  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '其他备注信息',
                                  `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
                                  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                                  `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                                  `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `code_uni`(`code`) USING BTREE COMMENT '响应码唯一',
                                  INDEX `query_filter`(`service_code`, `type`) USING BTREE COMMENT '平台筛选查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 134 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '响应码配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of response_code
-- ----------------------------
INSERT INTO `response_code` VALUES (1, 'E000', '0', 'E', '系统异常', 'E', '系统异常（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (2, 'E001', '0', 'E', '参数非法', 'E', '参数非法（默认）（免翻译）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (3, 'E002', '0', 'E', '尚未登陆', 'E', '尚未登陆（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (4, 'E003', '0', 'E', '无权访问', 'E', '无权访问（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (5, 'E004', '0', 'E', '路径不存在', 'E', '路径不存在（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (6, 'S000', '0', 'S', '处理成功', 'E', '处理成功（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (7, 'S001', '0', 'S', '登陆成功', 'E', '登陆成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (8, 'S002', '0', 'S', '密码重置成功', 'E', '密码重置成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (9, 'F000', '0', 'F', '处理失败', 'E', '处理失败（默认）', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (10, 'F001', '0', 'F', '登陆失败，请联系管理员', 'E', '登陆失败，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (11, 'F002', '0', 'F', '异常登陆，请联系管理员', 'E', '异常登陆，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (12, 'F003', '0', 'F', '密码重置失败，请联系管理员', 'E', '密码重置失败，请联系管理员', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (13, 'S100', '1', 'S', '字典创建成功', 'E', '字典创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (14, 'S101', '1', 'S', '字典编辑成功', 'E', '字典编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (15, 'S102', '1', 'S', '字典排序成功', 'E', '字典排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (16, 'S103', '1', 'S', '字典封存成功', 'E', '字典封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (17, 'F100', '1', 'F', '字典查询失败', 'E', '字典查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (18, 'F101', '1', 'F', '字典分组下字典值唯一', 'E', '字典分组下字典值唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (19, 'F102', '1', 'F', '内置字典禁止刪除', 'E', '内置字典禁止刪除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (20, 'F103', '1', 'F', '字典排序失败', 'E', '字典排序失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (21, 'S200', '2', 'S', '响应码创建成功', 'E', '响应码创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (22, 'S201', '2', 'S', '响应码创建成功,实际响应码为{{code}}', 'E', '响应码创建成功,实际响应码为{{code}}', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (23, 'S202', '2', 'S', '响应码编辑成功', 'E', '响应码编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (24, 'S203', '2', 'S', '响应码封存成功', 'E', '响应码封存成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (25, 'F200', '2', 'F', '响应码查询失败', 'E', '响应码查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (26, 'F201', '2', 'F', '响应码全局唯一', 'E', '响应码全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (27, 'F202', '2', 'F', '内置响应码禁止删除', 'E', '内置响应码禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (28, 'S300', '3', 'S', '路由创建成功', 'E', '路由创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (29, 'S301', '3', 'S', '路由编辑成功', 'E', '路由编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (30, 'S302', '3', 'S', '路由删除成功', 'E', '路由删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (31, 'F300', '3', 'F', '路由查询失败', 'E', '路由查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (32, 'F301', '3', 'F', '路由地址全局唯一', 'E', '路由地址全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (33, 'F302', '3', 'F', '路由名称全局唯一', 'E', '路由名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (34, 'F303', '3', 'F', '内置路由禁止删除', 'E', '内置路由禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (35, 'F304', '3', 'F', '路由删除', 'E', '路由删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (36, 'S400', '4', 'S', '权限创建成功', 'E', '权限创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (37, 'S401', '4', 'S', '权限编辑成功', 'E', '权限编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (38, 'S402', '4', 'S', '权限删除成功', 'E', '权限删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (39, 'S403', '4', 'S', '权限排序成功', 'E', '权限排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (40, 'F400', '4', 'F', '权限查询失败', 'E', '权限查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (41, 'F401', '4', 'F', '权限别名全局唯一', 'E', '权限别名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (42, 'F402', '4', 'F', '权限名称全局唯一', 'E', '权限名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (43, 'F403', '4', 'F', '内置权限禁止删除', 'E', '内置权限禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (44, 'F404', '4', 'F', '权限配置路由同步失败', 'E', '权限配置路由同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (45, 'F405', '4', 'F', '权限配置删除失败', 'E', '权限配置删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (46, 'S500', '5', 'S', '角色创建成功', 'E', '角色创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (47, 'S501', '5', 'S', '角色编辑成功', 'E', '角色编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (48, 'S502', '5', 'S', '角色删除失败', 'E', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (49, 'F500', '5', 'F', '角色查询失败', 'E', '角色查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (50, 'F501', '5', 'F', '内置角色禁止编辑', 'E', '内置角色禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (51, 'F502', '5', 'F', '角色名全局唯一', 'E', '角色名全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (52, 'F503', '5', 'F', '角色权限配置失败', 'E', '角色权限配置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (53, 'F504', '5', 'F', '角色删除失败', 'E', '角色删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (54, 'S600', '6', 'S', '部门创建成功', 'E', '集团部门创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (55, 'S601', '6', 'S', '部门编辑成功', 'E', '集团部门编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (56, 'S602', '6', 'S', '部门删除成功', 'E', '集团部门删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (57, 'S603', '6', 'S', '部门排序成功', 'E', '集团部门排序成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (58, 'S604', '6', 'S', '部门迁移成功', 'E', '集团部门迁移成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (59, 'F600', '6', 'F', '部门查询失败', 'E', '集团部门查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (60, 'F601', '6', 'F', '部门名称全局唯一', 'E', '集团部门名称全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (61, 'F602', '6', 'F', '内置部门禁止删除', 'E', '内置集团部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (62, 'F603', '6', 'F', '部门删除失败', 'E', '集团部门删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (63, 'F604', '6', 'F', '部门存在子部门禁止删除', 'E', '集团部门存在子部门禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (64, 'F605', '6', 'F', '部门存在成员禁止删除', 'E', '集团部门存在成员禁止删除', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (65, 'F606', '6', 'F', '部门迁移失败', 'E', '集团部门迁移失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (66, 'S700', '7', 'S', '登陆账号创建成功', 'E', '登陆账号创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (67, 'S701', '7', 'S', '登陆账号编辑成功', 'E', '登陆账号编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (68, 'S702', '7', 'S', '登陆账号删除成功', 'E', '登陆账号删除成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (69, 'S703', '7', 'S', '登陆账号重置成功', 'E', '登陆账号重置成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (70, 'F700', '7', 'F', '登陆账号查询失败', 'E', '登陆账号查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (71, 'F701', '7', 'F', '登陆账号全局唯一', 'E', '登陆账号全局唯一', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (72, 'F702', '7', 'F', '账号角色同步失败', 'E', '账号角色同步失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (73, 'F703', '7', 'F', '特殊账号禁止编辑', 'E', '特殊账号禁止编辑', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (74, 'F704', '7', 'F', '登陆账号删除失败', 'E', '登陆账号删除失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (75, 'F705', '7', 'F', '登陆账号重置失败', 'E', '登陆账号重置失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (76, 'S800', '8', 'S', '系统配置编辑成功', 'E', '系统配置编辑成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (77, 'F800', '8', 'F', '系统配置查询失败', 'E', '系统配置查询失败', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (78, 'S900', '9', 'S', '文章分类创建成功', 'E', '文章分类创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (79, 'S901', '9', 'S', '文章分类编辑成功', 'E', '文章分类编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (80, 'S902', '9', 'S', '文章分类删除成功', 'E', '文章分类删除成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (81, 'S903', '9', 'S', '文章分类迁移成功', 'E', '文章分类迁移成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (82, 'F900', '9', 'F', '文章分类查询失败', 'E', '文章分类查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (83, 'F901', '9', 'F', '文章分类迁移失败', 'E', '文章分类迁移失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (84, 'F902', '9', 'F', '文章分类未提交图片', 'E', '文章分类未提交图片', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (85, 'F903', '9', 'F', '文章分类地址全局唯一', 'E', '文章分类地址全局唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (86, 'F904', '9', 'F', '文章分类下存在数据禁止删除', 'E', '文章分类下存在数据禁止删除', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (87, 'F905', '9', 'F', '文章分类删除失败', 'E', '文章分类删除失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (88, 'S1000', '10', 'S', '资源添加成功', 'E', '资源配置表创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (89, 'S1001', '10', 'S', '资源编辑成功', 'E', '资源配置表编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (90, 'S1002', '10', 'S', '资源清理暂不开放', 'E', '资源清理暂不开放', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (91, 'F1000', '10', 'F', '资源查询失败', 'E', '资源配置表查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (92, 'F1001', '10', 'F', '资源文件上传失败', 'E', '资源文件上传失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (93, 'S1100', '11', 'S', '文章标签创建成功', 'E', '标签创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (94, 'S1101', '11', 'S', '文章标签编辑成功', 'E', '标签编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (95, 'S1102', '11', 'S', '文章标签删除成功', 'E', '标签删除成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (96, 'F1100', '11', 'F', '文章标签查询失败', 'E', '标签查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (97, 'F1101', '11', 'F', '文章标签未提交图片', 'E', '文章标签未提交图片', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (98, 'F1102', '11', 'F', '文章标签地址全局唯一', 'E', '标签地址全局唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (99, 'F1103', '11', 'F', '文章标签下数据删除失败', 'E', '文章标签下数据删除失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (100, 'F1104', '11', 'F', '文章标签删除失败', 'E', '标签删除成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (101, 'S1200', '12', 'S', '文章专题创建成功', 'E', '文章专题创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (102, 'S1201', '12', 'S', '文章专题编辑成功', 'E', '文章专题编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (103, 'S1202', '12', 'S', '文章专题删除成功', 'E', '文章专题删除成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (104, 'F1200', '12', 'F', '文章专题查询失败', 'E', '文章专题查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (105, 'F1201', '12', 'F', '文章标签未提交图片', 'E', '文章标签未提交图片', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (106, 'F1202', '12', 'F', '文章专题内文章应当唯一', 'E', '文章专题内文章应当唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (107, 'F1203', '12', 'F', '文章专题关联文章同步失败', 'E', '文章专题关联文章同步失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (108, 'F1204', '12', 'F', '文章专题删除', 'E', '文章专题删除', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (109, 'S1300', '13', 'S', '文章或页面创建成功', 'E', '文章或页面创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (110, 'S1301', '13', 'S', '文章或页面编辑成功', 'E', '文章或页面编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (111, 'F1300', '13', 'F', '文章或页面查询失败', 'E', '文章或页面查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (112, 'F1301', '13', 'F', '文章或页面必须有主图', 'E', '文章或页面必须有主图', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (113, 'F1302', '13', 'F', '文章关联标签更新失败', 'E', '文章关联标签更新失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (114, 'F1303', '13', 'F', '文章关联资源更新失败', 'E', '文章关联资源更新失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (115, 'F1304', '13', 'F', '文章关联详情更新失败', 'E', '文章关联详情更新失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (116, 'F1305', '13', 'F', '文章或页面地址全局唯一', 'E', '文章或页面地址全局唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (117, 'F1306', '13', 'F', '文章关联地址全局唯一', 'E', '文章关联地址全局唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (118, 'S1400', '14', 'S', 'Banner创建成功', 'E', 'Banner创建成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (119, 'S1401', '14', 'S', 'Banner编辑成功', 'E', 'Banner编辑成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (120, 'S1402', '14', 'S', 'Banner删除成功', 'E', 'Banner删除成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (121, 'F1400', '14', 'F', 'Banner查询失败', 'E', 'Banner查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (122, 'F1401', '14', 'F', 'Banner未提交图片', 'E', 'Banner未提交图片', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (123, 'F1402', '14', 'F', 'Banner地址全局唯一', 'E', 'Banner地址全局唯一', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (124, 'F1403', '14', 'F', 'Banner删除失败', 'E', 'Banner删除失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (125, 'S1500', '15', 'S', '评论更新成功', 'E', '评论更新成功', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (126, 'F1500', '15', 'F', '评论查询失败', 'E', '评论查询失败', '0', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (127, 'S1600', '16', 'S', '友情链接创建成功', 'E', '友情链接创建成功', '1', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (128, 'S1601', '16', 'S', '友情链接编辑成功', 'E', '友情链接编辑成功', '2', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (129, 'S1602', '16', 'S', '友情链接删除成功', 'E', '友情链接删除成功', '3', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (130, 'F1600', '16', 'F', '友情链接查询失败', 'E', '友情链接查询失败', '4', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (131, 'F1601', '16', 'F', '友情链接未提交图片', 'E', '友情链接未提交图片', '5', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (132, 'F1602', '16', 'F', '友情链接地址全局唯一', 'E', '友情链接地址全局唯一', '6', '0', NULL, NULL);
INSERT INTO `response_code` VALUES (133, 'F1603', '16', 'F', '友情链接删除失败', 'E', '友情链接删除失败', '7', '0', NULL, NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
                         `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                         `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
                         `remark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色备注',
                         `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
                         `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                         `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                         `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `name_uni`(`name`) USING BTREE COMMENT '角色名称唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (5, '管理员', '', '0', '0', '2024-06-29 13:03:06', '2025-05-28 11:14:04');

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
                                    `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                                    `role_id` bigint(0) NOT NULL COMMENT '角色ID',
                                    `permission_id` bigint(0) NOT NULL COMMENT '权限ID',
                                    `check_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '选择类型 0 check 1 halfCheck',
                                    PRIMARY KEY (`id`) USING BTREE,
                                    UNIQUE INDEX `role_permission_uni`(`role_id`, `permission_id`) USING BTREE COMMENT '角色权限全局唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1104 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (1049, 5, 1, '0');
INSERT INTO `role_permission` VALUES (1050, 5, 2, '0');
INSERT INTO `role_permission` VALUES (1051, 5, 3, '0');
INSERT INTO `role_permission` VALUES (1052, 5, 11, '0');
INSERT INTO `role_permission` VALUES (1053, 5, 4, '0');
INSERT INTO `role_permission` VALUES (1054, 5, 9, '0');
INSERT INTO `role_permission` VALUES (1055, 5, 44, '0');
INSERT INTO `role_permission` VALUES (1056, 5, 13, '0');
INSERT INTO `role_permission` VALUES (1057, 5, 14, '0');
INSERT INTO `role_permission` VALUES (1058, 5, 15, '0');
INSERT INTO `role_permission` VALUES (1059, 5, 16, '0');
INSERT INTO `role_permission` VALUES (1060, 5, 12, '0');
INSERT INTO `role_permission` VALUES (1061, 5, 45, '0');
INSERT INTO `role_permission` VALUES (1062, 5, 17, '0');
INSERT INTO `role_permission` VALUES (1063, 5, 18, '0');
INSERT INTO `role_permission` VALUES (1064, 5, 19, '0');
INSERT INTO `role_permission` VALUES (1065, 5, 21, '0');
INSERT INTO `role_permission` VALUES (1066, 5, 20, '0');
INSERT INTO `role_permission` VALUES (1067, 5, 7, '0');
INSERT INTO `role_permission` VALUES (1068, 5, 46, '0');
INSERT INTO `role_permission` VALUES (1069, 5, 22, '0');
INSERT INTO `role_permission` VALUES (1070, 5, 23, '0');
INSERT INTO `role_permission` VALUES (1071, 5, 24, '0');
INSERT INTO `role_permission` VALUES (1072, 5, 10, '0');
INSERT INTO `role_permission` VALUES (1073, 5, 47, '0');
INSERT INTO `role_permission` VALUES (1074, 5, 25, '0');
INSERT INTO `role_permission` VALUES (1075, 5, 26, '0');
INSERT INTO `role_permission` VALUES (1076, 5, 27, '0');
INSERT INTO `role_permission` VALUES (1077, 5, 28, '0');
INSERT INTO `role_permission` VALUES (1078, 5, 29, '0');
INSERT INTO `role_permission` VALUES (1079, 5, 48, '0');
INSERT INTO `role_permission` VALUES (1080, 5, 30, '0');
INSERT INTO `role_permission` VALUES (1081, 5, 31, '0');
INSERT INTO `role_permission` VALUES (1082, 5, 32, '0');
INSERT INTO `role_permission` VALUES (1083, 5, 33, '0');
INSERT INTO `role_permission` VALUES (1084, 5, 49, '0');
INSERT INTO `role_permission` VALUES (1085, 5, 36, '0');
INSERT INTO `role_permission` VALUES (1086, 5, 37, '0');
INSERT INTO `role_permission` VALUES (1087, 5, 38, '0');
INSERT INTO `role_permission` VALUES (1088, 5, 34, '0');
INSERT INTO `role_permission` VALUES (1089, 5, 50, '0');
INSERT INTO `role_permission` VALUES (1090, 5, 39, '0');
INSERT INTO `role_permission` VALUES (1091, 5, 40, '0');
INSERT INTO `role_permission` VALUES (1092, 5, 41, '0');
INSERT INTO `role_permission` VALUES (1093, 5, 42, '0');
INSERT INTO `role_permission` VALUES (1094, 5, 35, '0');
INSERT INTO `role_permission` VALUES (1095, 5, 52, '0');
INSERT INTO `role_permission` VALUES (1096, 5, 43, '0');
INSERT INTO `role_permission` VALUES (1097, 5, 53, '0');
INSERT INTO `role_permission` VALUES (1098, 5, 56, '0');
INSERT INTO `role_permission` VALUES (1099, 5, 57, '0');
INSERT INTO `role_permission` VALUES (1100, 5, 58, '0');
INSERT INTO `role_permission` VALUES (1101, 5, 59, '0');
INSERT INTO `role_permission` VALUES (1102, 5, 60, '0');
INSERT INTO `role_permission` VALUES (1103, 5, 61, '0');
INSERT INTO `role_permission` VALUES (1104, 5, 62, '0');

-- ----------------------------
-- Table structure for router
-- ----------------------------
DROP TABLE IF EXISTS `router`;
CREATE TABLE `router`  (
                           `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '默认数据ID',
                           `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由名称，用于界面展示，与权限关联',
                           `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容',
                           `type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '免授权路由 0 授权 1 免授权',
                           `service_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务编码（字典），为接口分组',
                           `log_in_db` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '日志入库 0 入库 1 不入库',
                           `req_log_print` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '请求日志打印 0 打印 1 不打印',
                           `req_log_secure` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求日志脱敏字段，逗号分隔',
                           `res_log_print` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '响应日志打印 0 打印 1 不打印',
                           `res_log_secure` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '响应日志脱敏字段，逗号分隔',
                           `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                           `mark` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '变更标识 0可变更1禁止变更',
                           `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态 0正常 1锁定 2封存',
                           `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                           `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '路由地址唯一',
                           INDEX `name_uni`(`name`) USING BTREE COMMENT '路由名称唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 108 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口路由' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router
-- ----------------------------
INSERT INTO `router` VALUES (1, '账号密码登陆', '/auth/login', '1', '0', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (2, '账号密码重置', '/auth/reset', '0', '0', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (3, '通用API示例', '/docs/sample', '0', '0', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (4, '登陆账号新建', '/plat/account/add', '0', '5', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (5, '登陆账号移除', '/plat/account/del', '0', '5', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (6, '登陆账号编辑', '/plat/account/edit', '0', '5', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (7, '登陆账号详情', '/plat/account/get', '0', '5', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:33:05');
INSERT INTO `router` VALUES (8, '登陆账号分页', '/plat/account/page', '0', '5', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (9, '登陆账号重置', '/plat/account/reset', '0', '5', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (10, '集团部门新建', '/plat/dept/add', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (11, '集团同级部门', '/plat/dept/bro', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (12, '集团部门移除', '/plat/dept/del', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (13, '集团部门编辑', '/plat/dept/edit', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (14, '集团部门详情', '/plat/dept/get', '0', '7', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:59');
INSERT INTO `router` VALUES (15, '集团部门排序', '/plat/dept/sort', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (16, '集团部门迁移', '/plat/dept/to', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (17, '集团部门树', '/plat/dept/tree', '0', '7', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (18, '字典新建', '/plat/dict/add', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (19, '字典同级查询', '/plat/dict/bro', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (20, '字典封存', '/plat/dict/del', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (21, '字典编辑', '/plat/dict/edit', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (22, '字典详情', '/plat/dict/get', '0', '1', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:44');
INSERT INTO `router` VALUES (23, '字典NextVal建议', '/plat/dict/nextVal', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (24, '字典分页', '/plat/dict/page', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (25, '读取字典', '/plat/dict/read', '0', '1', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:33:42');
INSERT INTO `router` VALUES (26, '排序处理', '/plat/dict/sort', '0', '1', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (27, '权限新建', '/plat/permission/add', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (28, '同级权限', '/plat/permission/bro', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (29, '权限封存', '/plat/permission/del', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (30, '权限编辑', '/plat/permission/edit', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (31, '权限详情', '/plat/permission/get', '0', '4', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:40');
INSERT INTO `router` VALUES (32, '权限排序', '/plat/permission/sort', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (33, '权限树', '/plat/permission/tree', '0', '4', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (34, '响应码新建', '/plat/response/add', '0', '2', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (35, '响应码封存', '/plat/response/del', '0', '2', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (36, '响应码编辑', '/plat/response/edit', '0', '2', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (37, '响应码详情', '/plat/response/get', '0', '2', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:34');
INSERT INTO `router` VALUES (38, '响应码NextVal建议', '/plat/response/nextVal', '0', '2', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (39, '响应码分页', '/plat/response/page', '0', '2', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (40, '角色新建', '/plat/role/add', '0', '6', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (41, '角色删除', '/plat/role/del', '0', '6', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (42, '角色编辑', '/plat/role/edit', '0', '6', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (43, '角色详情', '/plat/role/get', '0', '6', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:29');
INSERT INTO `router` VALUES (44, '角色列表', '/plat/role/list', '0', '6', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (45, '角色分页', '/plat/role/page', '0', '6', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (46, '路由接口新建', '/plat/router/add', '0', '3', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (47, '路由接口封存', '/plat/router/del', '0', '3', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (48, '路由接口编辑', '/plat/router/edit', '0', '3', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (49, '路由接口详情', '/plat/router/get', '0', '3', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:23');
INSERT INTO `router` VALUES (50, '路由接口分页', '/plat/router/page', '0', '3', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (51, '系统配置编辑', '/plat/sysConfig/edit', '0', '8', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (52, '系统配置详情', '/plat/sysConfig/get', '0', '8', '1', '1', '', '1', '', '', '1', '0', NULL, '2024-07-31 21:32:17');
INSERT INTO `router` VALUES (53, '账号登出', '/auth/logout', '1', '0', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (54, '账号权限信息', '/auth/mine', '0', '0', '1', '1', '', '1', '', '', '0', '0', '2024-08-23 14:06:48', '2024-08-23 14:06:48');
INSERT INTO `router` VALUES (55, '链接', '/blog/banner/add', '0', '14', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (56, '链接', '/blog/banner/del', '0', '14', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (57, '链接', '/blog/banner/edit', '0', '14', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (58, '链接', '/blog/banner/get', '0', '14', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (59, '链接', '/blog/banner/page', '0', '14', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (60, '分类', '/blog/category/add', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (61, '分类', '/blog/category/del', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (62, '分类', '/blog/category/edit', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (63, '分类', '/blog/category/get', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (64, '分类', '/blog/category/list', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (65, '分类', '/blog/category/merge', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (66, '分类', '/blog/category/page', '0', '9', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (67, '博客', '/blog/center/index', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (68, '友链', '/blog/links/add', '0', '16', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (69, '友链', '/blog/links/del', '0', '16', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (70, '友链', '/blog/links/edit', '0', '16', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (71, '友链', '/blog/links/get', '0', '16', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (72, '友链', '/blog/links/page', '0', '16', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (73, '文章', '/blog/post/add', '0', '13', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (74, '文章', '/blog/post/edit', '0', '13', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (75, '文章', '/blog/post/get', '0', '13', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (76, '文章', '/blog/post/page', '0', '13', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (77, '评论', '/blog/postComment/edit', '0', '15', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (78, '评论', '/blog/postComment/get', '0', '15', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (79, '评论', '/blog/postComment/page', '0', '15', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (80, '资源', '/blog/source/add', '0', '10', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (81, '资源', '/blog/source/del', '0', '10', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (82, '资源', '/blog/source/edit', '0', '10', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (83, '资源', '/blog/source/get', '0', '10', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (84, '资源', '/blog/source/page', '0', '10', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (85, '标签', '/blog/tag/add', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (86, '标签', '/blog/tag/del', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (87, '标签', '/blog/tag/edit', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (88, '标签', '/blog/tag/get', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (89, '标签', '/blog/tag/list', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (90, '标签', '/blog/tag/page', '0', '11', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (91, '专题', '/blog/topic/add', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (92, '专题', '/blog/topic/del', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (93, '专题', '/blog/topic/edit', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (94, '专题', '/blog/topic/get', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (95, '专题', '/blog/topic/page', '0', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (96, '博客页面', '/page/categoryList', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (97, '博客页面', '/page/comm/add', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (98, '博客页面', '/page/comments', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (99, '博客页面', '/page/index', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (100, '博客页面', '/page/link/add', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (101, '博客页面', '/page/link/scan', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (102, '博客页面', '/page/page', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (103, '博客页面', '/page/pages', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (104, '博客页面', '/page/post', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (105, '博客页面', '/page/post/good', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (106, '博客页面', '/page/posts', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (107, '博客页面', '/page/tags', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);
INSERT INTO `router` VALUES (108, '博客页面', '/page/link', '1', '12', '1', '1', NULL, '1', NULL, NULL, '1', '0', NULL, NULL);


-- ----------------------------
-- Table structure for router_log
-- ----------------------------
DROP TABLE IF EXISTS `router_log`;
CREATE TABLE `router_log`  (
                               `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                               `app_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用名',
                               `app_node` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用节点',
                               `app_trace_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用节点TraceID',
                               `req_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源IP',
                               `req_url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求地址',
                               `req_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求报文',
                               `req_at` datetime(0) NULL DEFAULT NULL COMMENT '请求时间',
                               `res_status` smallint(0) NULL DEFAULT NULL COMMENT '响应状态',
                               `res_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应报文',
                               `res_time` datetime(0) NULL DEFAULT NULL COMMENT '响应时间',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `log_uni`(`app_name`, `app_node`, `app_trace_id`) USING BTREE COMMENT '日志唯一性'
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '路由接口日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router_log
-- ----------------------------

-- ----------------------------
-- Table structure for source
-- ----------------------------
DROP TABLE IF EXISTS `source`;
CREATE TABLE `source`  (
                           `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                           `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源名',
                           `back_end` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片后缀，jpg png gif',
                           `file_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片类型，枚举：0_主图 1_图标 2_正文',
                           `file_path` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件子目录，年份目录',
                           `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                           `version` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件当前版本号',
                           `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                           `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of source
-- ----------------------------
-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
                               `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                               `login_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '并发限制开关，0限制 1不限制',
                               `login_num` int(0) NULL DEFAULT NULL COMMENT '最大登陆并发量，最小为1',
                               `login_fail_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败限制开关，0限制 1不限制',
                               `login_fail_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败限制 1秒 2分 3时 4天 ',
                               `login_fail_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败最大尝试次数',
                               `login_fail_lock_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆失败锁定 1秒 2分 3时 4天 ',
                               `login_fail_lock_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败锁定时长',
                               `login_fail_try_num` int(0) NULL DEFAULT NULL COMMENT '登陆失败尝试次数',
                               `logout_switch` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆过期开关，0限制 1不限制',
                               `logout_unit` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆过期单位，0永不过期 1秒 2分 3时 4天 ',
                               `logout_num` int(0) NULL DEFAULT NULL COMMENT '登陆过期长度数量',
                               `cdn_full_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件CDN地址（暂不支持）',
                               `file_full_path` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件存放根目录',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '0', 1, '0', '2', 5, '3', 12, 5, '0', '6', 1, NULL, 'D:\\Mebugs\\mebugs\\admin-ui\\public\\source\\');

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
                        `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                        `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
                        `url` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签地址',
                        `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                        `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                        `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                        `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                        `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`  (
                          `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                          `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
                          `url` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '专题地址',
                          `summary` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                          `source_id` bigint(0) NULL DEFAULT NULL COMMENT '主图资源ID',
                          `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_锁定 2_封存',
                          `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                          `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `url_uni`(`url`) USING BTREE COMMENT '地址唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章专题' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of topic
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
                         `client_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户端ID（前端生成，清除会丢失）',
                         `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
                         `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                         `summary` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
                         `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户编码（允许访问用户）默认=ID',
                         `third_url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '三方网站（需要审核）',
                         `wait_url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '待审核三方URL',
                         `source_id` bigint(0) NULL DEFAULT NULL COMMENT '头像资源ID（仅小程序提供）',
                         `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '小程序账号绑定登录（允许授权转移）',
                         `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '状态，枚举：0_正常 1_待审核 2_封存',
                         `create_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
                         `update_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `client_uni`(`client_id`) USING BTREE COMMENT '客户端ID唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '极简用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

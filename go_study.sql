/*
 Navicat Premium Dump SQL

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : go_study

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 15/11/2024 17:11:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `user_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '100' COMMENT '用户类型：(100系统用户)',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '手机',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户邮箱',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户头像',
  `signed` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '个人签名',
  `dashboard` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '后台首页类型',
  `status` bigint NULL DEFAULT 1 COMMENT '状态 (1正常 2停用)',
  `login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登陆IP',
  `backend_setting` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '后台设置数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `super` int NOT NULL DEFAULT 2 COMMENT '是否超级管理员 (1是 2否)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username_deletedat`(`username` ASC) USING BTREE,
  INDEX `idx_admin_users_user_name`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
INSERT INTO `admin_users` VALUES (1, '2024-07-11 00:00:00.000', '2024-08-19 17:25:46.000', 'superAdmin', '$2a$10$H6/FynuF4p0BT2jd4aAEIuiY3MMIr1cPdzJKZ5iMDg4/dCPxNrHz2', '100', 'superAdmin', '13888888888', 'admin@qq.com', '', '行胜于言，质胜于华。', 'statistics', 1, '127.0.0.1', '{\"mode\":\"light\",\"tag\":true,\"menuCollapse\":false,\"menuWidth\":230,\"layout\":\"classic\",\"skin\":\"mine\",\"i18n\":false,\"language\":\"zh_CN\",\"animation\":\"ma-slide-down\",\"color\":\"#165dff\",\"ws\":false}', '', 1);
INSERT INTO `admin_users` VALUES (30, '2024-11-15 16:52:41.748', '2024-11-15 16:52:41.748', 'admin', '$2a$10$Vt.bJlPDD3cSDnLOVsL2..46jcoQ2uPug/k6lZZ7V/Mk5hZa0J2GC', '100', '', '', '', '', '', '', 1, '', '', '', 2);

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `status` smallint NOT NULL DEFAULT 1 COMMENT '状态 (1启用 2禁用)',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公告内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articles
-- ----------------------------
INSERT INTO `articles` VALUES (13, '2024-10-31 16:40:40.371', '2024-11-04 14:17:37.689', '哈哈哈哈1', 1, '<p>哈哈哈哈哈哈哈哈哈哈哈哈1</p>');
INSERT INTO `articles` VALUES (14, '2024-10-31 16:41:04.846', '2024-11-15 16:51:24.806', '文章', 2, '<p>哈哈哈哈2哈哈哈哈2哈哈哈哈2哈哈哈哈2哈哈哈哈2</p>');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 450 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (397, 'g', '30', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (436, 'p', 'admin', '/backend/articles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (433, 'p', 'admin', '/backend/articles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (434, 'p', 'admin', '/backend/articles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (435, 'p', 'admin', '/backend/articles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (437, 'p', 'admin', '/backend/articles/change', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (425, 'p', 'admin', '/backend/buttons', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (424, 'p', 'admin', '/backend/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (429, 'p', 'admin', '/backend/notices', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (426, 'p', 'admin', '/backend/notices', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (427, 'p', 'admin', '/backend/notices', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (428, 'p', 'admin', '/backend/notices', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (430, 'p', 'admin', '/backend/notices/change', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (447, 'p', 'admin', '/backend/roles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (444, 'p', 'admin', '/backend/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (445, 'p', 'admin', '/backend/roles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (446, 'p', 'admin', '/backend/roles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (449, 'p', 'admin', '/backend/roles/get/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (448, 'p', 'admin', '/backend/roles/set/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (431, 'p', 'admin', '/backend/upload/local/img', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (432, 'p', 'admin', '/backend/upload/local/video', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (441, 'p', 'admin', '/backend/users', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (438, 'p', 'admin', '/backend/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (439, 'p', 'admin', '/backend/users', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (440, 'p', 'admin', '/backend/users', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (442, 'p', 'admin', '/backend/users/get/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (443, 'p', 'admin', '/backend/users/set/role', 'POST', '', '', '');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由访问路径',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由name',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由重定向地址',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求接口地址',
  `act` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '视图文件路径',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单和面包屑对应的图标',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由标题(菜单名称)',
  `activeMenu` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '是否在菜单中隐藏,需要高亮的path',
  `isLink` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由外链时填写的访问地址',
  `isHide` smallint NOT NULL DEFAULT 2 COMMENT '是否在菜单中隐藏 (1是 2否)',
  `isFull` smallint NOT NULL DEFAULT 2 COMMENT '菜单是否全屏 (1是 2否)',
  `isAffix` smallint NOT NULL DEFAULT 1 COMMENT '菜单是否固定在标签页中 (1是 2否)',
  `isKeepAlive` smallint NOT NULL DEFAULT 2 COMMENT '当前路由是否缓存 (1是 2否)',
  `status` smallint NOT NULL DEFAULT 1 COMMENT '状态 (1正常 2停用)',
  `type` smallint NOT NULL DEFAULT 1 COMMENT '类型 (1菜单2按钮3外链4Iframe)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 62 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, NULL, '2024-11-04 14:07:32.892', 0, '/home/index', 'home', '', '', 'GET', '/home/index', 'HomeFilled', '首页', '', '', 2, 2, 1, 2, 1, 1, 99);
INSERT INTO `menus` VALUES (3, NULL, NULL, 0, '/cms', 'cms', '/cms/notice', '', '', '', 'Lollipop', '内容管理', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (4, NULL, '2024-11-04 11:32:35.063', 3, '/cms/notice', 'notice', '', '/backend/notices', 'GET', '/cms/notice/index', 'menu', '公告管理', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (5, NULL, NULL, 4, '/cms/notice/detail/:id', 'noticeDetail', '', '', '', '/cms/notice/detail', 'menu', '公告 详情', '/cms/notice', '', 1, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (6, NULL, '2024-11-04 11:31:44.123', 4, '', 'create', '', '/backend/notices', 'POST', '', '', '添加', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (7, NULL, '2024-11-04 11:33:44.041', 4, '', 'update', '', '/backend/notices', 'PUT', '', '', '编辑', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (8, NULL, '2024-11-04 11:33:47.909', 4, '', 'delete', '', '/backend/notices', 'DELETE', '', '', '删除', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (10, NULL, '2024-11-04 11:40:03.682', 4, '', 'change', '', '/backend/notices/change', 'POST', '', '', '修改状态', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (12, NULL, '2024-09-02 15:59:39.192', 0, '/authority', 'authority', '/authority/menu', '', '', '', 'Lock', '权限管理', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (13, NULL, '2024-09-02 17:18:24.867', 12, '/authority/menu', 'menuAuthority', '', '', '', '/authority/menu/index', 'menu', '菜单权限', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (30, '2024-09-02 10:35:03.467', '2024-11-04 11:46:41.797', 12, '/authority/admin_user', 'adminUser', '', '/backend/users', 'GET', '/authority/admin_user/index', 'Avatar', '管理员', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (31, '2024-09-02 16:10:02.983', '2024-11-04 11:45:03.431', 12, '/authority/role', 'role', '', '/backend/roles', 'GET', '/authority/role/index', 'Stamp', '角色管理', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (34, '2024-09-03 10:28:22.793', '2024-11-04 11:31:23.293', 1, '', 'none', '', '/backend/menus', 'GET', '', '', '用户菜单数据源', '', '', 1, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (35, '2024-09-03 10:28:53.299', '2024-11-04 11:31:26.915', 1, '', 'none', '', '/backend/buttons', 'GET', '', '', '用户按钮数据源', '', '', 1, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (36, NULL, '2024-11-04 11:46:53.547', 13, '', 'create', '', '/backend/menus', 'POST', '', '', '添加', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (37, NULL, '2024-11-04 11:48:18.520', 13, '', 'update', '', '/backend/menus', 'PUT', '', '', '编辑', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (38, NULL, '2024-11-04 11:48:21.419', 13, '', 'delete', '', '/backend/menus', 'DELETE', '', '', '删除', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (42, NULL, '2024-11-04 11:46:38.651', 30, '', 'create', '', '/backend/users', 'POST', '', '', '添加', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (43, NULL, '2024-11-04 11:46:35.643', 30, '', 'update', '', '/backend/users', 'PUT', '', '', '编辑', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (44, NULL, '2024-11-04 11:46:32.784', 30, '', 'delete', '', '/backend/users', 'DELETE', '', '', '删除', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (45, NULL, '2024-11-04 11:46:28.277', 30, '', 'role', '', '/backend/users/get/role', 'POST', '', '', '查看角色', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (46, NULL, '2024-11-04 11:46:25.212', 30, '', 'role', '', '/backend/users/set/role', 'POST', '', '', '设置角色', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (47, NULL, '2024-11-04 11:45:09.579', 31, '', 'create', '', '/backend/roles', 'POST', '', '', '添加', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (48, NULL, '2024-11-04 13:07:14.315', 31, '', 'update', '', '/backend/roles', 'PUT', '', '', '编辑', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (49, NULL, '2024-11-04 13:07:18.462', 31, '', 'delete', '', '/backend/roles', 'DELETE', '', '', '删除', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (50, NULL, '2024-11-04 11:46:02.201', 31, '', 'authority', '', '/backend/roles/set/menu', 'POST', '', '', '查看角色', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (51, NULL, '2024-11-04 11:46:05.540', 31, '', 'authority', '', '/backend/roles/get/menu', 'POST', '', '', '设置角色', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (52, NULL, '2024-11-04 11:34:18.637', 3, '/cms/upload', 'upload', '', '', '', '/cms/upload/index', 'upload', '文件上传', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (53, NULL, '2024-11-04 14:13:59.287', 52, '', 'upload-img', '', '/backend/upload/local/img', 'POST', '', '', '图片上传', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (54, NULL, '2024-11-04 14:14:07.684', 52, '', 'upload-video', '', '/backend/upload/local/video', 'POST', '', '', '视频上传', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (55, '2024-10-31 16:31:36.808', '2024-11-04 11:35:23.780', 3, '/cms/article', 'article', '', '/backend/articles', 'GET', '/cms/article/index', 'Burger', '文章示例', '', '', 2, 2, 2, 2, 1, 1, 0);
INSERT INTO `menus` VALUES (56, NULL, '2024-11-04 11:35:27.869', 55, '', 'create', '', '/backend/articles', 'POST', '', '', '添加', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (57, NULL, '2024-11-04 11:35:38.638', 55, '', 'update', '', '/backend/articles', 'PUT', '', '', '编辑', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (58, NULL, '2024-11-04 11:35:42.685', 55, '', 'delete', '', '/backend/articles', 'DELETE', '', '', '删除', '', '', 1, 2, 2, 2, 1, 2, 0);
INSERT INTO `menus` VALUES (60, NULL, '2024-11-04 14:14:50.351', 55, '', 'change', '', '/backend/articles/change', 'POST', '', '', '修改状态', '', '', 1, 2, 2, 2, 1, 2, 0);

-- ----------------------------
-- Table structure for notices
-- ----------------------------
DROP TABLE IF EXISTS `notices`;
CREATE TABLE `notices`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `status` smallint NOT NULL DEFAULT 1 COMMENT '状态 (1启用 2禁用)',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公告内容',
  `deleted_at` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `udx_title`(`title` ASC, `deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notices
-- ----------------------------
INSERT INTO `notices` VALUES (1, '2024-11-04 14:42:30.404', '2024-11-04 14:42:30.404', '111', 1, '<p>111111111</p>', 1730702636);
INSERT INTO `notices` VALUES (5, '2024-11-04 14:43:59.083', '2024-11-15 16:42:39.125', '333333333', 1, '<p>333333333333333<img src=\"https://asqy-1318214684.cos.ap-mumbai.myqcloud.com/1731660150860742400_5213.jpg\" alt=\"\" data-href=\"\" style=\"\"/></p>', 0);

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `role_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menu_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1353 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES (1322, '2024-11-15 16:55:21.805', '2024-11-15 16:55:21.805', 10, 1);
INSERT INTO `role_menus` VALUES (1323, '2024-11-15 16:55:21.807', '2024-11-15 16:55:21.807', 10, 34);
INSERT INTO `role_menus` VALUES (1324, '2024-11-15 16:55:21.810', '2024-11-15 16:55:21.810', 10, 35);
INSERT INTO `role_menus` VALUES (1325, '2024-11-15 16:55:21.811', '2024-11-15 16:55:21.811', 10, 3);
INSERT INTO `role_menus` VALUES (1326, '2024-11-15 16:55:21.814', '2024-11-15 16:55:21.814', 10, 4);
INSERT INTO `role_menus` VALUES (1327, '2024-11-15 16:55:21.815', '2024-11-15 16:55:21.815', 10, 5);
INSERT INTO `role_menus` VALUES (1328, '2024-11-15 16:55:21.818', '2024-11-15 16:55:21.818', 10, 6);
INSERT INTO `role_menus` VALUES (1329, '2024-11-15 16:55:21.820', '2024-11-15 16:55:21.820', 10, 7);
INSERT INTO `role_menus` VALUES (1330, '2024-11-15 16:55:21.822', '2024-11-15 16:55:21.822', 10, 8);
INSERT INTO `role_menus` VALUES (1331, '2024-11-15 16:55:21.824', '2024-11-15 16:55:21.824', 10, 10);
INSERT INTO `role_menus` VALUES (1332, '2024-11-15 16:55:21.826', '2024-11-15 16:55:21.826', 10, 52);
INSERT INTO `role_menus` VALUES (1333, '2024-11-15 16:55:21.828', '2024-11-15 16:55:21.828', 10, 53);
INSERT INTO `role_menus` VALUES (1334, '2024-11-15 16:55:21.831', '2024-11-15 16:55:21.831', 10, 54);
INSERT INTO `role_menus` VALUES (1335, '2024-11-15 16:55:21.833', '2024-11-15 16:55:21.833', 10, 55);
INSERT INTO `role_menus` VALUES (1336, '2024-11-15 16:55:21.835', '2024-11-15 16:55:21.835', 10, 56);
INSERT INTO `role_menus` VALUES (1337, '2024-11-15 16:55:21.837', '2024-11-15 16:55:21.837', 10, 57);
INSERT INTO `role_menus` VALUES (1338, '2024-11-15 16:55:21.840', '2024-11-15 16:55:21.840', 10, 58);
INSERT INTO `role_menus` VALUES (1339, '2024-11-15 16:55:21.842', '2024-11-15 16:55:21.842', 10, 60);
INSERT INTO `role_menus` VALUES (1340, '2024-11-15 16:55:21.843', '2024-11-15 16:55:21.843', 10, 12);
INSERT INTO `role_menus` VALUES (1341, '2024-11-15 16:55:21.846', '2024-11-15 16:55:21.846', 10, 30);
INSERT INTO `role_menus` VALUES (1342, '2024-11-15 16:55:21.848', '2024-11-15 16:55:21.848', 10, 42);
INSERT INTO `role_menus` VALUES (1343, '2024-11-15 16:55:21.851', '2024-11-15 16:55:21.851', 10, 43);
INSERT INTO `role_menus` VALUES (1344, '2024-11-15 16:55:21.853', '2024-11-15 16:55:21.853', 10, 44);
INSERT INTO `role_menus` VALUES (1345, '2024-11-15 16:55:21.855', '2024-11-15 16:55:21.855', 10, 45);
INSERT INTO `role_menus` VALUES (1346, '2024-11-15 16:55:21.858', '2024-11-15 16:55:21.858', 10, 46);
INSERT INTO `role_menus` VALUES (1347, '2024-11-15 16:55:21.860', '2024-11-15 16:55:21.860', 10, 31);
INSERT INTO `role_menus` VALUES (1348, '2024-11-15 16:55:21.863', '2024-11-15 16:55:21.863', 10, 47);
INSERT INTO `role_menus` VALUES (1349, '2024-11-15 16:55:21.865', '2024-11-15 16:55:21.865', 10, 48);
INSERT INTO `role_menus` VALUES (1350, '2024-11-15 16:55:21.868', '2024-11-15 16:55:21.868', 10, 49);
INSERT INTO `role_menus` VALUES (1351, '2024-11-15 16:55:21.870', '2024-11-15 16:55:21.870', 10, 50);
INSERT INTO `role_menus` VALUES (1352, '2024-11-15 16:55:21.872', '2024-11-15 16:55:21.872', 10, 51);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色代码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_roles_code`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (10, '2024-11-04 15:57:19.174', '2024-11-04 15:57:19.174', '管理员', 'admin');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `pid` bigint NOT NULL DEFAULT 0 COMMENT '父ID',
  `real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `id_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

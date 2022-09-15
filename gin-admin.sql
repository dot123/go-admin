/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS `gin-admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gin-admin`;

CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(512) DEFAULT NULL,
  `v0` varchar(512) DEFAULT NULL,
  `v1` varchar(512) DEFAULT NULL,
  `v2` varchar(512) DEFAULT NULL,
  `v3` varchar(512) DEFAULT NULL,
  `v4` varchar(512) DEFAULT NULL,
  `v5` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(256) DEFAULT NULL,
  `url` varchar(256) DEFAULT NULL,
  `tag` varchar(256) DEFAULT NULL,
  `key` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `notice` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '''创建时间''',
  `start_time` datetime DEFAULT NULL COMMENT '''开始时间''',
  `end_time` datetime DEFAULT NULL COMMENT '''结束时间''',
  `title` varchar(256) DEFAULT NULL COMMENT '''标题''',
  `content` varchar(256) DEFAULT NULL COMMENT '''内容''',
  `operator` varchar(256) DEFAULT NULL COMMENT '''操作者''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `sys_api` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `handle` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `type` varchar(16) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `type`, `action`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES
	(1, 'github.com/dot123/gin-gorm-admin/api/v1.(*AuthApi).Login-fm', '登陆', '/api/v1/login', '', 'POST', '2022-09-14 14:15:42.965', '2022-09-14 14:15:42.965', '0000-00-00 00:00:00.000', 0, 0),
	(2, 'github.com/dot123/gin-gorm-admin/api/v1.(*AuthApi).Logout-fm', '登出', '/api/v1/logout', '', 'POST', '2022-09-14 14:15:43.012', '2022-09-14 14:15:43.012', '0000-00-00 00:00:00.000', 0, 0),
	(3, 'github.com/dot123/gin-gorm-admin/api/v1.(*AuthApi).RefreshToken-fm', '刷新token', '/api/v1/refresh_token', '', 'POST', '2022-09-14 14:15:43.094', '2022-09-14 14:15:43.094', '0000-00-00 00:00:00.000', 0, 0),
	(4, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Insert-fm', '创建角色', '/api/v1/role', '', 'POST', '2022-09-14 14:15:43.160', '2022-09-14 14:15:43.160', '0000-00-00 00:00:00.000', 0, 0),
	(5, 'github.com/dot123/gin-gorm-admin/api/v1.(*FileApi).UploadFile-fm', '上传文件', '/api/v1/public/uploadFile', '', 'POST', '2022-09-14 14:15:43.235', '2022-09-14 14:15:43.235', '0000-00-00 00:00:00.000', 0, 0),
	(6, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysPostApi).Insert-fm', '添加岗位', '/api/v1/post', '', 'POST', '2022-09-14 14:15:43.319', '2022-09-14 14:15:43.319', '0000-00-00 00:00:00.000', 0, 0),
	(7, 'github.com/dot123/gin-gorm-admin/api/v1.(*MsgApi).Create-fm', '新建公告', '/api/v1/msg/notice', '', 'POST', '2022-09-14 14:15:43.403', '2022-09-14 14:15:43.403', '0000-00-00 00:00:00.000', 0, 0),
	(8, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).Insert-fm', '创建菜单', '/api/v1/menu', '', 'POST', '2022-09-14 14:15:43.468', '2022-09-14 14:15:43.468', '0000-00-00 00:00:00.000', 0, 0),
	(9, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).Insert-fm', '创建用户', '/api/v1/sys-user', '', 'POST', '2022-09-14 14:15:43.519', '2022-09-14 14:15:43.519', '0000-00-00 00:00:00.000', 0, 0),
	(10, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).InsetAvatar-fm', '修改头像', '/api/v1/user/avatar', '', 'POST', '2022-09-14 14:15:43.560', '2022-09-14 14:15:43.560', '0000-00-00 00:00:00.000', 0, 0),
	(11, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).Insert-fm', '添加部门', '/api/v1/dept', '', 'POST', '2022-09-14 14:15:43.603', '2022-09-14 14:15:43.603', '0000-00-00 00:00:00.000', 0, 0),
	(12, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).GetPage-fm', 'Menu列表数据', '/api/v1/menu', '', 'GET', '2022-09-14 14:15:43.667', '2022-09-14 14:15:43.667', '0000-00-00 00:00:00.000', 0, 0),
	(13, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).Get-fm', 'Menu详情数据', '/api/v1/menu/:id', '', 'GET', '2022-09-14 14:15:43.719', '2022-09-14 14:15:43.719', '0000-00-00 00:00:00.000', 0, 0),
	(14, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).GetMenuRole-fm', '根据登录角色名称获取菜单列表数据（左菜单使用）', '/api/v1/menurole', '', 'GET', '2022-09-14 14:15:43.784', '2022-09-14 14:15:43.784', '0000-00-00 00:00:00.000', 0, 0),
	(15, 'github.com/dot123/gin-gorm-admin/api/v1.(*MonitorApi).Index-fm', '获取服务器状态', '/api/v1/monitor/index', '', 'GET', '2022-09-14 14:15:43.836', '2022-09-14 14:15:43.836', '0000-00-00 00:00:00.000', 0, 0),
	(16, 'github.com/dot123/gin-gorm-admin/api/v1.(*MsgApi).GetPage-fm', '获取公告列表', '/api/v1/msg/notice', '', 'GET', '2022-09-14 14:15:43.878', '2022-09-14 14:15:43.878', '0000-00-00 00:00:00.000', 0, 0),
	(17, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysApi).GetPage-fm', '获取接口管理列表', '/api/v1/sys-api', '', 'GET', '2022-09-14 14:15:43.927', '2022-09-14 14:15:43.927', '0000-00-00 00:00:00.000', 0, 0),
	(18, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysApi).Get-fm', '获取接口管理', '/api/v1/sys-api/:id', '', 'GET', '2022-09-14 14:15:43.969', '2022-09-14 14:15:43.969', '0000-00-00 00:00:00.000', 0, 0),
	(19, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).GetPage-fm', '列表用户信息数据', '/api/v1/sys-user', '', 'GET', '2022-09-14 14:15:44.052', '2022-09-14 14:15:44.052', '0000-00-00 00:00:00.000', 0, 0),
	(20, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).Get-fm', '获取用户', '/api/v1/sys-user/:userId', '', 'GET', '2022-09-14 14:15:44.117', '2022-09-14 14:15:44.117', '0000-00-00 00:00:00.000', 0, 0),
	(21, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).GetPage-fm', '角色列表数据', '/api/v1/role', '', 'GET', '2022-09-14 14:15:44.167', '2022-09-14 14:15:44.167', '0000-00-00 00:00:00.000', 0, 0),
	(22, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Get-fm', '获取Role数据', '/api/v1/role/:roleId', '', 'GET', '2022-09-14 14:15:44.254', '2022-09-14 14:15:44.254', '0000-00-00 00:00:00.000', 0, 0),
	(23, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).GetDeptTreeRoleSelect-fm', '部门树形', '/api/v1/roleDeptTreeselect/:roleId', '', 'GET', '2022-09-14 14:15:44.317', '2022-09-14 14:15:44.317', '0000-00-00 00:00:00.000', 0, 0),
	(24, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).GetMenuTreeSelect-fm', '根据角色ID查询菜单下拉树结构', '/api/v1/roleMenuTreeselect/:roleId', '', 'GET', '2022-09-14 14:15:44.367', '2022-09-14 14:15:44.367', '0000-00-00 00:00:00.000', 0, 0),
	(25, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).GetPage-fm', '分页部门列表数据', '/api/v1/dept', '', 'GET', '2022-09-14 14:15:44.419', '2022-09-14 14:15:44.419', '0000-00-00 00:00:00.000', 0, 0),
	(26, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).Get-fm', '获取部门数据', '/api/v1/dept/:deptId', '', 'GET', '2022-09-14 14:15:44.486', '2022-09-14 14:15:44.486', '0000-00-00 00:00:00.000', 0, 0),
	(27, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).Get2Tree-fm', '左侧部门树', '/api/v1/deptTree', '', 'GET', '2022-09-14 14:15:44.606', '2022-09-14 14:15:44.606', '0000-00-00 00:00:00.000', 0, 0),
	(28, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).GetProfile-fm', '获取个人中心用户', '/api/v1/user/profile/:userId', '', 'GET', '2022-09-14 14:15:44.744', '2022-09-14 14:15:44.744', '0000-00-00 00:00:00.000', 0, 0),
	(29, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).GetInfo-fm', '获取个人信息', '/api/v1/user/getinfo', '', 'GET', '2022-09-14 14:15:44.795', '2022-09-14 14:15:44.795', '0000-00-00 00:00:00.000', 0, 0),
	(30, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysPostApi).GetPage-fm', '岗位列表数据', '/api/v1/post', '', 'GET', '2022-09-14 14:15:44.859', '2022-09-14 14:15:44.859', '0000-00-00 00:00:00.000', 0, 0),
	(31, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysPostApi).Get-fm', '获取岗位信息', '/api/v1/post/:postId', '', 'GET', '2022-09-14 14:15:44.908', '2022-09-14 14:15:44.908', '0000-00-00 00:00:00.000', 0, 0),
	(32, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).ResetPwd-fm', '重置用户密码', '/api/v1/user/pwd/reset', '', 'PUT', '2022-09-14 14:15:44.993', '2022-09-14 14:15:44.993', '0000-00-00 00:00:00.000', 0, 0),
	(33, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).UpdatePwd-fm', '修改密码', '/api/v1/user/pwd/set', '', 'PUT', '2022-09-14 14:15:45.036', '2022-09-14 14:15:45.036', '0000-00-00 00:00:00.000', 0, 0),
	(34, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).UpdateStatus-fm', '修改用户状态', '/api/v1/user/status', '', 'PUT', '2022-09-14 14:15:45.101', '2022-09-14 14:15:45.101', '0000-00-00 00:00:00.000', 0, 0),
	(35, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Update-fm', '修改用户角色', '/api/v1/role', '', 'PUT', '2022-09-14 14:15:45.159', '2022-09-14 14:15:45.159', '0000-00-00 00:00:00.000', 0, 0),
	(36, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Update2Status-fm', '修改用户角色状态', '/api/v1/role-status', '', 'PUT', '2022-09-14 14:15:45.210', '2022-09-14 14:15:45.210', '0000-00-00 00:00:00.000', 0, 0),
	(37, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Update2DataScope-fm', '更新角色数据权限', '/api/v1/role-datascope', '', 'PUT', '2022-09-14 14:15:45.261', '2022-09-14 14:15:45.261', '0000-00-00 00:00:00.000', 0, 0),
	(38, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysApi).Update-fm', '修改接口管理', '/api/v1/sys-api', '', 'PUT', '2022-09-14 14:15:45.325', '2022-09-14 14:15:45.325', '0000-00-00 00:00:00.000', 0, 0),
	(39, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).Update-fm', '修改用户数据', '/api/v1/sys-user', '', 'PUT', '2022-09-14 14:15:45.411', '2022-09-14 14:15:45.411', '0000-00-00 00:00:00.000', 0, 0),
	(40, 'github.com/dot123/gin-gorm-admin/api/v1.(*MsgApi).Update-fm', '更新公告', '/api/v1/msg/notice', '', 'PUT', '2022-09-14 14:15:45.476', '2022-09-14 14:15:45.476', '0000-00-00 00:00:00.000', 0, 0),
	(41, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).Update-fm', '修改菜单', '/api/v1/menu', '', 'PUT', '2022-09-14 14:15:45.544', '2022-09-14 14:15:45.544', '0000-00-00 00:00:00.000', 0, 0),
	(42, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysPostApi).Update-fm', '修改岗位', '/api/v1/post', '', 'PUT', '2022-09-14 14:15:45.609', '2022-09-14 14:15:45.609', '0000-00-00 00:00:00.000', 0, 0),
	(43, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).Update-fm', '修改部门', '/api/v1/dept', '', 'PUT', '2022-09-14 14:15:45.678', '2022-09-14 14:15:45.678', '0000-00-00 00:00:00.000', 0, 0),
	(44, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysApi).Delete-fm', '删除接口管理', '/api/v1/sys-api', '', 'DELETE', '2022-09-14 14:15:45.728', '2022-09-14 14:15:45.728', '0000-00-00 00:00:00.000', 0, 0),
	(45, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysUserApi).Delete-fm', '删除用户数据', '/api/v1/sys-user', '', 'DELETE', '2022-09-14 14:15:45.792', '2022-09-14 14:15:45.792', '0000-00-00 00:00:00.000', 0, 0),
	(46, 'github.com/dot123/gin-gorm-admin/api/v1.(*MsgApi).Delete-fm', '删除公告', '/api/v1/msg/notice/:id', '', 'DELETE', '2022-09-14 14:15:45.861', '2022-09-14 14:15:45.861', '0000-00-00 00:00:00.000', 0, 0),
	(47, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysMenuApi).Delete-fm', '删除菜单', '/api/v1/menu', '', 'DELETE', '2022-09-14 14:15:45.926', '2022-09-14 14:15:45.926', '0000-00-00 00:00:00.000', 0, 0),
	(48, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysRoleApi).Delete-fm', '删除用户角色', '/api/v1/role/:roleId', '', 'DELETE', '2022-09-14 14:15:45.994', '2022-09-14 14:15:45.994', '0000-00-00 00:00:00.000', 0, 0),
	(49, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysPostApi).Delete-fm', '删除岗位', '/api/v1/post', '', 'DELETE', '2022-09-14 14:15:46.061', '2022-09-14 14:15:46.061', '0000-00-00 00:00:00.000', 0, 0),
	(50, 'github.com/dot123/gin-gorm-admin/api/v1.(*SysDeptApi).Delete-fm', '删除部门', '/api/v1/dept', '', 'DELETE', '2022-09-14 14:15:46.125', '2022-09-14 14:15:46.125', '0000-00-00 00:00:00.000', 0, 0);

CREATE TABLE IF NOT EXISTS `sys_dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 0, '/0/1/', '爱拓科技', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:06:44.960', NULL),
	(7, 1, '/0/1/7/', '研发部', 1, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:35:00.109', NULL),
	(8, 1, '/0/1/8/', '运维部', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:41:39.747', NULL),
	(9, 1, '/0/1/9/', '客服部', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:05.993', NULL),
	(10, 1, '/0/1/10/', '人力资源', 3, 'aituo', '13782218188', 'atuo@aituo.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:08.503', NULL);

CREATE TABLE IF NOT EXISTS `sys_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(128) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL,
  `parent_id` bigint DEFAULT NULL,
  `no_cache` tinyint(1) DEFAULT NULL,
  `breadcrumb` tinyint(1) DEFAULT '0',
  `component` varchar(255) DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `visible` tinyint(1) DEFAULT '0',
  `is_frame` tinyint(1) DEFAULT '0',
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=543 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `title`, `icon`, `path`, `paths`, `menu_type`, `action`, `permission`, `parent_id`, `no_cache`, `breadcrumb`, `component`, `sort`, `visible`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(2, 'Admin', '系统管理', 'api-server', '/admin', '/0/2', 'M', '无', '', 0, 1, 0, 'Layout', 10, 0, 1, 0, 1, '2021-05-20 21:58:45.679', '2021-06-17 11:48:40.703', NULL),
	(3, 'SysUserManage', '用户管理', 'user', '/admin/sys-user', '/0/2/3', 'C', '无', 'admin:sysUser:list', 2, 0, 0, '/admin/sys-user/index', 10, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 20:31:14.305', NULL),
	(43, '', '新增管理员', 'app-group-fill', '', '/0/2/3/43', 'F', 'POST', 'admin:sysUser:add', 3, 0, 0, '', 10, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 20:31:14.305', NULL),
	(44, '', '查询管理员', 'app-group-fill', '', '/0/2/3/44', 'F', 'GET', 'admin:sysUser:query', 3, 0, 0, '', 40, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 20:31:14.305', NULL),
	(45, '', '修改管理员', 'app-group-fill', '', '/0/2/3/45', 'F', 'PUT', 'admin:sysUser:edit', 3, 0, 0, '', 30, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 20:31:14.305', NULL),
	(46, '', '删除管理员', 'app-group-fill', '', '/0/2/3/46', 'F', 'DELETE', 'admin:sysUser:remove', 3, 0, 0, '', 20, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 20:31:14.305', NULL),
	(51, 'SysMenuManage', '菜单管理', 'tree-table', '/admin/sys-menu', '/0/2/51', 'C', '无', 'admin:sysMenu:list', 2, 1, 0, '/admin/sys-menu/index', 30, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(52, 'SysRoleManage', '角色管理', 'peoples', '/admin/sys-role', '/0/2/52', 'C', '无', 'admin:sysRole:list', 2, 1, 0, '/admin/sys-role/index', 20, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(56, 'SysDeptManage', '部门管理', 'tree', '/admin/sys-dept', '/0/2/56', 'C', '无', 'admin:sysDept:list', 2, 0, 0, '/admin/sys-dept/index', 40, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(57, 'SysPostManage', '岗位管理', 'pass', '/admin/sys-post', '/0/2/57', 'C', '无', 'admin:sysPost:list', 2, 0, 0, '/admin/sys-post/index', 50, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(220, '', '新增菜单', 'app-group-fill', '', '/0/2/51/220', 'F', '', 'admin:sysMenu:add', 51, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(221, '', '修改菜单', 'app-group-fill', '', '/0/2/51/221', 'F', '', 'admin:sysMenu:edit', 51, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(222, '', '查询菜单', 'app-group-fill', '', '/0/2/51/222', 'F', '', 'admin:sysMenu:query', 51, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(223, '', '删除菜单', 'app-group-fill', '', '/0/2/51/223', 'F', '', 'admin:sysMenu:remove', 51, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(224, '', '新增角色', 'app-group-fill', '', '/0/2/52/224', 'F', '', 'admin:sysRole:add', 52, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(225, '', '查询角色', 'app-group-fill', '', '/0/2/52/225', 'F', '', 'admin:sysRole:query', 52, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(226, '', '修改角色', 'app-group-fill', '', '/0/2/52/226', 'F', '', 'admin:sysRole:update', 52, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(227, '', '删除角色', 'app-group-fill', '', '/0/2/52/227', 'F', '', 'admin:sysRole:remove', 52, 0, 0, '', 1, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(228, '', '查询部门', 'app-group-fill', '', '/0/2/56/228', 'F', '', 'admin:sysDept:query', 56, 0, 0, '', 40, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(229, '', '新增部门', 'app-group-fill', '', '/0/2/56/229', 'F', '', 'admin:sysDept:add', 56, 0, 0, '', 10, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(230, '', '修改部门', 'app-group-fill', '', '/0/2/56/230', 'F', '', 'admin:sysDept:edit', 56, 0, 0, '', 30, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(231, '', '删除部门', 'app-group-fill', '', '/0/2/56/231', 'F', '', 'admin:sysDept:remove', 56, 0, 0, '', 20, 0, 1, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(232, '', '查询岗位', 'app-group-fill', '', '/0/2/57/232', 'F', '', 'admin:sysPost:query', 57, 0, 0, '', 0, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(233, '', '新增岗位', 'app-group-fill', '', '/0/2/57/233', 'F', '', 'admin:sysPost:add', 57, 0, 0, '', 0, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(234, '', '修改岗位', 'app-group-fill', '', '/0/2/57/234', 'F', '', 'admin:sysPost:edit', 57, 0, 0, '', 0, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(235, '', '删除岗位', 'app-group-fill', '', '/0/2/57/235', 'F', '', 'admin:sysPost:remove', 57, 0, 0, '', 0, 0, 1, 1, 1, '2020-04-11 15:52:48.000', '2021-06-17 11:48:40.703', NULL),
	(528, 'SysApiManage', '接口管理', 'api-doc', '/admin/sys-api', '/0/527/528', 'C', '无', 'admin:sysApi:list', 2, 0, 0, '/admin/sys-api/index', 0, 0, 0, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(529, '', '查询接口', 'app-group-fill', '', '/0/527/528/529', 'F', '无', 'admin:sysApi:query', 528, 0, 0, '', 40, 0, 0, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL),
	(531, '', '修改接口', 'app-group-fill', '', '/0/527/528/531', 'F', '无', 'admin:sysApi:edit', 528, 0, 0, '', 30, 0, 0, 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', NULL);

CREATE TABLE IF NOT EXISTS `sys_menu_api_rule` (
  `sys_menu_menu_id` bigint NOT NULL,
  `sys_api_id` bigint NOT NULL,
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`),
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES
	(224, 4),
	(233, 6),
	(220, 8),
	(43, 9),
	(229, 11),
	(51, 12),
	(222, 13),
	(528, 17),
	(529, 18),
	(3, 19),
	(44, 20),
	(52, 21),
	(225, 22),
	(56, 25),
	(228, 26),
	(57, 30),
	(232, 31),
	(226, 35),
	(531, 38),
	(45, 39),
	(221, 41),
	(234, 42),
	(230, 43),
	(46, 45),
	(223, 47),
	(227, 48),
	(235, 49),
	(231, 50);

CREATE TABLE IF NOT EXISTS `sys_post` (
  `post_id` bigint NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_post` (`post_id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, '首席执行官', 'CEO', 0, 2, '首席执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL),
	(2, '首席技术执行官', 'CTO', 2, 2, '首席技术执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL),
	(3, '首席运营官', 'COO', 3, 2, '测试工程师', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

CREATE TABLE IF NOT EXISTS `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` tinyint DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, '系统管理员', 2, 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

CREATE TABLE IF NOT EXISTS `sys_role_dept` (
  `role_id` bigint NOT NULL,
  `dept_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`),
  KEY `fk_sys_role_dept_sys_dept` (`dept_id`),
  CONSTRAINT `fk_sys_role_dept_sys_dept` FOREIGN KEY (`dept_id`) REFERENCES `sys_dept` (`dept_id`),
  CONSTRAINT `fk_sys_role_dept_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `sys_role_menu` (
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `nick_name` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `role_id` bigint DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `dept_id` bigint DEFAULT NULL,
  `post_id` bigint DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_user_dept` (`dept_id`),
  CONSTRAINT `fk_sys_user_dept` FOREIGN KEY (`dept_id`) REFERENCES `sys_dept` (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'zhangwj', '13818888888', 1, '', '', 1, '1@qq.com', 1, 1, '', 2, 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

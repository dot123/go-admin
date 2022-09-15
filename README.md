# gin-gorm-admin

基于[gin-gorm-admin](https://github.com/dot123/gin-gorm-admin)实现的权限管理系统。

> 账号：admin 密码：123456

## 特性

* 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
* 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestRateLimiter、Recover、GZIP）
* 基于Casbin的 RBAC 访问控制模型
* 基于[jwt](https://github.com/golang-jwt/jwt) 认证
* 基于[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验
* 用Docker上云
* 在token过期后的一个小时内，用户再次操作会要求重新登陆
* 基于[swaggo](https://github.com/swaggo)为Go工程生成自动化接口文档
* 基于[wire](https://github.com/google/wire)依赖注入
* 基于[gorm](https://gorm.io/zh_CN/)全功能ORM
* 基于[air](https://github.com/cosmtrek/air)自动编译，重启程序
* 基于redis限制请求频率

##  内置
1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识，接口权限等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 接口文档：根据业务代码自动生成相关的api接口文档。

### 项目结构
<pre><code>
├─api
├─cmd
├─configs
├─docs
├─internal
│  ├─config
│  ├─contextx
│  ├─errors
│  ├─ginx
│  ├─middleware
│  ├─models
│  ├─schema
│  ├─service
│  └─validate
└─pkg
    ├─fileStore
    ├─gormx
    ├─hash
    ├─helper
    ├─logger
    ├─monitor
    ├─rabbitMQ
    ├─redisHelper
    ├─store
    ├─timer
    ├─types
    └─validate
</code></pre>

### 下载依赖

<pre><code>depend.cmd</code></pre>

### 代码生成与运行

##### 生成

<pre><code>generate.cmd</code></pre>

##### 数据库

<pre><code>gin-admin.sql</code></pre>

##### 运行

<pre><code>run.cmd 或go run ./cmd/gin-gorm-admin/ web -c ./configs/config.toml</code></pre>

##### sys_api表的数据如何添加

在项目启动时，使用`-a true` 系统会自动添加缺少的接口数据

<pre><code>go run ./cmd/gin-gorm-admin/ web -c ./configs/config.toml -a true</code></pre>

##### docker部署

<pre><code>deploy.cmd</code></pre>

## 服务程序开发实战 - Agenda
### 1、概述

   利用命令行 或 web 客户端调用远端服务是服务开发的重要内容。其中，要点是如何实现 API First 开发，使得团队协作变得更有效率。

#### 任务目标

1. 熟悉 API 设计工具，实现从资源（领域）建模，到 API 设计的过程
2. 使用 Github ，通过 API 文档，实现 agenda 命令行项目 与 RESTful 服务项目同步开发
3. 使用 API 设计工具提供 Mock 服务，两个团队独立测试 API
4. 使用 travis 测试相关模块
4. 利用 dockerfile 在 docker hub 上构建一个镜像，同时包含 agenda cli 和 agenda service， 如果 mysql 包含 服务器 和 客户端一样
### 2、agenda 开发项目

重构、或新建 agenda 项目，根目录必须包含
> cli 目录
service 目录
.travis
apiary.apib
dockerfile
LICENSE
README.md
README-yourid.md
记录你的工作摘要（个人评分依据）

#### - API 开发
- 使用 API Blueprint 设计 API
- 资源 URL 命名符合 RESTful 设计标准
- 资源 CRUD 基本完整

#### - API 客户端开发
- 可用命令 5 个以上
- 必须有 XXX-test.go 文件


#### - 服务端开发
- 使用 sqlite3 作为数据库
- 建议使用课程提供的服务端框架
- 必须有 XXX-test.go 文件

#### - 容器镜像制作
- 在 docker hub 上生成镜像
- base 镜像 go-1.8
- 需要加载 sqlite3
- 同时包含客户端与服务器

#### - README.md
- 有 build pass 标签
- 有简短使用说明
- 有系统测试的结果（包含如何下载镜像，如何启动服务器，如何使用命令行，cli 的 mock 测试结果， 综合系统测试结果）

#### - README-yourid.md
- fork 项目的位置
- 个人工作摘要（每次提交）
- 项目小结

## Agenda 命令设计

### 命令实现
>agenda help [option]

列出命令说明(可选是否列出具体功能的说明)。

>agenda register -uUserName --password pass -email=a@xxx.com -phone=phoneNum

用户注册，如果用户名已被使用，返回错误信息；如果注册成功，返回成功信息。

>agenda login -uUserName --password pass

用户登录，登录失败返回失败原因;登录成功，返回成功信息，并列出可选操作。

>agenda logout

用户退出登录，返回成功信息并列出可选操作。

>agenda -lUser

已登录用户查询已注册用户信息（用户名、邮箱、电话）

>agenda delete

已登录用户注销帐号，操作成功返回成功信息；否则，返回失败信息。若成功，删除一切与该用户的信息。

>agenda mkmeeting --title Name --participator user1 [user2 ....] --stime start --etime end

成功，则返回成功信息及注册信息；失败，则返回失败原因。

>agenda meetingadd --participator user1 [user2 ...]

成功，则返回新增参与者信息；失败，返回失败原因。

> agenda meetingdel --participator user1 [user2 ...]

成功，则返回成功信息（如果会议因为删除参与者而删除，返回额外信息）；失败，返回失败原因。

>agenda querymeeting -stime start -etime end

已登录用户查询自己的会议议程。

>agenda meetingdel --title Name

已登录用户删除会议。

>agenda meetingout --ttile Name

已登录用户退出自己参与的某一会议，若因此造成会议参与者为0,则会议被删除。

>agenda meetingclear

已登录用户清空自己发起的所有会议。

![](https://raw.githubusercontent.com/moandy/cloudgo/67124ed6590af80f3c969706eb6e2483a6760a09/picture/p1.png)
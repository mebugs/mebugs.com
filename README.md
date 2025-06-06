米虫 | [mebugs.com](http://www.mebugs.com) 独立博客的开源程序仓库。

本仓库技术栈：GoLang VueNext Vite Nginx MySQL Redis Nuxt。

# 概述

## 重构缘由：

1. 彻底解决 PHP 本身内存开销过高问题
2. 享受 GoLang 云原生的技术红利
3. 网页排版整体重排

## 仓库结构：

1. admin/admin-ui 管理后台前端，Go/VueNext
2. page-ui 采用Nuxt实现SSR渲染，兼顾SEO和访问性能

## 其他说明：

- 本仓库内容，仅供阅读、学习。
- 本仓库后端/前端基于Smart开发。[Gitee](https://gitee.com/siteol-com/smart)  [GitHub](https://github.com/siteol-com/smart) 


# 安装

当前尚未计划编写Docker镜像，因为大部分个人博主的服务器资源有限，请参考以下手顺。

本博客不是一个通用的软件，具有极高的个人定制性，因此比较建议您按照个人意愿二次开发。

本手顺内容具有一定的专业性，不建议纯新手尝试，给您带来不便，进请谅解。

手顺如有遗漏或错误地方，请提交评论。

## 服务端安装

- 本地启动建议通过`GoLand`。

通过Golang打包成二进制文件。

Windows环境打包命令：

```bash
go env -w CGO_ENABLED=0 
go env -w GOOS=linux 
go env -w GOARCH=amd64 
go build
```

编译完成后，看`smart`的二进制文件，你需要先恢复正常的配置，避免Goland本地启动异常。

```bash
go env -w CGO_ENABLED=1
go env -w GOOS=windows 
```

请将`admin/config`下的`smart_init.sql`刷入您的数据库中。

具体库名你可按照您自己希望的数据库来，但您需要记住同步修改配置`config_prod.json`。

您需要将`smart`的二进制文件，`restart.sh`，`config`目录一起提交到环境中，保持如下结构。

```bash
- config
  - config_prod.json
- smart
- restart.sh
```

通过`restart.sh`启动应用，应用的日志会写入`logs`目录下。

如果您未修改端口，应用默认监听`8000`端口。

## 管理平台

- 管理平台本地依赖`Node`建议18+，最优20+。建议使用`pnpm`管理您的依赖。

本地编译命令。

```bash
pnpm run build
```

编译完成在`admin-ui`下的`dist`为您的编译结果文件。

您需要在`admin-ui\dist\source`下适用于存储上传图片的位置，当您将`dist`下的内容部署到生产后。

需要在数据库修改`sys_config`的`file_full_path`，建议绝对路径。

您可以通过管理平台的系统管理来修改这个数据。

将`dist`下的静态文件部署到您想放的位置，你需要记住这个位置，后续Nginx会需要配置这个路径。

本手顺建议的位置在`/www/wwwroot/admin-ui`（推荐您使用宝塔面板来配置），将`dist`下的内容传到该路径下。

静态前端编译后无需启动，由Nginx指向。

`sys_config`的`file_full_path` = `/www/wwwroot/admin-ui/source`。

如果您在本地运行是，这个地址配置示例：`'D:\\Mebugs\\mebugs\\admin-ui\\public\\source\\'`，根据您实际地址修改。

创建NginxServer不在本手顺范围内，您需要追加的Nginx配置大致如下，建议您为管理平台开辟单独的子域名。

```nginx
server
{
    listen 80; # 您自己HTTP或HTTPS端口
    server_name xxx.xxx.com; # 您自己的服务域名
	# 默认首页
    index index.html;
	# 项目位置
    root /www/wwwroot/admin-ui;
	# VUE核心配置
    try_files $uri $uri/ /index.html;
    #转发api
    location /api/ {
    proxy_pass  http://127.0.0.1:8000/; # 转发规则
    proxy_set_header Host $proxy_host; 
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }   
    # 其他更多
}
```

## Nuxt客户端部署

- Nuxt可以很好的在服务端渲染网页，比纯静态页面有更好的切换性能，即切换页面无需重新加载。

你需要在环境上安装一个Node服务，建议18+，最优20+（安装Node略）。

`nuxt.config.ts`的`runtimeConfig`需要修改为您自己所希望的SEO信息。

您需要自行修改`components\Foot.vue`下的个人内容。

本地编译命令。

```bash
pnpm run build
```

编译成果位于`mebugs\page-ui\.output`。

上传`.output`到环境中，本手顺示例位于`/www/wwwroot/page-ui`。

启动命令`node /www/wwwroot/page-ui/.output/server/index.mjs  >> page.log 2>&1 &`

应用默认端口为`3000`，Nginx配置参考如下。

```bash
server
{
    listen 80; # 您自己HTTP或HTTPS端口
    server_name xxx.xxx.com; # 您自己的服务域名
	#转发客户端请求
    location / {
	proxy_set_header	Host $host;
	proxy_set_header  	X-Real-IP $remote_addr;
	proxy_set_header  	X-Forwarded-For $proxy_add_x_forwarded_for;
	proxy_pass			http://127.0.0.1:3000/;
    }  
    #转发api
    location /api/ {
    proxy_pass  http://127.0.0.1:8000/; # 转发规则
    proxy_set_header Host $proxy_host; 
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }   
    # 其他更多
}
```

## 其他说明

`/www/wwwroot/page-ui/.output/public/source`生产中建议删除，创建一个软连接指向`/www/wwwroot/admin-ui/source`。

让客户端可以直接都渠道，上传的图片。


# 感谢 

感谢您的阅读。
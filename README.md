# gcwguide

#### 介绍
使用与公司内部或个人的简易版网址导航工具

#### 示例站点：https://nav.jianean.com/
[demo](https://nav.jianean.com/)

###### 操作账号
- 用户名：add
- 密码：123456
- 权限：只能添加，不能删除和编辑

#### 软件架构
1.前端：vue3 + element
2.后端：golang + gin
3.存储：文件存储，直接存储在服务器上，未使用任何数据库
4.当前版本不支持自动备份数据，请自行备份，备份目录为 config.yaml指定的路径：store/websites/webs.json (以实际自己的配置为准，相对于接口服务gcapi的根目录)


#### 安装教程

以下均指Centos7.x服务器

###### 方式一

1. 构建后端：
   1. 更改配置文件：gcapi/config/config.yaml 配置如端口号，上传目录，上传图片的访问域名（可以web域名保持一致，斜杠/结尾）
   2. 安装golang >= 1.14.3
   3. 配置go mod
   4. 上传gcapi整个文件夹到服务器，示例保存目录：/data/www/gcwguide/gcapi
   5. 进入gcapi目录，执行 go build
   6. 除了生成的可执行文件 gcapi和config/config.yaml不可删除外其他的gcapi中文件都可去掉了（也可以保留）。
   7. 将gcapi.service 中的执行文件路径改成自己的路径，并移动到/etc/systemd/system，gcapi.service示例：
```
[Unit]
Description=Gcapi Api Service
After=network.target

[Service]
Type=simple
WorkingDirectory=/data/server/api/gcapi
ExecStart=/data/server/api/gcapi/gcapi
Restart=always

[Install]
WantedBy=multi-user.target

```
   8. 执行：systemctl daemon-reload 
   9.  执行：systemctl start gcapi.service 启动接口服务
   10. 执行：systemctl status gcapi.service 查看接口服务启动状态，或者执行netstat -ntlp | grep :xxxx端口 查看端口是否被监听
   11. 后端完毕


2. 构建前端：（vue3+element）
   1. npm install
   2. npm run build
   3. 上传./dist文件夹中的全部文件到自己的服务器指定的目录，示例：/data/www/gcwguide/gcweb


3. 配置nginx，
   1. 示例配置：
```
server {
    listen       80;
    server_name  nav.jianean.com;

    #charset koi8-r;
    access_log  /var/log/nginx/nav.jianean.com.access.log  main;
    error_log   /var/log/nginx/nav.jianean.com.error.log;
    error_page   500 502 503 504  /50x.html;
	
    gzip            on;
    gzip_types      text/plain application/xml text/css application/javascript;
    gzip_min_length 1000;

    location / {
        #expires $expires;
        proxy_redirect                      off;
        proxy_set_header Host               $host;
        proxy_set_header X-Real-IP          $remote_addr;
        proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto  $scheme;
        proxy_read_timeout          1m;
        proxy_connect_timeout       1m;
        proxy_pass http://127.0.0.1:xxxx; # gcapi中的config/config.yaml的port端口号
    }
}

```
   3. 执行 nginx -t
   4. 执行 nginx -s reload
   5. 访问域名，查看是否可以正常访问了，如果不能正常访问请自行排查原因，或可在 https://www.jianean.com 上留言 或发送邮件到 374695440@qq.com 

###### 方式二
使用已经打包好的文件
-   前端：gcweb/dist/ 下的全部文件
-   后端：gcapi 下的 ./gcapi ./config/config.yaml 一个可执行文件和 config 目录及 config.yaml
-   按照方式一的后面几部配置

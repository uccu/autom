# AUTOM


#### 功能简要

web钩子连接地址：http://[host]:[port]/hook

自动接收来自git的推送事件（目前支持push/tag_push）来构建运行目标项目, 并用反代或端口映射来暴露端口

- 配置文件
```json
[{
    "network_name": "autom",  // 网络名称
    "subnet": "172.20.0.0/16", // 网段
    "containers": [{
            "name": "test", // 项目名称
            "tag": false, // 是否更新tag时候触发，不包括删除tag，为false时则更新分支时触发
            "branch": "master", // 分支名字，tag为true时无效
            "ip": "172.20.0.2", // 创建容器ip
            "url": "git@test.com:test/test.git", // 项目git地址
            "volumes": {}, // true dir: container dir的字典
            "token": "123123123123123" // gitlab 验证用TOKEN
        },
        {
            "name": "test",
            "tag": true,
            "ip": "172.20.0.3",
            "url": "git@test.com:test/test.git",
            "volumes": {},
            "token": "123123123123123"
        }
    ]
}]
```

- 运行

```sh
autom start
```

- 停止

```sh
autom stop
```

- 后台运行

```sh
autom start -d
```

- 指定端口运行

```sh
autom start -p 443
```

- 指定配置

```sh
autom start -c base.json
```

- 指定运行目录

```sh
autom start -e /home/autom
```

- 指定输出日志文件

```sh
autom start -l info.log
```


#### 未来规划

- 优化通过pid文件的方式开启关闭服务
- docker 支持远程api操作
- 非git方式，文件上传方式/url下载方式更新项目
- 用户系统
- windows系统接收项目上线消息功能
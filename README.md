## Books
图书管理系统后端

## 介绍

* 该系统应用于web端
* 实现了book的存储


## 安装和使用

克隆仓库到本地
```shell
git clone https://github.com/zhendong233/Books.git
```

创建mysql容器并启动
```shell
docker-compose up -d books-mysql
```

创建数据表
```shell
go run ./migration
```

运行
```shell
go run main.go
```
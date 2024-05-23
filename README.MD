# GoFrameWork
`GoFrameWork` is a modular, powerful, high-performance and enterprise-class application development framework of Golang.


# 本仓库使用方法

注意go版本必须大于 1.21

~~~sh
# 将本项目下载到和你的项目评级的文件夹下
git clone https://github.com/tekintian/gofw.git

# 进入你的项目根目录
cd yourproject
#
# 然后执行下面的命令将gogf/gf的模块引用指向你下载的本仓库目录，相对路径即可
# 主模块绑定
go mod edit -replace=github.com/gogf/gf/v2=../gofw
# 添加依赖
go get -v github.com/gogf/gf/v2


# mysql模块绑定， 其他的模块和这个方法一样，注意修改为对应的目录即可
go mod edit -replace=github.com/gogf/gf/contrib/drivers/mysql/v2=../gofw/contrib/drivers/mysql
go get github.com/gogf/gf/contrib/drivers/mysql/v2

# pgsql驱动模块
go mod edit -replace=github.com/gogf/gf/contrib/drivers/pgsql/v2=../gofw/contrib/drivers/pgsql
go get github.com/gogf/gf/contrib/drivers/pgsql/v2

#nacos配置
go mod edit -replace=github.com/gogf/gf/contrib/config/nacos/v2=../gofw/contrib/config/nacos
go get github.com/gogf/gf/contrib/config/nacos/v2


# 依赖绑定并引入完成后执行 go mod tidy对模块依赖进行更新
go mod tidy

# ps: 前面为模块名称， 后面为模块对应的你本地的路径，即 go.mod文件所在的路径， 其他的模块的绑定方式一样
~~~

- cli tool install
~~~sh
# 查看当前模块的安装路径
go list -f '{{.Target}}'

# 进入工具模块所在路径
cd cmd/gf
# 编译并安装
go install .



~~~


# Features

- modular, loosely coupled design
- rich components, out-of-the-box
- automatic codes generating for efficiency
- simple and easy to use, detailed documentation
- interface designed components, with high scalability
- fully supported tracing and error stack feature
- specially developed and powerful ORM component
- robust engineering design specifications
- convenient development CLI tool provide
- OpenTelemetry observability features support
- OpenAPIV3 documentation generating, automatically
- much, much more...ready to explore?


# Limitation

```
golang version >= 1.21
```

# Documentation

- Chinese Official Site(中文官网): [https://goframe.org](https://goframe.org/display/gf)
- Chinese Pages Document(中文镜像文档): [https://pages.goframe.org](https://pages.goframe.org)
- Chinese Offline Document(中文离线文档): [https://github.com/gogf/goframe.org-pdf](https://github.com/gogf/goframe.org-pdf)
- GoDoc API: [https://pkg.go.dev/github.com/gogf/gf/v2](https://pkg.go.dev/github.com/gogf/gf/v2)

# License

`GoFrame` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.


# Thanks
github.com/gogf/gf

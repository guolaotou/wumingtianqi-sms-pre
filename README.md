# 无名天气

## 配置文件使用
- conf目录下存放配置文件
    - 参考config.template.json, 在本地新建config.json文件
    - 将其中的数据库配置等改成自己的实际配置

- 配置文件的加载代码在config文件下
    - config.go控制加载文件
    - content.go映射config.json配置文件的内容

## 运行方法
将项目放到go目录下，然后在项目目录下运行
```shell script
go run main.go
```


## 打包
1. 包需要在mac上运行
    export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && go build -o wumingtianqi.out -v wumingtianqi

2. 包需要再linux上运行
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build -o wumingtianqi.out -v wumingtianqi

## 加速测试方法
mac系统，项目目录下运行加速系统时间的脚本
```shell script
sudo go run scripts/speed_macbook_time.go
```

## 开发过程及文档
见docs